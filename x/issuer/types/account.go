package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountState string

const (
	APPROVED      AccountState = "approved"
	PENDING       AccountState = "pending"
	FROZENACCOUNT AccountState = "frozen"
)

type Account struct {
	State   AccountState   `json:"state" yaml:"state"`
	Address sdk.AccAddress `json:"address" yaml:"address"`
}

// NewAccount - initialize a new issuer
func NewAccount(address sdk.AccAddress, state AccountState) Account {
	return Account{
		State:   state,
		Address: address,
	}
}
