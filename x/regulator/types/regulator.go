package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DidDocument is the data model defined by w3c
type Regulator struct {
	Address sdk.AccAddress `json:"address"`
}

func NewRegulator(address sdk.AccAddress) Regulator {
	return Regulator{
		Address: address,
	}
}
