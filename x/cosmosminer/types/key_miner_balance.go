package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MinerBalanceKeyPrefix is the prefix to retrieve all MinerBalance
	MinerBalanceKeyPrefix = "MinerBalance/value/"
)

// MinerBalanceKey returns the store key to retrieve a MinerBalance from the index fields
func MinerBalanceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
