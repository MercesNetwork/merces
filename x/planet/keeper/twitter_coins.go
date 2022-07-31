package keeper

import (
	"github.com/MercesNetwork/merces/x/planet/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) getTwitterAccountScopedStore(ctx sdk.Context, username string) prefix.Store {
	store := ctx.KVStore(k.storeKey)

	return prefix.NewStore(store, k.createTwitterAccountScopedPrefix(username))
}

func (k Keeper) createTwitterAccountScopedPrefix(username string) []byte {
	return append(types.KeyPrefix(types.TwitterCoinsKeyPrefix), types.KeyPrefix(username)...)
}

// SetTwitterCoins set a specific twitterCoins in the store from its index
func (k Keeper) SetTwitterCoins(ctx sdk.Context, username string, twitterCoins types.TwitterCoins) {
	//store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TwitterCoinsKeyPrefix))
	store := k.getTwitterAccountScopedStore(ctx, username)
	b := k.cdc.MustMarshal(&twitterCoins)
	store.Set(types.TwitterCoinsKey(
		twitterCoins.Index,
	), b)
}

// GetTwitterCoins returns a twitterCoins from its index
func (k Keeper) GetTwitterCoins(
	ctx sdk.Context,
	username string,
	index string,

) (val types.TwitterCoins, found bool) {
	//store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TwitterCoinsKeyPrefix))
	store := k.getTwitterAccountScopedStore(ctx, username)

	b := store.Get(types.TwitterCoinsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTwitterCoins removes a twitterCoins from the store
func (k Keeper) RemoveTwitterCoins(
	ctx sdk.Context,
	username string,
	index string,

) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TwitterCoinsKeyPrefix))
	store := k.getTwitterAccountScopedStore(ctx, username)
	store.Delete(types.TwitterCoinsKey(
		index,
	))
}

// GetAllTwitterCoins returns all twitterCoins
func (k Keeper) GetAllTwitterCoins(ctx sdk.Context, username string) (list []types.TwitterCoins) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TwitterCoinsKeyPrefix))
	store := k.getTwitterAccountScopedStore(ctx, username)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TwitterCoins
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
