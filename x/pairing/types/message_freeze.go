package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgFreeze   = "freeze"
	ReasonFlagName  = "reason"
	ReasonMaxLength = 50
)

var _ sdk.Msg = &MsgFreezeProvider{}

func NewMsgFreeze(creator string, chainIds []string, reason string) *MsgFreezeProvider {
	return &MsgFreezeProvider{
		Creator:  creator,
		ChainIds: chainIds,
		Reason:   reason,
	}
}

func (msg *MsgFreezeProvider) Route() string {
	return RouterKey
}

func (msg *MsgFreezeProvider) Type() string {
	return TypeMsgFreeze
}

func (msg *MsgFreezeProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.GetReason()) > ReasonMaxLength {
		return sdkerrors.Wrapf(FreezeReasonTooLongError, "invalid freeze reason error (%s) ", FreezeReasonTooLongError.Error())
	}
	return nil
}
