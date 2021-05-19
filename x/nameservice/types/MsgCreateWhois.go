package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuyName{}

type MsgBuyName struct {
  Buyer sdk.AccAddress `json:"buyer" yaml:"buyer"`
  Name string  `json:"name" yaml:"name"`
  Bid sdk.Coins `json:"bid" yaml:"bid"`
}

func NewMsgBuyName(buyer sdk.AccAddress, name string, bid sdk.Coins) MsgBuyName {
  return MsgBuyName{
		Buyer: buyer,
    Name: name,
    Bid: bid,
	}
}

func (msg MsgBuyName) Route() string {
  return RouterKey
}

func (msg MsgBuyName) Type() string {
  return "BuyName"
}

func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Buyer)}
}

func (msg MsgBuyName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyName) ValidateBasic() error {
	if msg.Buyer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}
