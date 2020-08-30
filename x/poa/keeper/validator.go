package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetValidator(ctx sdk.Context, key string, validator types.Validator) {
	k.Set(ctx, []byte(key), types.ValidatorsKey, validator)
}

func (k Keeper) GetValidator(ctx sdk.Context, key string) (types.Validator, bool) {
	val, found := k.Get(ctx, []byte(key), types.ValidatorsKey, k.UnmarshalValidator)
	return val.(types.Validator), found
}

func (k Keeper) UnmarshalValidator(value []byte) (interface{}, bool) {
	validator := types.Validator{}
	err := k.cdc.UnmarshalBinaryLengthPrefixed(value, &validator)
	if err != nil {
		return types.Validator{}, false
	}
	return validator, true
}

func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	val := k.GetAll(ctx, types.ValidatorsKey, k.UnmarshalValidator)

	// TODO: Make this nicer
	for _, value := range val {
		validators = append(validators, value.(types.Validator))
	}

	return validators
}
