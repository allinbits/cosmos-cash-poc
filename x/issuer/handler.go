package issuer

import (
	"fmt"
	"strconv"

	"github.com/allinbits/cosmos-cash-poa/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateIssuer:
			return handleMsgCreateIssuer(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateIssuer(ctx sdk.Context, msg types.MsgCreateIssuer, k keeper.Keeper) (*sdk.Result, error) {
	if _, found := k.GetIssuer(ctx, msg.Address); found {
		return nil, nil
	}

	issuer := types.NewIssuer(
		msg.Name,
		msg.Address,
		msg.Token,
	)

	amount, err := strconv.Atoi(msg.Amount)
	if err != nil {
		return nil, nil
	}

	k.SetIssuer(ctx, msg.Address, issuer)
	checkIssuer, found := k.GetIssuer(ctx, msg.Address)
	if !found {
		return nil, nil
	}

	k.CoinKeeper.SetCoins(ctx, checkIssuer.Address, sdk.NewCoins(sdk.NewInt64Coin(msg.Token, int64(amount))))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateIssuer,
			sdk.NewAttribute(types.AttributeKeyIssuerAddress, msg.Address.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

//func handleMsgIssuerBurnToken(ctx sdk.Context, msg msg.MsgCreateToken, k keeper.Keeper) (*sdk.Result, error) {
//}

//func handleMsgIssuerMintToken(ctx sdk.Context, msg msg.MsgCreateToken, k keeper.Keeper) (*sdk.Result, error) {
//}
