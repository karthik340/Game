package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlaceBet = "place_bet"

var _ sdk.Msg = &MsgPlaceBet{}

func NewMsgPlaceBet(creator string, fee sdk.Coin, bet sdk.Coin) *MsgPlaceBet {
	return &MsgPlaceBet{
		Creator: creator,
		Fee:     fee,
		Bet:     bet,
	}
}

func (msg *MsgPlaceBet) Route() string {
	return RouterKey
}

func (msg *MsgPlaceBet) Type() string {
	return TypeMsgPlaceBet
}

func (msg *MsgPlaceBet) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPlaceBet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPlaceBet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
