package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Verifiable Credential message
type MsgCreateVerifiableCredential struct {
	DidUrl            string            `json:"didurl"`
	Context           string            `json:"@context"`
	ID                string            `json:"id"`
	VcType            string            `json:"type"`
	Issuer            string            `json:"issuer"`
	CredentialSubject CredentialSubject `json:"credentialsubject"`
	Proof             Proof             `json:"proof"`
	Owner             sdk.AccAddress    `json:"owner"`
}

func NewMsgCreateVerifiableCredential(didurl string, context string, id string, vctype string, issuer string, proof Proof, owner sdk.AccAddress) MsgCreateVerifiableCredential {
	return MsgCreateVerifiableCredential{
		DidUrl:  didurl,
		Context: context,
		ID:      id,
		VcType:  vctype,
		Issuer:  issuer,
		Proof:   proof,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgCreateVerifiableCredential) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateVerifiableCredential) Type() string { return "create_verifiable_credential" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateVerifiableCredential) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgCreateVerifiableCredential) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateVerifiableCredential) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgCreateVerifiableCredential{}, "did/MsgCreateVerifiableCredential", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
