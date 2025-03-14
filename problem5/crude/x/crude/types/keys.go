package types

const (
	// ModuleName defines the module name
	ModuleName = "crude"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_crude"
)

var (
	ParamsKey = []byte("p_crude")

	ResourceKey = "Resource/value"

	ResourceCountKey = "Resource/count"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
