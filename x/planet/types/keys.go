package types

import fmt "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "planet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_planet"

	StoreDNSKey                      = "dns"
	StoreDNSRegistryReverseSuffixKey = "DNSRegistryReverse"
	StoreDNSUsersSuffixKey           = "uns"
	StoreDNSCoinsSuffixKey           = "DNSCoins"
)

var (
	DnsRegReverseStoreKey = fmt.Sprintf("%s-%s", ModuleName, StoreDNSRegistryReverseSuffixKey)
	DnsUsersKey           = fmt.Sprintf("%s-%s", ModuleName, StoreDNSUsersSuffixKey)
	DnsCoinsKey           = fmt.Sprintf("%s-%s", ModuleName, StoreDNSCoinsSuffixKey)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

//func GetStoreDNSRegistryKey() string {
//	fmt.Sprintf("%s_")
//}

//func (k Keeper) GetStoreDNSRegistry(ctx sdk.Context) sdk.KVStore {
//	key := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s", k.storeKey.Name(), types.StoreDNSRegistrySuffixKey))
//	return prefix.NewStore(ctx.KVStore(key), types.KeyPrefix(types.DNSRegistryKeyPrefix))
//}
//
//func (k Keeper) GetStoreDNSRegistryReverse(ctx sdk.Context) sdk.KVStore {
//	key := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s", k.storeKey.Name(), types.StoreDNSRegistrySuffixKey))
//	return prefix.NewStore(ctx.KVStore(key), types.KeyPrefix(types.DNSRegistryReverseKeyPrefix))
//}
//
//func (k Keeper) GetStoreDNSUsers(ctx sdk.Context, domainName string) sdk.KVStore {
//	key := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s-%s", k.storeKey.Name(), types.StoreDNSUsersSuffixKey, domainName))
//	return ctx.KVStore(key)
//}
//
//func (k Keeper) GetStoreDNSCoins(ctx sdk.Context, domainName string) sdk.KVStore {
//	return ctx.KVStore(k.GetStoreDNSCoinsKey(domainName))
//}
//
//func (k Keeper) GetStoreDNSCoinsKey(domainName string) *sdk.KVStoreKey {
//	return sdk.NewKVStoreKey(fmt.Sprintf("%s-%s-%s", k.storeKey.Name(), types.StoreDNSCoinsSuffixKey, domainName))
//}
//
//func (k Keeper) GetStoreDNSUserCoins(ctx sdk.Context, domainName, username string) sdk.KVStore {
//	store := k.GetStoreDNSCoins(ctx, domainName)
//	return prefix.NewStore(store, types.KeyPrefix(username))
//}
