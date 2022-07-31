package planet

import (
	"github.com/MercesNetwork/merces/x/planet/keeper"
	"github.com/MercesNetwork/merces/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the twitterCoins
	//for _, elem := range genState.TwitterCoinsList {
	//	// k.SetTwitterCoins(ctx, elem)
	//}
	// Set all the dNSRegistry
	for _, elem := range genState.DNSRegistryList {
		k.SetDNSRegistry(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	//genesis.TwitterCoinsList = k.GetAllTwitterCoins(ctx)
	genesis.DNSRegistryList = k.GetAllDNSRegistry(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
