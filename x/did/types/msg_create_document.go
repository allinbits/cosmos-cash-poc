package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgCreateDidDocument struct {
	Context        string         `json:"context"`
	ID             string         `json:"id"`
	Authentication PubKeys        `json:"authentication"`
	Services       Services       `json:"service,omitempty"`
	Owner          sdk.AccAddress `json:"owner"`
}

func NewMsgCreateDidDocument(context string, id string, authentication PubKeys, services Services, owner sdk.AccAddress) MsgCreateDidDocument {
	return MsgCreateDidDocument{
		Context:        context,
		ID:             id,
		Authentication: authentication,
		Services:       services,
		Owner:          owner,
	}
}

// Route should return the name of the module
func (msg MsgCreateDidDocument) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateDidDocument) Type() string { return "create_did_document" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDidDocument) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgCreateDidDocument) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateDidDocument) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgCreateDidDocument{}, "issuer/MsgCreateDidDocument", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
