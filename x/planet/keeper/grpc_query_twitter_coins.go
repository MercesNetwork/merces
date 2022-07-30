package keeper

import (
	"context"

	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TwitterCoinsAll(c context.Context, req *types.QueryAllTwitterCoinsRequest) (*types.QueryAllTwitterCoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var twitterCoinss []types.TwitterCoins
	ctx := sdk.UnwrapSDKContext(c)

	// store := ctx.KVStore(k.storeKey)
	// twitterCoinsStore := prefix.NewStore(store, types.KeyPrefix(types.TwitterCoinsKeyPrefix))
	twitterCoinsStore := k.getTwitterAccountScopedStore(ctx, req.GetUsername())

	pageRes, err := query.Paginate(twitterCoinsStore, req.Pagination, func(key []byte, value []byte) error {
		var twitterCoins types.TwitterCoins
		if err := k.cdc.Unmarshal(value, &twitterCoins); err != nil {
			return err
		}

		twitterCoinss = append(twitterCoinss, twitterCoins)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTwitterCoinsResponse{TwitterCoins: twitterCoinss, Pagination: pageRes}, nil
}

func (k Keeper) TwitterCoins(c context.Context, req *types.QueryGetTwitterCoinsRequest) (*types.QueryGetTwitterCoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTwitterCoins(
		ctx,
		req.Username,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTwitterCoinsResponse{TwitterCoins: val}, nil
}
