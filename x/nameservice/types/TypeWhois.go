package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

type Whois struct {
		Owner sdk.AccAddress `json:"owner" yaml:"owner"`
    Value string `json:"value" yaml:"value"`
    Price sdk.Coins `json:"price" yaml:"price"`
}

func newWhois() Whois {
	return Whois {
		Price: MinNamePrice,
	}
}
