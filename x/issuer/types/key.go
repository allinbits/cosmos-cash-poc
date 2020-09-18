package types

var (
	IssuersKey = []byte{0x51} // prefix for each key to a Issuer
	TokensKey  = []byte{0x53} // prefix for each key to a Token
)

const (
	// ModuleName is the name of the module
	ModuleName = "issuer"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)
