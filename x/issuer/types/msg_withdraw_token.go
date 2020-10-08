package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgWithdrawToken struct {
	Token  string         `json:"token"`
	Amount string         `json:"amount"`
	Owner  sdk.AccAddress `json:"owner"`
}

func NewMsgWithdrawToken(token string, amount string, owner sdk.AccAddress) MsgWithdrawToken {
	return MsgWithdrawToken{
		Token:  token,
		Amount: amount,
		Owner:  owner,
	}
}

// Route should return the name of the module
func (msg MsgWithdrawToken) Route() string { return RouterKey }

// Type should return the action
func (msg MsgWithdrawToken) Type() string { return "burn_token" }

// ValidateBasic runs stateless checks on the message
func (msg MsgWithdrawToken) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgWithdrawToken) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgWithdrawToken) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgWithdrawToken{}, "issuer/MsgWithdrawToken", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
