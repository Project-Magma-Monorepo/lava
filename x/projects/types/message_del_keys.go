package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDelKeys = "del_keys"

var _ sdk.Msg = &MsgDelKeys{}

func NewMsgDelKeys(creator, projectID string, projectKeys []ProjectKey) *MsgDelKeys {
	return &MsgDelKeys{
		Creator:     creator,
		Project:     projectID,
		ProjectKeys: projectKeys,
	}
}

func (msg *MsgDelKeys) Route() string {
	return RouterKey
}

func (msg *MsgDelKeys) Type() string {
	return TypeMsgDelKeys
}

func (msg *MsgDelKeys) ValidateBasic() error {
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
