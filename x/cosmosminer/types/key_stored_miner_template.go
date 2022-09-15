package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredMinerTemplateKeyPrefix is the prefix to retrieve all StoredMinerTemplate
	StoredMinerTemplateKeyPrefix = "StoredMinerTemplate/value/"
)

// StoredMinerTemplateKey returns the store key to retrieve a StoredMinerTemplate from the index fields
func StoredMinerTemplateKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
