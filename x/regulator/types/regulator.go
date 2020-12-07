package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DidDocument is the data model defined by w3c
type Regualtor struct {
	Address sdk.AccAddress `json:"address"`
}

func NewRegualtor(address sdk.AccAddress) Regualtor {
	return Regualtor{
		Address: address,
	}
}
