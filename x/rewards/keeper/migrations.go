package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// MigrateVersion1To2 sets the min IPRPC cost to be 100LAVA = 100,000,000ulava
func (m Migrator) MigrateVersion1To2(ctx sdk.Context) error {
	return m.keeper.SetIprpcData(ctx, sdk.NewCoin(m.keeper.stakingKeeper.BondDenom(ctx), math.NewInt(100000000)), []string{})
}
