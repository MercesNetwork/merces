package simulation

import (
	"math/rand"

	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgTransferFromWalletToTwitter(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTransferFromWalletToTwitter{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TransferFromWalletToTwitter simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "TransferFromWalletToTwitter simulation not implemented"), nil, nil
	}
}