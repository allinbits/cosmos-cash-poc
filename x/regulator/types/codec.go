package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	//exported "github.com/cosmos/cosmos-sdk/x/auth/exported"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	//cdc.RegisterInterface((*exported.GenesisAccount)(nil), 1)

	cdc.RegisterConcrete(MsgCreateRegualtor{}, "regualtor/MsgCreateRegualtor", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
