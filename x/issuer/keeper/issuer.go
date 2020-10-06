package keeper

import (
	"github.com/allinbits/cosmos-cash-poa/x/issuer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetIssuer(ctx sdk.Context, key []byte, issuer types.Issuer) {
	k.Set(ctx, key, types.IssuersKey, issuer)

	k.Set(ctx, []byte(issuer.Token), types.TokensKey, issuer)
}

func (k Keeper) GetIssuer(ctx sdk.Context, key []byte) (types.Issuer, bool) {
	val, found := k.Get(ctx, key, types.IssuersKey, k.UnmarshalIssuer)
	return val.(types.Issuer), found
}

func (k Keeper) GetIssuerByToken(ctx sdk.Context, token string) (types.Issuer, bool) {
	val, found := k.Get(ctx, []byte(token), types.TokensKey, k.UnmarshalIssuer)
	return val.(types.Issuer), found
}

func (k Keeper) UnmarshalIssuer(value []byte) (interface{}, bool) {
	issuer := types.Issuer{}
	err := k.cdc.UnmarshalBinaryBare(value, &issuer)
	if err != nil {
		return types.Issuer{}, false
	}

	if issuer.Address.Empty() {
		return types.Issuer{}, false
	}

	return issuer, true
}
