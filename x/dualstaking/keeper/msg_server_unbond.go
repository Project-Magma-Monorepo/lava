package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/v5/utils"
	commontypes "github.com/lavanet/lava/v5/utils/common/types"
	"github.com/lavanet/lava/v5/x/dualstaking/types"
)

func (k msgServer) Unbond(goCtx context.Context, msg *types.MsgUnbond) (*types.MsgUnbondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgUnbondResponse{}, k.Keeper.UnbondFull(ctx, msg.Creator, msg.Validator, msg.Provider, msg.Amount, false)
}

// UnbondFul uses staking module for to unbond with hooks
func (k Keeper) UnbondFull(ctx sdk.Context, delegator string, validator string, provider string, amount sdk.Coin, stake bool) error {
	// 1.redelegate from the provider to the empty provider
	// 2.calls staking module to unbond from the validator
	// 3.calls the hooks to than unbond from the empty provider

	addr, err := sdk.ValAddressFromBech32(validator)
	if err != nil {
		return err
	}
	delegatorAddress, err := sdk.AccAddressFromBech32(delegator)
	if err != nil {
		return err
	}
	shares, err := k.stakingKeeper.ValidateUnbondAmount(
		ctx, delegatorAddress, addr, amount.Amount,
	)
	if err != nil {
		return err
	}

	err = utils.ValidateCoins(ctx, k.stakingKeeper.BondDenom(ctx), amount, false)
	if err != nil {
		return err
	}

	err = k.Redelegate(
		ctx,
		delegator,
		provider,
		commontypes.EMPTY_PROVIDER,
		amount,
		stake,
	)
	if err != nil {
		return err
	}
	_, err = k.stakingKeeper.Undelegate(ctx, delegatorAddress, addr, shares)
	if err != nil {
		return err
	}

	logger := k.Logger(ctx)
	details := map[string]string{
		"delegator": delegator,
		"provider":  provider,
		"amount":    amount.String(),
	}
	utils.LogLavaEvent(ctx, logger, types.UnbondingEventName, details, "Unbond")

	return nil
}
