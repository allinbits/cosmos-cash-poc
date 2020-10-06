package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IssuerState string

const (
	ACCEPTED IssuerState = "accepted"
	FROZEN   IssuerState = "frozen"
)

type Issuer struct {
	Name    string         `json:"name" yaml:"name"`
	Token   string         `json:"token" yaml:"token"`
	Fee     uint16         `json:"fee" yaml:"fee"`
	State   IssuerState    `json:"state" yaml:"state"`
	Address sdk.AccAddress `json:"address" yaml:"address"`
}

// NewIssuer - initialize a new issuer
func NewIssuer(name string, address sdk.AccAddress, token string) Issuer {
	return Issuer{
		Name:    name,
		Token:   token,
		Fee:     1,
		State:   ACCEPTED,
		Address: address,
	}
}
