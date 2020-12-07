package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Vote messages
type MsgCreateRegualtor struct {
	Address sdk.AccAddress `json:"address"`
	Owner   sdk.AccAddress `json:"owner"`
}

func NewMsgCreateRegualtor(address sdk.AccAddress, owner sdk.AccAddress) MsgCreateRegualtor {
	return MsgCreateRegualtor{
		Address: address,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgCreateRegualtor) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateRegualtor) Type() string { return "create_regulator" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateRegualtor) ValidateBasic() error {
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgCreateRegualtor) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRegualtor) GetSignBytes() []byte {
	ModuleCdc := codec.New()
	ModuleCdc.RegisterConcrete(MsgCreateRegualtor{}, "regulator/MsgCreateRegualtor", nil)
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
