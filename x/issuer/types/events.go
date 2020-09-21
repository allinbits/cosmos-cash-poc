package types

// issuer module event types
const (
	EventTypeCreateIssuer = "create-issuer"
	EventTypeMintToken    = "mint-token"
	EventTypeBurnToken    = "burn-token"

	AttributeKeyIssuerAddress = "issuer-address"
	AttributeKeyIssuerAmount  = "issuer-amount"
	AttributeKeyMinterAddress = "minter-address"
	AttributeKeyMinterAmount  = "minter-amount"
	AttributeKeyBurnerAddress = "minter-address"
	AttributeKeyBurnerAmount  = "burner-amount"

	AttributeValueCategory = ModuleName
)
