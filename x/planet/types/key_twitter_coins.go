package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TwitterCoinsKeyPrefix is the prefix to retrieve all TwitterCoins
	TwitterCoinsKeyPrefix = "TwitterCoins/value/"
)

// TwitterCoinsKey returns the store key to retrieve a TwitterCoins from the index fields
func TwitterCoinsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
