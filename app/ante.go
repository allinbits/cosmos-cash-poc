package app

import (
	didkeeper "github.com/allinbits/cosmos-cash-poa/x/did/keeper"
	issuerante "github.com/allinbits/cosmos-cash-poa/x/issuer/ante"
	issuerkeeper "github.com/allinbits/cosmos-cash-poa/x/issuer/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	supplykeeper "github.com/cosmos/cosmos-sdk/x/supply"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(
	ak authkeeper.AccountKeeper,
	supplyKeeper supplykeeper.Keeper,
	ik issuerkeeper.Keeper,
	dk didkeeper.Keeper,
	sigGasConsumer authante.SignatureVerificationGasConsumer,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		authante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		authante.NewMempoolFeeDecorator(),
		authante.NewValidateBasicDecorator(),
		authante.NewValidateMemoDecorator(ak),
		authante.NewConsumeGasForTxSizeDecorator(ak),
		authante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		authante.NewValidateSigCountDecorator(ak),
		authante.NewDeductFeeDecorator(ak, supplyKeeper),
		authante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		authante.NewSigVerificationDecorator(ak),
		authante.NewIncrementSequenceDecorator(ak), // innermost AnteDecorator
		issuerante.NewDeductIssuerFeeDecorator(ik, dk),
	)
}
