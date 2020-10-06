package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgBurnToken struct {
	Token  string         `json:"token"`
	Amount string         `json:"amount"`
	Issuer sdk.AccAddress `json:"issuer"`
}

func NewMsgBurnToken(token string, amount string, issuer sdk.AccAddress) MsgBurnToken {
	return MsgBurnToken{
		Token:  token,
		Amount: amount,
		Issuer: issuer,
	}
}

// Route should return the name of the module
func (msg MsgBurnToken) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBurnToken) Type() string { return "burn_token" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBurnToken) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgBurnToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

// GetSignBytes encodes the message for signing
func (msg MsgBurnToken) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgBurnToken{}, "issuer/MsgBurnToken", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
