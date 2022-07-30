package keeper

import (
	"context"

	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DNSRegistryAll(c context.Context, req *types.QueryAllDNSRegistryRequest) (*types.QueryAllDNSRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var dNSRegistrys []types.DNSRegistry
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	dNSRegistryStore := prefix.NewStore(store, types.KeyPrefix(types.DNSRegistryKeyPrefix))

	pageRes, err := query.Paginate(dNSRegistryStore, req.Pagination, func(key []byte, value []byte) error {
		var dNSRegistry types.DNSRegistry
		if err := k.cdc.Unmarshal(value, &dNSRegistry); err != nil {
			return err
		}

		dNSRegistrys = append(dNSRegistrys, dNSRegistry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDNSRegistryResponse{DNSRegistry: dNSRegistrys, Pagination: pageRes}, nil
}

func (k Keeper) DNSRegistry(c context.Context, req *types.QueryGetDNSRegistryRequest) (*types.QueryGetDNSRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDNSRegistry(
		ctx,
		req.Domain,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDNSRegistryResponse{DNSRegistry: val}, nil
}
