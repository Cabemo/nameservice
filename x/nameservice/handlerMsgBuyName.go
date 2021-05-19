package nameservice

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/Cabemo/nameservice/x/nameservice/keeper"
	"github.com/Cabemo/nameservice/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, k keeper.Keeper, msg types.MsgBuyName) (*sdk.Result, error) {
	// Checks if the the bid price is greater than the price paid by the current owner
	if k.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough") // If not, throw an error
	}
	if k.HasOwner(ctx, msg.Name) {
		err := k.CoinKeeper.SendCoins(ctx, msg.Buyer, k.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := k.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return nil, err
		}
	}
	k.SetOwner(ctx, msg.Name, msg.Buyer)
	k.SetPrice(ctx, msg.Name, msg.Bid)
	return &sdk.Result{}, nil
}
