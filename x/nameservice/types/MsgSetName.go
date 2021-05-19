package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetName{}

type MsgSetName struct {
  Owner sdk.AccAddress `json:"owner" yaml:"owner"`
  Name string `json:"name" yaml:"name"`
  Value string `json:"value" yaml:"value"`
}

func NewMsgSetName(owner sdk.AccAddress, name string, value string) MsgSetName {
  return MsgSetName{
		Owner: owner,
    Name: name,
    Value: value,
	}
}

func (msg MsgSetName) Route() string {
  return RouterKey
}

func (msg MsgSetName) Type() string {
  return "SetName"
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg MsgSetName) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetName) ValidateBasic() error {
  if msg.Owner.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
  }
  if len(msg.Name) == 0 || len(msg.Value) == 0 {
    return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "name and/or value weren't defined")
  }
  return nil
}
