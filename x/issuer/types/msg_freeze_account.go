package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgFreezeAccount struct {
	Account sdk.AccAddress `json:"account"`
	Issuer  sdk.AccAddress `json:"issuer"`
}

func NewMsgFreezeAccount(account sdk.AccAddress, issuer sdk.AccAddress) MsgFreezeAccount {
	return MsgFreezeAccount{
		Account: account,
		Issuer:  issuer,
	}
}

// Route should return the name of the module
func (msg MsgFreezeAccount) Route() string { return RouterKey }

// Type should return the action
func (msg MsgFreezeAccount) Type() string { return "freeze_account" }

// ValidateBasic runs stateless checks on the message
func (msg MsgFreezeAccount) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgFreezeAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Issuer}
}

// GetSignBytes encodes the message for signing
func (msg MsgFreezeAccount) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgFreezeAccount{}, "issuer/MsgFreezeAccount", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
