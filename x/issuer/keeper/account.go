package keeper

import (
	"github.com/allinbits/cosmos-cash-poc/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetAccount(ctx sdk.Context, key []byte, account types.Account) {
	k.Set(ctx, key, types.AccountsKey, account)
}

func (k Keeper) GetAccount(ctx sdk.Context, key []byte) (types.Account, bool) {
	val, found := k.Get(ctx, key, types.AccountsKey, k.UnmarshalAccount)
	return val.(types.Account), found
}

func (k Keeper) UnmarshalAccount(value []byte) (interface{}, bool) {
	account := types.Account{}
	err := k.cdc.UnmarshalBinaryBare(value, &account)
	if err != nil {
		return types.Account{}, false
	}

	if account.Address.Empty() {
		return types.Account{}, false
	}

	return account, true
}
