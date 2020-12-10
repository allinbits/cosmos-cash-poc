package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetDidDocument(ctx sdk.Context, key []byte, document types.DidDocument) {
	k.Set(ctx, key, types.DidDocumentKey, document)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, key []byte) (types.DidDocument, bool) {
	val, found := k.Get(ctx, key, types.DidDocumentKey, k.UnmarshalDidDocument)
	return val.(types.DidDocument), found
}

func (k Keeper) UnmarshalDidDocument(value []byte) (interface{}, bool) {
	document := types.DidDocument{}
	err := k.cdc.UnmarshalBinaryBare(value, &document)
	if err != nil {
		return types.DidDocument{}, false
	}

	if document.Context == "" {
		return types.DidDocument{}, false
	}

	return document, true
}
