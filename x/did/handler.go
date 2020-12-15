package did

import (
	"fmt"
	"strings"

	"github.com/allinbits/cosmos-cash-poc/x/did/keeper"
	"github.com/allinbits/cosmos-cash-poc/x/did/types"
	issuerkeeper "github.com/allinbits/cosmos-cash-poc/x/issuer/keeper"
	regulatorkeeper "github.com/allinbits/cosmos-cash-poc/x/regulator/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the did type messages
func NewHandler(k keeper.Keeper, rk regulatorkeeper.Keeper, ik issuerkeeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateDidDocument:
			return handleMsgCreateDidDocument(ctx, msg, k)
		case types.MsgCreateVerifiableCredential:
			return handleMsgCreateVerifiableCredential(ctx, msg, k, rk, ik)
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

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreateVerifiableCredential(ctx sdk.Context, msg types.MsgCreateVerifiableCredential, k keeper.Keeper, rk regulatorkeeper.Keeper, ik issuerkeeper.Keeper) (*sdk.Result, error) {
	// TODO: check if issuer exists
	document, found := k.GetDidDocument(ctx, []byte(msg.DidUrl))
	if !found {
		return nil, fmt.Errorf("no did found for user: %s", msg.DidUrl)
	}

	issuerAddress, _ := sdk.AccAddressFromBech32(msg.Issuer)

	var cred types.CredentialSubject

	identiferAddress := strings.Split(msg.DidUrl, ":")
	if msg.Issuer == identiferAddress[2] {
		cred = types.NewCredentialSubject("Regulator", true)
	} else {
		if _, found := rk.GetRegulator(ctx, issuerAddress); found {
			cred = types.NewCredentialSubject("Issuer", true)
		}
	}

	if _, found := ik.GetIssuer(ctx, issuerAddress); found {
		cred = types.NewCredentialSubject("User", true)
	}

	storeValue := msg.Issuer + ":" + msg.DidUrl

	verifiableCredential := types.NewVerifiableCredential(
		msg.Context,
		storeValue,
		msg.VcType,
		msg.Issuer,
		cred,
		msg.Proof,
	)

	service := types.NewService(storeValue, "role", "cash-bc")

	//TODO: check if serivce id exists

	document.Service = append(document.Service, service)

	k.SetVerifiableCredential(ctx, []byte(storeValue), verifiableCredential)
	k.SetDidDocument(ctx, []byte(document.ID), document)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
