package simulation

import (
	"math/rand"

	"github.com/MercesNetwork/merces/x/planet/keeper"
	"github.com/MercesNetwork/merces/x/planet/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgTransferFromTwitterToTwitterByMerces(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTransferFromTwitterToTwitterByMerces{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TransferFromTwitterToTwitterByMerces simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TransferFromTwitterToTwitterByMerces simulation not implemented"), nil, nil
	}
}
