package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Issuer struct {
	Name       string         `json:"name" yaml:"name"`
	Token      string         `json:"token" yaml:"token"`
	Fee        uint16         `json:"fee" yaml:"fee"`
	IsAccepted bool           `json:"isaccepted" yaml:"isaccepted"`
	Address    sdk.AccAddress `json:"address" yaml:"address"`
}

// NewIssuer - initialize a new issuer
func NewIssuer(name string, address sdk.AccAddress, token string) Issuer {
	return Issuer{
		Name:       name,
		Token:      token,
		Fee:        1,
		IsAccepted: true,
		Address:    address,
	}
}
