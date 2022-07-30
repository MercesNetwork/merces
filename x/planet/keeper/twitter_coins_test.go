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

func createNTwitterCoins(keeper *keeper.Keeper, ctx sdk.Context, username string, n int) []types.TwitterCoins {
	items := make([]types.TwitterCoins, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTwitterCoins(ctx, username, items[i])
	}
	return items
}

func TestTwitterCoinsGet(t *testing.T) {
	username := "merces"
	keeper, ctx := keepertest.PlanetKeeper(t)
	items := createNTwitterCoins(keeper, ctx, username, 10)
	for _, item := range items {
		rst, found := keeper.GetTwitterCoins(ctx,
			username,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTwitterCoinsRemove(t *testing.T) {
	username := "merces"
	keeper, ctx := keepertest.PlanetKeeper(t)
	items := createNTwitterCoins(keeper, ctx, username, 10)
	for _, item := range items {
		keeper.RemoveTwitterCoins(ctx,
			username,
			item.Index,
		)
		_, found := keeper.GetTwitterCoins(ctx,
			username,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTwitterCoinsGetAll(t *testing.T) {
	username := "merces"
	keeper, ctx := keepertest.PlanetKeeper(t)
	items := createNTwitterCoins(keeper, ctx, username, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTwitterCoins(ctx, username)),
	)
}
