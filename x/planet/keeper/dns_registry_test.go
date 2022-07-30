package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/MercesToken/planet/testutil/keeper"
	"github.com/MercesToken/planet/testutil/nullify"
	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDNSRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DNSRegistry {
	items := make([]types.DNSRegistry, n)
	for i := range items {
		items[i].Domain = strconv.Itoa(i)

		keeper.SetDNSRegistry(ctx, items[i])
	}
	return items
}

func TestDNSRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.PlanetKeeper(t)
	items := createNDNSRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDNSRegistry(ctx,
			item.Domain,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDNSRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.PlanetKeeper(t)
	items := createNDNSRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDNSRegistry(ctx,
			item.Domain,
		)
		_, found := keeper.GetDNSRegistry(ctx,
			item.Domain,
		)
		require.False(t, found)
	}
}

func TestDNSRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.PlanetKeeper(t)
	items := createNDNSRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDNSRegistry(ctx)),
	)
}
