package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/poa/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetValidator(ctx sdk.Context, key string, validator types.Validator) {
	// set validator by name
	k.Set(ctx, []byte(key), types.ValidatorsKey, validator)

	// set validiator by address
	k.Set(ctx, validator.Address, types.ValidatorsByAddressKey, validator)
}

func (k Keeper) GetValidator(ctx sdk.Context, key string) (types.Validator, bool) {
	val, found := k.Get(ctx, []byte(key), types.ValidatorsKey, k.UnmarshalValidator)
	return val.(types.Validator), found
}

func (k Keeper) GetValidatorByAddress(ctx sdk.Context, valAddress sdk.ValAddress) (types.Validator, bool) {
	val, found := k.Get(ctx, valAddress, types.ValidatorsByAddressKey, k.UnmarshalValidator)
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

func (k Keeper) GetAllAcceptedValidators(ctx sdk.Context) (validators []types.Validator) {
	val := k.GetAll(ctx, types.ValidatorsKey, k.UnmarshalValidator)

	if len(val) == 1 {
		return append(validators, val[0].(types.Validator))

	}

	for _, value := range val {
		validator := value.(types.Validator)
		if validator.Accepted == true {
			validators = append(validators, validator)
		}
	}

	return validators
}
