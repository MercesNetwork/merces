package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/MercesNetwork/merces/testutil/keeper"
	"github.com/MercesNetwork/merces/testutil/nullify"
	"github.com/MercesNetwork/merces/x/planet/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDNSRegistryQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PlanetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDNSRegistry(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDNSRegistryRequest
		response *types.QueryGetDNSRegistryResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDNSRegistryRequest{
				Domain: msgs[0].Domain,
			},
			response: &types.QueryGetDNSRegistryResponse{DNSRegistry: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDNSRegistryRequest{
				Domain: msgs[1].Domain,
			},
			response: &types.QueryGetDNSRegistryResponse{DNSRegistry: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDNSRegistryRequest{
				Domain: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.DNSRegistry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDNSRegistryQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PlanetKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDNSRegistry(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDNSRegistryRequest {
		return &types.QueryAllDNSRegistryRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DNSRegistryAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DNSRegistry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DNSRegistry),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DNSRegistryAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.DNSRegistry), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.DNSRegistry),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DNSRegistryAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.DNSRegistry),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DNSRegistryAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
