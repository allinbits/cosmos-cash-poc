package did

import (
	"fmt"

	"github.com/allinbits/cosmos-cash-poa/x/did/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the did type messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateDidDocument:
			return handleMsgCreateDidDocument(ctx, msg, k)
		case types.MsgCreateVerifiableCredential:
			return handleMsgCreateVerifiableCredential(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateDidDocument(ctx sdk.Context, msg types.MsgCreateDidDocument, k keeper.Keeper) (*sdk.Result, error) {
	// TODO: check if document exists
	didDocument := types.NewDidDocument(
		msg.Context,
		msg.ID,
		msg.Authentication,
		msg.Services,
	)

	k.SetDidDocument(ctx, []byte(msg.ID), didDocument)

	//document, _ := k.GetDidDocument(ctx, []byte(msg.ID))
	//fmt.Println(document)

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

func handleMsgCreateVerifiableCredential(ctx sdk.Context, msg types.MsgCreateVerifiableCredential, k keeper.Keeper) (*sdk.Result, error) {
	// TODO: check if issuer exists
	verifiableCredential := types.NewVerifiableCredential(
		msg.Context,
		msg.ID,
		msg.VcType,
		msg.Issuer,
		msg.Proof,
	)

	k.SetVerifiableCredential(ctx, []byte(msg.ID), verifiableCredential)

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
