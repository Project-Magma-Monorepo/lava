package types

import (
	"math"
	"strconv"
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func initCtxAndTimerStores(t *testing.T, count int) (sdk.Context, []*TimerStore) {
	ctx, cdc := initCtx(t)
	tstore := make([]*TimerStore, count)

	for i := 0; i < count; i++ {
		timerKey := "mock_timer_" + strconv.Itoa(i)
		tstore[i] = NewTimerStore(mockStoreKey, cdc, timerKey)
	}

	return ctx, tstore
}

type timerTemplate struct {
	op    string
	name  string
	store int
	value uint64
	key   string
	data  string
	fire  int
}

// helper to automate testing operations
func testWithTimerTemplate(t *testing.T, playbook []timerTemplate, countTS int) {
	ctx, tstore := initCtxAndTimerStores(t, countTS)

	var (
		callbackBlockCount int
		callbackBlockKey   string
		callbackBlockData  string
		callbackTimeCount  int
		callbackTimeKey    string
		callbackTimeData   string
	)

	callbackHeight := func(ctx sdk.Context, key, data []byte) {
		callbackBlockCount += 1
		callbackBlockKey += string(key)
		callbackBlockData += string(data)
	}

	callbackTime := func(ctx sdk.Context, key, data []byte) {
		callbackTimeCount += 1
		callbackTimeKey += string(key)
		callbackTimeData += string(data)
	}

	for i := range tstore {
		tstore[i] = tstore[i].
			WithCallbackByBlockHeight(callbackHeight).
			WithCallbackByBlockTime(callbackTime)
	}

	for _, play := range playbook {
		what := play.op + " " + play.name +
			" value: " + strconv.Itoa(int(play.value)) +
			" key: " + play.key +
			" data: " + play.data
		key := []byte(play.key)
		data := []byte(play.data)
		switch play.op {
		case "addheight":
			tstore[play.store].AddTimerByBlockHeight(ctx, play.value, key, data)
		case "hasheight":
			has := tstore[play.store].HasTimerByBlockHeight(ctx, play.value, key)
			require.Equal(t, play.data == "has", has)
		case "delheight":
			tstore[play.store].DelTimerByBlockHeight(ctx, play.value, key)
		case "addtime":
			tstore[play.store].AddTimerByBlockTime(ctx, play.value, key, data)
		case "hastime":
			has := tstore[play.store].HasTimerByBlockTime(ctx, play.value, key)
			require.Equal(t, play.data == "has", has)
		case "deltime":
			tstore[play.store].DelTimerByBlockTime(ctx, play.value, key)
		case "nextheight":
			value := tstore[play.store].GetNextTimeoutBlockHeight(ctx)
			require.Equal(t, play.value, value, what)
		case "nexttime":
			value := tstore[play.store].GetNextTimeoutBlockTime(ctx)
			require.Equal(t, play.value, value, what)
		case "tickheight":
			callbackBlockCount = 0
			callbackBlockKey = ""
			callbackBlockData = ""
			ctx = ctx.WithBlockHeight(int64(play.value))
			tstore[play.store].Tick(ctx)
			require.Equal(t, play.fire, callbackBlockCount, what)
			require.Equal(t, play.key, callbackBlockKey, what)
			require.Equal(t, play.data, callbackBlockData, what)
		case "ticktime":
			callbackTimeCount = 0
			callbackTimeKey = ""
			callbackTimeData = ""
			ctx = ctx.WithBlockTime(time.Unix(int64(play.value), 0))
			tstore[play.store].Tick(ctx)
			require.Equal(t, play.fire, callbackTimeCount, what)
			require.Equal(t, play.key, callbackTimeKey, what)
			require.Equal(t, play.data, callbackTimeData, what)
		}
	}
}

// Test single timer by block height
func TestTimerBlockHeight(t *testing.T) {
	playbook := []timerTemplate{
		{op: "nextheight", name: "next timeout infinity", value: math.MaxUint64},
		{op: "tickheight", name: "tick without timers", value: 100, fire: 0},
		{op: "addheight", name: "add timer no-1", value: 120, key: "a", data: "no-1."},
		{op: "hasheight", name: "has timer no-1", value: 120, key: "a", data: "has"},
		{op: "nextheight", name: "next timeout no-1", value: 120},
		{op: "tickheight", name: "tick before timer no-1", value: 110, fire: 0},
		{op: "tickheight", name: "tick after timer no-1", value: 130, key: "a", fire: 1, data: "no-1."},
		{op: "hasheight", name: "gone timer no-1", value: 120, key: "a", data: "gone"},
		{op: "nextheight", name: "next timeout no-1", value: math.MaxUint64},
		{op: "addheight", name: "add timer no-2", value: 140, key: "a", data: "no-2."},
		{op: "hasheight", name: "has timer no-2", value: 140, key: "a", data: "has"},
		{op: "tickheight", name: "tick exactly on timer no-2", value: 140, key: "a", fire: 1, data: "no-2."},
		{op: "nextheight", name: "next timeout infinity again", value: math.MaxUint64},
	}

	testWithTimerTemplate(t, playbook, 1)
}

// Test single timer by block time
func TestTimerBlockTime(t *testing.T) {
	playbook := []timerTemplate{
		{op: "nexttime", name: "next timeout infinity", value: math.MaxUint64},
		{op: "ticktime", name: "tick without timers", value: 100, fire: 0},
		{op: "addtime", name: "add timer no-1", value: 120, key: "b", data: "no-1."},
		{op: "hastime", name: "has timer no-1", value: 120, key: "b", data: "has"},
		{op: "nexttime", name: "next timeout no-1", value: 120},
		{op: "ticktime", name: "tick before timer no-1", value: 110, fire: 0},
		{op: "ticktime", name: "tick after timer no-1", value: 130, key: "b", fire: 1, data: "no-1."},
		{op: "hastime", name: "gone timer no-1", value: 120, key: "b", data: "gone"},
	}

	testWithTimerTemplate(t, playbook, 1)
}

// Test new timer earlier than next timeout
func TestTimerEarlierThenNext(t *testing.T) {
	playbook := []timerTemplate{
		{op: "tickheight", name: "tick without timers", value: 100, fire: 0},
		{op: "addheight", name: "add timer no 1", value: 120, key: "a", data: "no-1."},
		{op: "nextheight", name: "next timeout no-1", value: 120},
		{op: "addheight", name: "add timer no 2 (as first)", value: 110, key: "b", data: "no-2."},
		{op: "nextheight", name: "next timeout no-2", value: 110},
		{op: "tickheight", name: "tick before all", value: 105, fire: 0},
		{op: "tickheight", name: "tick between no-2, no-1", value: 115, key: "b", fire: 1, data: "no-2."},
		{op: "nextheight", name: "next timeout no-1 again", value: 120},
		{op: "tickheight", name: "tick after no-1", value: 125, fire: 1, key: "a", data: "no-1."},
	}

	testWithTimerTemplate(t, playbook, 1)
}

// Test multiple timers (by block height)
func TestMultipleTimers(t *testing.T) {
	playbook := []timerTemplate{
		{op: "tickheight", name: "tick without timers", value: 100, fire: 0},
		{op: "addheight", name: "add timer no 1", value: 120, key: "a", data: "no-1."},
		// also use two timers with same block
		{op: "addheight", name: "add timer no 2a", value: 130, key: "bx", data: "no-2a."},
		{op: "addheight", name: "add timer no 2b", value: 130, key: "by", data: "no-2b."},
		{op: "addheight", name: "add timer no 3", value: 140, key: "c", data: "no-3."},
		{op: "addheight", name: "add timer no 4", value: 150, key: "d", data: "no-4."},
		// also overwrite existing timer
		{op: "addheight", name: "add timer no 4 (again)", value: 150, key: "d", data: "no-4x."},
		{op: "tickheight", name: "tick before all", value: 110, fire: 0},
		{op: "tickheight", name: "tick between no-2,no-3", value: 135, key: "abxby", fire: 3, data: "no-1.no-2a.no-2b."},
		{op: "nextheight", name: "next timeout no-3", value: 140},
		{op: "tickheight", name: "tick after all", value: 155, fire: 2, key: "cd", data: "no-3.no-4x."},
	}

	testWithTimerTemplate(t, playbook, 1)
}

// Test delete timers (by block height)
func TestDeleteTimers(t *testing.T) {
	playbook := []timerTemplate{
		{op: "tickheight", name: "tick without timers", value: 100, fire: 0},
		{op: "addheight", name: "add timer no 1", value: 120, key: "a", data: "no-1."},
		{op: "addheight", name: "add timer no 2a", value: 130, key: "bx", data: "no-2a."},
		{op: "addheight", name: "add timer no 2b", value: 130, key: "by", data: "no-2b."},
		{op: "addheight", name: "add timer no 3", value: 140, key: "c", data: "no-3."},
		{op: "tickheight", name: "tick before all", value: 110, fire: 0},
		{op: "delheight", name: "del timer no 2a", value: 130, key: "bx"},
		{op: "hasheight", name: "gone timer no 2a", value: 130, key: "bx", data: "gone"},
		{op: "tickheight", name: "tick between no-2,no-3", value: 135, key: "aby", fire: 2, data: "no-1.no-2b."},
		{op: "delheight", name: "del timer no 3", value: 140, key: "c"},
		{op: "hasheight", name: "gone timer no 3", value: 140, key: "c", data: "gone"},
		{op: "nextheight", name: "next timeout no-3", value: 140},
		{op: "tickheight", name: "tick after all", value: 155, fire: 0, key: "", data: ""},
	}

	testWithTimerTemplate(t, playbook, 1)
}

// Test delete non-existent timers (by block height)
func TestBadDeleteTimers(t *testing.T) {
	playbooks := [][]timerTemplate{
		{
			{op: "tickheight", name: "tick without timers", value: 100, fire: 0},
			{op: "delheight", name: "del non-existing timer", value: 130, key: "bx"},
		},
		{
			{op: "tickheight", name: "tick without timers", value: 100, fire: 0},
			{op: "addheight", name: "add timer no 1", value: 120, key: "a", data: "no-1."},
			{op: "delheight", name: "del non-existing timer", value: 120, key: "a"},
			{op: "delheight", name: "del non-existing timer( again)", value: 120, key: "a"},
		},
	}

	for _, p := range playbooks {
		what := p[0].op + " " + p[0].name
		require.Panics(t, func() { testWithTimerTemplate(t, p, 1) }, what)
	}
}

// Test expiry before current block
func TestTimerEarlyExpiry(t *testing.T) {
	playbooks := [][]timerTemplate{
		{
			{op: "tickheight", name: "tick height without timers", value: 100, fire: 0},
			{op: "addheight", name: "add timer no 1", value: 100, key: "a", data: "no-1."},
		},
		{
			{op: "ticktime", name: "tick time without timers", value: 100, fire: 0},
			{op: "addtime", name: "add timer no 1", value: 100, key: "a", data: "no-1."},
		},
	}

	for _, p := range playbooks {
		what := p[0].op + " " + p[0].name
		require.Panics(t, func() { testWithTimerTemplate(t, p, 1) }, what)
	}
}

var (
	mockStoreKey    = storetypes.NewKVStoreKey("storeKey")
	mockMemStoreKey = storetypes.NewMemoryStoreKey("storeMemKey")
)

// Helper function to init a mock keeper and context
func initCtx(t *testing.T) (sdk.Context, *codec.ProtoCodec) {
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	stateStore.MountStoreWithDB(mockStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(mockMemStoreKey, storetypes.StoreTypeMemory, nil)

	require.NoError(t, stateStore.LoadLatestVersion())

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return ctx, cdc
}
