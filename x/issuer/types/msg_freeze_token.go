package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgFreezeToken struct {
	Token  string         `json:"token"`
	Issuer sdk.AccAddress `json:"issuer"`
}

func NewMsgFreezeToken(token string, issuer sdk.AccAddress) MsgFreezeToken {
	return MsgFreezeToken{
		Token:  token,
		Issuer: issuer,
	}
}

// Route should return the name of the module
func (msg MsgFreezeToken) Route() string { return RouterKey }

// Type should return the action
func (msg MsgFreezeToken) Type() string { return "burn_token" }

// ValidateBasic runs stateless checks on the message
func (msg MsgFreezeToken) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgFreezeToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

// GetSignBytes encodes the message for signing
func (msg MsgFreezeToken) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgFreezeToken{}, "issuer/MsgFreezeToken", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
