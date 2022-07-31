package planet_test

import (
	"testing"

	keepertest "github.com/MercesNetwork/merces/testutil/keeper"
	"github.com/MercesNetwork/merces/testutil/nullify"
	"github.com/MercesNetwork/merces/x/planet"
	"github.com/MercesNetwork/merces/x/planet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlanetKeeper(t)
	planet.InitGenesis(ctx, *k, genesisState)
	got := planet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TwitterCoinsList, got.TwitterCoinsList)
	require.ElementsMatch(t, genesisState.DNSRegistryList, got.DNSRegistryList)
	require.ElementsMatch(t, genesisState.DNSRegistryReverseList, got.DNSRegistryReverseList)
	// this line is used by starport scaffolding # genesis/test/assert
}
