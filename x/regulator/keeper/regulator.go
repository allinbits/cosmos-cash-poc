package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/regulator/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRegulator(ctx sdk.Context, key []byte, reg types.Regulator) {
	k.Set(ctx, key, types.RegulatorKey, reg)
}

func (k Keeper) GetRegulator(ctx sdk.Context, key []byte) (types.Regulator, bool) {
	val, found := k.Get(ctx, key, types.RegulatorKey, k.UnmarshalRegulator)
	return val.(types.Regulator), found
}

func (k Keeper) UnmarshalRegulator(value []byte) (interface{}, bool) {
	reg := types.Regulator{}
	err := k.cdc.UnmarshalBinaryBare(value, &reg)
	if err != nil {
		return types.Regulator{}, false
	}

	return reg, true
}
