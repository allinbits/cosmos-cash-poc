package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgMintToken struct {
	Token  string         `json:"token"`
	Amount string         `json:"amount"`
	Issuer sdk.AccAddress `json:"issuer"`
}

func NewMsgMintToken(token string, amount string, issuer sdk.AccAddress) MsgMintToken {
	return MsgMintToken{
		Token:  token,
		Amount: amount,
		Issuer: issuer,
	}
}

// Route should return the name of the module
func (msg MsgMintToken) Route() string { return RouterKey }

// Type should return the action
func (msg MsgMintToken) Type() string { return "mint_token" }

// ValidateBasic runs stateless checks on the message
func (msg MsgMintToken) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgMintToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

// GetSignBytes encodes the message for signing
func (msg MsgMintToken) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgMintToken{}, "issuer/MsgMintToken", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
