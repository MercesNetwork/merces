package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DNSRegistryReverseKeyPrefix is the prefix to retrieve all DNSRegistryReverse
	DNSRegistryReverseKeyPrefix = "DNSRegistryReverse/value/"
)

// DNSRegistryReverseKey returns the store key to retrieve a DNSRegistryReverse from the index fields
func DNSRegistryReverseKey(
	publicKey string,
) []byte {
	var key []byte

	publicKeyBytes := []byte(publicKey)
	key = append(key, publicKeyBytes...)
	key = append(key, []byte("/")...)

	return key
}
