package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdate{}

func NewMsgUpdate(creator string, name string, value string, id uint64) *MsgUpdate {
	return &MsgUpdate{
		Creator: creator,
		Name:    name,
		Value:   value,
		Id:      id,
	}
}

func (msg *MsgUpdate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
