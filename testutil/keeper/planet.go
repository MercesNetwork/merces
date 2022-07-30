package keeper

import (
	"testing"

	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

//func mountStores(stateStore storetypes.CommitMultiStore, storeKey storetypes.StoreKey, db tmdb.DB) {
//
//	dnsRegStoreKey := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s", storeKey.Name(), types.StoreDNSRegistrySuffixKey))
//	dnsRegReverseStoreKey := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s", storeKey.Name(), types.StoreDNSRegistryReverseSuffixKey))
//	dnsUsersKey := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s", storeKey.Name(), types.StoreDNSUsersSuffixKey))
//	dnsCoinsKey := sdk.NewKVStoreKey(fmt.Sprintf("%s-%s", storeKey.Name(), types.StoreDNSCoinsSuffixKey))
//	fmt.Println("store", dnsRegReverseStoreKey.Name())
//
//	stateStore.MountStoreWithDB(dnsRegStoreKey, sdk.StoreTypeIAVL, db)
//	stateStore.MountStoreWithDB(dnsRegReverseStoreKey, sdk.StoreTypeIAVL, db)
//	stateStore.MountStoreWithDB(dnsUsersKey, sdk.StoreTypeIAVL, db)
//	stateStore.MountStoreWithDB(dnsCoinsKey, sdk.StoreTypeIAVL, db)
//}

func PlanetKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	// mountStores(stateStore, storeKey, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"PlanetParams",
	)

	k := keeper.NewKeeper(
		*new(bankkeeper.Keeper),
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
