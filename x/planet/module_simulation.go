package planet

import (
	"math/rand"

	"github.com/MercesNetwork/merces/testutil/sample"
	planetsimulation "github.com/MercesNetwork/merces/x/planet/simulation"
	"github.com/MercesNetwork/merces/x/planet/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = planetsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgTransferFromTwitterToWalletByMerces = "op_weight_msg_transfer_from_twitter_to_wallet_by_merces"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferFromTwitterToWalletByMerces int = 100

	opWeightMsgTransferFromTwitterToTwitterByMerces = "op_weight_msg_transfer_from_twitter_to_twitter_by_merces"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferFromTwitterToTwitterByMerces int = 100

	opWeightMsgTransferFromWalletToTwitter = "op_weight_msg_transfer_from_wallet_to_twitter"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTransferFromWalletToTwitter int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	planetGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&planetGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgTransferFromTwitterToWalletByMerces int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferFromTwitterToWalletByMerces, &weightMsgTransferFromTwitterToWalletByMerces, nil,
		func(_ *rand.Rand) {
			weightMsgTransferFromTwitterToWalletByMerces = defaultWeightMsgTransferFromTwitterToWalletByMerces
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferFromTwitterToWalletByMerces,
		planetsimulation.SimulateMsgTransferFromTwitterToWalletByMerces(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferFromTwitterToTwitterByMerces int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferFromTwitterToTwitterByMerces, &weightMsgTransferFromTwitterToTwitterByMerces, nil,
		func(_ *rand.Rand) {
			weightMsgTransferFromTwitterToTwitterByMerces = defaultWeightMsgTransferFromTwitterToTwitterByMerces
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferFromTwitterToTwitterByMerces,
		planetsimulation.SimulateMsgTransferFromTwitterToTwitterByMerces(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTransferFromWalletToTwitter int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTransferFromWalletToTwitter, &weightMsgTransferFromWalletToTwitter, nil,
		func(_ *rand.Rand) {
			weightMsgTransferFromWalletToTwitter = defaultWeightMsgTransferFromWalletToTwitter
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransferFromWalletToTwitter,
		planetsimulation.SimulateMsgTransferFromWalletToTwitter(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
