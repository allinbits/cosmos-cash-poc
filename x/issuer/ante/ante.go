package ante

import (
	"fmt"
	didkeeper "github.com/allinbits/cosmos-cash-poa/x/did/keeper"
	didtypes "github.com/allinbits/cosmos-cash-poa/x/did/types"
	"github.com/allinbits/cosmos-cash-poa/x/issuer/keeper"
	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
)

// NewAnteHandler returns an AnteHandler
func NewAnteHandler(ik keeper.Keeper, dk didkeeper.Keeper) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewDeductIssuerFeeDecorator(ik, dk),
	)
}

// DeductIssuerFeeDecorator deducts fees from the every send transaction
type DeductIssuerFeeDecorator struct {
	ik keeper.Keeper
	dk didkeeper.Keeper
}

func NewDeductIssuerFeeDecorator(ik keeper.Keeper, dk didkeeper.Keeper) DeductIssuerFeeDecorator {
	return DeductIssuerFeeDecorator{
		ik: ik,
		dk: dk,
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
				did, foundDid := difd.dk.GetDidDocument(ctx, []byte(didtypes.DidIdentifer+sendMsg.ToAddress.String()))
				if !foundDid {
					return ctx, fmt.Errorf("user has no identity")
				}

				cred, foundCred := difd.dk.GetVerifiableCredential(ctx, []byte(did.Service[0].ID))
				if !foundCred {
					return ctx, fmt.Errorf("user has no credentials")
				}

				role := cred.CredentialSubject.Role

				if role != "User" && role != "Issuer" && role != "Regulator" {
					return ctx, fmt.Errorf("user has incorrect role")
				}

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
