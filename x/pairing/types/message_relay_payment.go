package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRelayPayment = "relay_payment"

var _ sdk.Msg = &MsgRelayPayment{}

func NewMsgRelayPayment(creator string, relays []*RelaySession, description string, latestBlocks []*LatestBlockReport) *MsgRelayPayment {
	return &MsgRelayPayment{
		Creator:            creator,
		Relays:             relays,
		DescriptionString:  description,
		LatestBlockReports: latestBlocks,
	}
}

func (msg *MsgRelayPayment) Route() string {
	return RouterKey
}

func (msg *MsgRelayPayment) Type() string {
	return TypeMsgRelayPayment
}

func (msg *MsgRelayPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
