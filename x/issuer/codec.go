package issuer

import (
	"github.com/allinbits/cosmos-cash-poc/x/issuer/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.MsgCreateIssuer{}, "issuer/MsgCreateIssuer", nil)
	cdc.RegisterConcrete(types.MsgMintToken{}, "issuer/MsgMintToken", nil)
	cdc.RegisterConcrete(types.MsgBurnToken{}, "issuer/MsgBurnToken", nil)
	cdc.RegisterConcrete(types.MsgFreezeToken{}, "issuer/MsgFreezeToken", nil)
	cdc.RegisterConcrete(types.MsgUnfreezeToken{}, "issuer/MsgUnfreezeToken", nil)
	cdc.RegisterConcrete(types.MsgFreezeAccount{}, "issuer/MsgreezeToken", nil)
	cdc.RegisterConcrete(types.MsgWithdrawToken{}, "issuer/MsgWithdrawToken", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
