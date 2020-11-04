package ante

import (
	"fmt"
	"github.com/allinbits/cosmos-cash-poa/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
)

// NewAnteHandler returns an AnteHandler
func NewAnteHandler(ik keeper.Keeper) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewDeductIssuerFeeDecorator(ik),
	)
}

// DeductIssuerFeeDecorator deducts fees from the every send transaction
type DeductIssuerFeeDecorator struct {
	ik keeper.Keeper
}

func NewDeductIssuerFeeDecorator(ik keeper.Keeper) DeductIssuerFeeDecorator {
	return DeductIssuerFeeDecorator{
		ik: ik,
	}
}

func (difd DeductIssuerFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		if msg.Type() == "send" {
			sendMsg := msg.(bank.MsgSend)
			issuer, found := difd.ik.GetIssuerByToken(ctx, sendMsg.Amount[0].Denom)
			if issuer.State == types.FROZEN {
				return ctx, fmt.Errorf("token is frozen")
			}

			if found {
				account, found := difd.ik.GetAccount(ctx, sendMsg.FromAddress)
				if found {
					if account.State == types.FROZENACCOUNT {
						return ctx, fmt.Errorf("account is frozen")
					}
				}

				account, found = difd.ik.GetAccount(ctx, sendMsg.ToAddress)
				if found {
					if account.State == types.FROZENACCOUNT {
						return ctx, fmt.Errorf("account is frozen")
					}
				}

				difd.ik.CoinKeeper.SendCoins(ctx, sendMsg.FromAddress, issuer.Address, sdk.NewCoins(sdk.NewInt64Coin(issuer.Token, int64(issuer.Fee))))
			}
		}
	}

	return next(ctx, tx, simulate)
}
