package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConflictVoteCommit = "conflict_vote_commit"

var _ sdk.Msg = &MsgConflictVoteCommit{}

func NewMsgConflictVoteCommit(creator, voteID string, hash []byte) *MsgConflictVoteCommit {
	return &MsgConflictVoteCommit{
		Creator: creator,
		VoteID:  voteID,
		Hash:    hash,
	}
}

func (msg *MsgConflictVoteCommit) Route() string {
	return RouterKey
}

func (msg *MsgConflictVoteCommit) Type() string {
	return TypeMsgConflictVoteCommit
}

func (msg *MsgConflictVoteCommit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
