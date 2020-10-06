package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgUnfreezeToken struct {
	Token  string         `json:"token"`
	Issuer sdk.AccAddress `json:"issuer"`
}

func NewMsgUnfreezeToken(token string, issuer sdk.AccAddress) MsgUnfreezeToken {
	return MsgUnfreezeToken{
		Token:  token,
		Issuer: issuer,
	}
}

// Route should return the name of the module
func (msg MsgUnfreezeToken) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnfreezeToken) Type() string { return "burn_token" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUnfreezeToken) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgUnfreezeToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnfreezeToken) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgUnfreezeToken{}, "issuer/MsgUnfreezeToken", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
