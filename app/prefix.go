package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	AccountAddressPrefix   = "cosmos"
	AccountPubKeyPrefix    = "cosmospub"
	ValidatorAddressPrefix = "cosmosvaloper"
	ValidatorPubKeyPrefix  = "cosmosvaloperpub"
	ConsNodeAddressPrefix  = "cosmosvalcons"
	ConsNodePubKeyPrefix   = "cosmosvalconspub"
)

func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}
