package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredMinerKeyPrefix is the prefix to retrieve all StoredMiner
	StoredMinerKeyPrefix = "StoredMiner/value/"
)

// StoredMinerKey returns the store key to retrieve a StoredMiner from the index fields
func StoredMinerKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
