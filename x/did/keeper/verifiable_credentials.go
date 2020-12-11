package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetVerifiableCredential(ctx sdk.Context, key []byte, document types.VerifiableCredential) {
	k.Set(ctx, key, types.VerifiableCredentialKey, document)
}

func (k Keeper) GetVerifiableCredential(ctx sdk.Context, key []byte) (types.VerifiableCredential, bool) {
	val, found := k.Get(ctx, key, types.VerifiableCredentialKey, k.UnmarshalVerifiableCredential)
	return val.(types.VerifiableCredential), found
}

func (k Keeper) UnmarshalVerifiableCredential(value []byte) (interface{}, bool) {
	document := types.VerifiableCredential{}
	err := k.cdc.UnmarshalBinaryBare(value, &document)
	if err != nil {
		return types.VerifiableCredential{}, false
	}

	if document.Context == "" {
		return types.VerifiableCredential{}, false
	}

	return document, true
}
