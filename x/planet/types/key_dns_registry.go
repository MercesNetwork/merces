package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DNSRegistryKeyPrefix is the prefix to retrieve all DNSRegistry
	DNSRegistryKeyPrefix = "DNS/value/"
)

// DNSRegistryKey returns the store key to retrieve a DNSRegistry from the index fields
func DNSRegistryKey(
	domain string,
) []byte {
	var key []byte

	domainBytes := []byte(domain)
	key = append(key, domainBytes...)
	key = append(key, []byte("/")...)

	return key
}
