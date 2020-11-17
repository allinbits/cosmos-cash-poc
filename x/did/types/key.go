package types

var (
	DidDocumentKey = []byte{0x61} // prefix for each key to a DidDocument
)

const (
	// ModuleName is the name of the module
	ModuleName = "did"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName

	// did identifer prefix
	DidIdentifer = "did:cosmos:"

	// w3c spec context definition
	Context = "https://www.w3.org/ns/did/v1"
)
