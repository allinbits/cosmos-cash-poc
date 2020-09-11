package poa

import (
	"fmt"

	"github.com/allinbits/cosmos-cash-poa/x/poa/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/poa/msg"
	"github.com/allinbits/cosmos-cash-poa/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateValidatorPOA:
			return handleMsgCreateValidatorPOA(ctx, msg, k)
		case types.MsgVoteValidator:
			return handleMsgVoteValidator(ctx, msg, k)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateValidatorPOA(ctx sdk.Context, msg msg.MsgCreateValidatorPOA, k keeper.Keeper) (*sdk.Result, error) {
	// check to see if the name has been registered before
	if _, found := k.GetValidator(ctx, msg.Name); found {
		return nil, nil
	}

	validator := types.NewValidator(
		msg.Name,
		msg.Address,
		msg.PubKey,
		stakingtypes.Description{"nil", "nil", "nil", "nil", "nil"},
	)

	k.SetValidator(ctx, msg.Name, validator)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			stakingtypes.EventTypeCreateValidator,
			sdk.NewAttribute(stakingtypes.AttributeKeyValidator, msg.Address.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgVoteValidator(ctx sdk.Context, msg msg.MsgVoteValidator, k keeper.Keeper) (*sdk.Result, error) {
	_, found := k.GetValidator(ctx, msg.Name)
	if !found {
		return nil, nil
	}

	vote := types.NewVote(
		msg.Voter,
		msg.Name,
		true,
	)

	k.SetVote(ctx, vote)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeVote,
			sdk.NewAttribute(stakingtypes.AttributeKeyValidator, msg.Voter.String()),
			sdk.NewAttribute(types.AttributeKeyCandidate, msg.Name),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
