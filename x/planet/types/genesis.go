package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TwitterCoinsList: []TwitterCoins{},
		DNSRegistryList:  []DNSRegistry{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in twitterCoins
	twitterCoinsIndexMap := make(map[string]struct{})

	for _, elem := range gs.TwitterCoinsList {
		index := string(TwitterCoinsKey(elem.Index))
		if _, ok := twitterCoinsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for twitterCoins")
		}
		twitterCoinsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in dNSRegistry
	dNSRegistryIndexMap := make(map[string]struct{})

	for _, elem := range gs.DNSRegistryList {
		index := string(DNSRegistryKey(elem.Domain))
		if _, ok := dNSRegistryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for dNSRegistry")
		}
		dNSRegistryIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
