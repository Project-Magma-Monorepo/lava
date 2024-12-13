package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDelProject = "del_project"

var _ sdk.Msg = &MsgDelProject{}

func NewMsgDelProject(creator, name string) *MsgDelProject {
	return &MsgDelProject{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgDelProject) Route() string {
	return RouterKey
}

func (msg *MsgDelProject) Type() string {
	return TypeMsgDelProject
}

func (msg *MsgDelProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
