package nameservice

import (
	"fmt"

	"github.com/Cabemo/nameservice/x/nameservice/keeper"
	"github.com/Cabemo/nameservice/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgBuyName:
			return handleMsgBuyName(ctx, k, msg)
		case types.MsgSetName:
			return handleMsgSetName(ctx, k, msg)
		case types.MsgDeleteName:
			return handleMsgDeleteName(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
