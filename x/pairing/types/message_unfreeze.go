package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnfreeze = "unfreeze"

var _ sdk.Msg = &MsgUnfreezeProvider{}

func NewMsgUnfreeze(creator string, chainIds []string) *MsgUnfreezeProvider {
	return &MsgUnfreezeProvider{
		Creator:  creator,
		ChainIds: chainIds,
	}
}

func (msg *MsgUnfreezeProvider) Route() string {
	return RouterKey
}

func (msg *MsgUnfreezeProvider) Type() string {
	return TypeMsgUnfreeze
}

func (msg *MsgUnfreezeProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnfreezeProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
