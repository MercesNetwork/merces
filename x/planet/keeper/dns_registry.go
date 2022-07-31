package keeper

import (
	"github.com/MercesNetwork/merces/x/planet/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDNSRegistry set a specific dNSRegistry in the store from its index
func (k Keeper) SetDNSRegistry(ctx sdk.Context, dNSRegistry types.DNSRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DNSRegistryKeyPrefix))
	//store := k.GetDNSScopedStore(ctx)
	b := k.cdc.MustMarshal(&dNSRegistry)
	store.Set(types.DNSRegistryKey(
		dNSRegistry.Domain,
	), b)
}

// GetDNSRegistry returns a dNSRegistry from its index
func (k Keeper) GetDNSRegistry(
	ctx sdk.Context,
	domain string,

) (val types.DNSRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DNSRegistryKeyPrefix))
	//store := k.GetDNSScopedStore(ctx)

	b := store.Get(types.DNSRegistryKey(
		domain,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDNSRegistry removes a dNSRegistry from the store
func (k Keeper) RemoveDNSRegistry(
	ctx sdk.Context,
	domain string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DNSRegistryKeyPrefix))
	// store := k.GetDNSScopedStore(ctx)
	store.Delete(types.DNSRegistryKey(
		domain,
	))
}

// GetAllDNSRegistry returns all dNSRegistry
func (k Keeper) GetAllDNSRegistry(ctx sdk.Context) (list []types.DNSRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DNSRegistryKeyPrefix))
	// store := k.GetDNSScopedStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DNSRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
