package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddKeys = "add_keys"

var _ sdk.Msg = &MsgAddKeys{}

func NewMsgAddKeys(creator, projectID string, projectKeys []ProjectKey) *MsgAddKeys {
	return &MsgAddKeys{
		Creator:     creator,
		Project:     projectID,
		ProjectKeys: projectKeys,
	}
}

func (msg *MsgAddKeys) Route() string {
	return RouterKey
}

func (msg *MsgAddKeys) Type() string {
	return TypeMsgAddKeys
}

func (msg *MsgAddKeys) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	for _, key := range msg.ProjectKeys {
		_, err := sdk.AccAddressFromBech32(key.Key)
		if err != nil {
			return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid project key address (%s)", err)
		}
	}

	return nil
}
