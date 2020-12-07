package types

var (
	RegualtorKey = []byte{0x61} // prefix for each key to a Regualtor
)

const (
	// ModuleName is the name of the module
	ModuleName = "regulator"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)
