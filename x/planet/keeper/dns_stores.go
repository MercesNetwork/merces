package keeper

import (
	"fmt"

	"github.com/MercesNetwork/merces/x/planet/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetStoreDNSRegistryReverse(ctx sdk.Context) sdk.KVStore {
	key := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s", k.storeKey.Name(), types.StoreDNSRegistryReverseSuffixKey))
	return prefix.NewStore(ctx.KVStore(key), types.KeyPrefix(types.DNSRegistryReverseKeyPrefix))
}

func (k Keeper) GetStoreDNSUsers(ctx sdk.Context, domainName string) sdk.KVStore {
	key := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s-%s", k.storeKey.Name(), types.StoreDNSUsersSuffixKey, domainName))
	return ctx.KVStore(key)
}

func (k Keeper) GetStoreDNSCoins(ctx sdk.Context, domainName string) sdk.KVStore {
	return ctx.KVStore(k.GetStoreDNSCoinsKey(domainName))
}

func (k Keeper) GetStoreDNSCoinsKey(domainName string) *sdk.KVStoreKey {
	return sdk.NewKVStoreKey(fmt.Sprintf("%s-%s-%s", k.storeKey.Name(), types.StoreDNSCoinsSuffixKey, domainName))
}

func (k Keeper) GetStoreDNSUserCoins(ctx sdk.Context, domainName, username string) sdk.KVStore {
	store := k.GetStoreDNSCoins(ctx, domainName)
	return prefix.NewStore(store, types.KeyPrefix(username))
}

func (k Keeper) GetDNSScopedStore(ctx sdk.Context) prefix.Store {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, k.createDNSScopedPrefix())
}

func (k Keeper) createDNSScopedPrefix() []byte {
	return append(types.KeyPrefix(types.ModuleName), types.KeyPrefix(types.StoreDNSKey)...)
}
