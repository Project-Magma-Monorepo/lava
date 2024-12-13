package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConflictVoteReveal = "conflict_vote_reveal"

var _ sdk.Msg = &MsgConflictVoteReveal{}

func NewMsgConflictVoteReveal(creator, voteID string, nonce int64, hash []byte) *MsgConflictVoteReveal {
	return &MsgConflictVoteReveal{
		Creator: creator,
		VoteID:  voteID,
		Nonce:   nonce,
		Hash:    hash,
	}
}

func (msg *MsgConflictVoteReveal) Route() string {
	return RouterKey
}

func (msg *MsgConflictVoteReveal) Type() string {
	return TypeMsgConflictVoteReveal
}

func (msg *MsgConflictVoteReveal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
