package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgCreateIssuer struct {
	Name    string         `json:"name"`
	Address sdk.AccAddress `json:"issuer"`
	Token   string         `json:"token"`
	Amount  string         `json:"amount"`
	Owner   sdk.AccAddress `json:"owner"`
}

func NewMsgCreateIssuer(name string, address sdk.AccAddress, token string, amount string, owner sdk.AccAddress) MsgCreateIssuer {
	return MsgCreateIssuer{
		Name:    name,
		Address: address,
		Token:   token,
		Amount:  amount,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgCreateIssuer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateIssuer) Type() string { return "create_issuer" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateIssuer) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgCreateIssuer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateIssuer) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgCreateIssuer{}, "issuer/MsgCreateIssuer", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
