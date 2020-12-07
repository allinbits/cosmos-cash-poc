package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/regulator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRegualtor(ctx sdk.Context, key []byte, document types.Regualtor) {
	k.Set(ctx, key, types.RegualtorKey, document)
}

func (k Keeper) GetRegualtor(ctx sdk.Context, key []byte) (types.Regualtor, bool) {
	val, found := k.Get(ctx, key, types.RegualtorKey, k.UnmarshalRegualtor)
	return val.(types.Regualtor), found
}

func (k Keeper) UnmarshalRegualtor(value []byte) (interface{}, bool) {
	document := types.Regualtor{}
	err := k.cdc.UnmarshalBinaryBare(value, &document)
	if err != nil {
		return types.Regualtor{}, false
	}

	return document, true
}
