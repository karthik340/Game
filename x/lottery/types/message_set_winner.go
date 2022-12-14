package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetWinner = "set_winner"

var _ sdk.Msg = &MsgSetWinner{}

func NewMsgSetWinner(validator string, winner string) *MsgSetWinner {
	return &MsgSetWinner{
		Validator: validator,
		Winner:    winner,
	}
}

func (msg *MsgSetWinner) Route() string {
	return RouterKey
}

func (msg *MsgSetWinner) Type() string {
	return TypeMsgSetWinner
}

func (msg *MsgSetWinner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetWinner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetWinner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Validator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
