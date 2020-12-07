package regulator

import (
	"fmt"

	"github.com/allinbits/cosmos-cash-poa/x/regulator/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/regulator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the regulator type messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		// 	return handleMsg<Action>(ctx, k, msg)
		case types.MsgCreateRegualtor:
			return handleMsgCreateRegualtor(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handle<Action> does x
func handleMsgCreateRegualtor(ctx sdk.Context, msg types.MsgCreateRegualtor, k keeper.Keeper) (*sdk.Result, error) {
	// Only trusted regualtors can create other regualors
	_, found := k.GetRegualtor(ctx, []byte(msg.Owner))
	if !found {
		return nil, nil //fmt.Error("regualator not found")
	}

	regualtor := types.NewRegualtor(
		msg.Address,
	)

	k.SetRegualtor(ctx, []byte(msg.Address), regualtor)

	/*	ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.EventTypeCreateIssuer,
				sdk.NewAttribute(types.AttributeKeyIssuerAddress, msg.Address.String()),
				sdk.NewAttribute(types.AttributeKeyIssuerAmount, msg.Amount),
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
			),
		})
	/*/
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
