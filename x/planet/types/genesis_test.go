package types_test

import (
	"testing"

	"github.com/MercesToken/planet/x/planet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				TwitterCoinsList: []types.TwitterCoins{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				DNSRegistryList: []types.DNSRegistry{
					{
						Domain: "0",
					},
					{
						Domain: "1",
					},
				},
				DNSRegistryReverseList: []types.DNSRegistryReverse{
					{
						PublicKey: "0",
					},
					{
						PublicKey: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated twitterCoins",
			genState: &types.GenesisState{
				TwitterCoinsList: []types.TwitterCoins{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated dNSRegistry",
			genState: &types.GenesisState{
				DNSRegistryList: []types.DNSRegistry{
					{
						Domain: "0",
					},
					{
						Domain: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated dNSRegistryReverse",
			genState: &types.GenesisState{
				DNSRegistryReverseList: []types.DNSRegistryReverse{
					{
						PublicKey: "0",
					},
					{
						PublicKey: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
