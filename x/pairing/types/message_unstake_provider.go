package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnstakeProvider = "unstake_provider"

var _ sdk.Msg = &MsgUnstakeProvider{}

func NewMsgUnstakeProvider(creator, chainID, validator string) *MsgUnstakeProvider {
	return &MsgUnstakeProvider{
		Creator:   creator,
		ChainID:   chainID,
		Validator: validator,
	}
}

func (msg *MsgUnstakeProvider) Route() string {
	return RouterKey
}

func (msg *MsgUnstakeProvider) Type() string {
	return TypeMsgUnstakeProvider
}

func (msg *MsgUnstakeProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnstakeProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}
