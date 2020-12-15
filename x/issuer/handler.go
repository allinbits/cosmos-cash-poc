package issuer

import (
	"fmt"
	"strconv"

	didkeeper "github.com/allinbits/cosmos-cash-poc/x/did/keeper"
	didtypes "github.com/allinbits/cosmos-cash-poc/x/did/types"
	"github.com/allinbits/cosmos-cash-poc/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash-poc/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper, idk didkeeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateIssuer:
			return handleMsgCreateIssuer(ctx, msg, k, idk)
		case types.MsgBurnToken:
			return handleMsgBurnToken(ctx, msg, k)
		case types.MsgMintToken:
			return handleMsgMintToken(ctx, msg, k)
		case types.MsgFreezeToken:
			return handleMsgFreezeToken(ctx, msg, k)
		case types.MsgUnfreezeToken:
			return handleMsgUnfreezeToken(ctx, msg, k)
		case types.MsgWithdrawToken:
			return handleMsgWithdrawToken(ctx, msg, k)
		case types.MsgFreezeAccount:
			return handleMsgFreezeAccount(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateIssuer(ctx sdk.Context, msg types.MsgCreateIssuer, k keeper.Keeper, idk didkeeper.Keeper) (*sdk.Result, error) {
	// TODO: check if regulator
	document, found := idk.GetDidDocument(ctx, []byte(didtypes.DidIdentifer+msg.Owner.String()))
	if !found {
		return nil, fmt.Errorf("identity not found")
	}
	cred, found := idk.GetVerifiableCredential(ctx, []byte(document.Service[0].ID))
	if !found {
		return nil, fmt.Errorf("creds not found")
	}

	if cred.CredentialSubject.Role != "Issuer" {
		return nil, fmt.Errorf("role incorrect")
	}

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
			sdk.NewAttribute(types.AttributeKeyIssuerAmount, msg.Amount),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgBurnToken(ctx sdk.Context, msg types.MsgBurnToken, k keeper.Keeper) (*sdk.Result, error) {
	issuer, found := k.GetIssuer(ctx, msg.Issuer)
	if !found {
		return nil, nil
	}

	amount, err := strconv.Atoi(msg.Amount)
	if err != nil {
		return nil, nil
	}

	k.CoinKeeper.SubtractCoins(ctx, issuer.Address, sdk.NewCoins(sdk.NewInt64Coin(msg.Token, int64(amount))))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurnToken,
			sdk.NewAttribute(types.AttributeKeyBurnerAddress, msg.Issuer.String()),
			sdk.NewAttribute(types.AttributeKeyBurnerAmount, msg.Amount),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgMintToken(ctx sdk.Context, msg types.MsgMintToken, k keeper.Keeper) (*sdk.Result, error) {
	issuer, found := k.GetIssuer(ctx, msg.Issuer)
	if !found {
		return nil, nil
	}

	amount, err := strconv.Atoi(msg.Amount)
	if err != nil {
		return nil, nil
	}

	k.CoinKeeper.AddCoins(ctx, issuer.Address, sdk.NewCoins(sdk.NewInt64Coin(msg.Token, int64(amount))))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintToken,
			sdk.NewAttribute(types.AttributeKeyMinterAddress, msg.Issuer.String()),
			sdk.NewAttribute(types.AttributeKeyMinterAmount, msg.Amount),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Issuer.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgFreezeToken(ctx sdk.Context, msg types.MsgFreezeToken, k keeper.Keeper) (*sdk.Result, error) {
	issuer, found := k.GetIssuer(ctx, msg.Issuer)
	if !found {
		return nil, nil
	}

	issuer.State = types.FROZEN
	k.SetIssuer(ctx, msg.Issuer, issuer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFreezeToken,
			sdk.NewAttribute(types.AttributeKeyIssuerAddress, msg.Issuer.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgUnfreezeToken(ctx sdk.Context, msg types.MsgUnfreezeToken, k keeper.Keeper) (*sdk.Result, error) {
	issuer, found := k.GetIssuer(ctx, msg.Issuer)
	if !found {
		return nil, nil
	}

	issuer.State = types.ACCEPTED
	k.SetIssuer(ctx, msg.Issuer, issuer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnfreezeToken,
			sdk.NewAttribute(types.AttributeKeyIssuerAddress, msg.Issuer.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgWithdrawToken(ctx sdk.Context, msg types.MsgWithdrawToken, k keeper.Keeper) (*sdk.Result, error) {
	// TODO: have some check here

	amount, err := strconv.Atoi(msg.Amount)
	if err != nil {
		return nil, nil
	}

	k.SupplyKeeper.SendCoinsFromAccountToModule(ctx, msg.Owner, types.ModuleName, sdk.NewCoins(sdk.NewInt64Coin(msg.Token, int64(amount))))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawToken,
			sdk.NewAttribute(types.AttributeKeyIssuerAddress, msg.Owner.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgFreezeAccount(ctx sdk.Context, msg types.MsgFreezeAccount, k keeper.Keeper) (*sdk.Result, error) {
	account := types.NewAccount(
		msg.Account,
		types.FROZENACCOUNT,
	)

	k.SetAccount(ctx, msg.Account, account)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawToken,
			sdk.NewAttribute(types.AttributeKeyIssuerAddress, msg.Issuer.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
