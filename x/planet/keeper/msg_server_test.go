package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/MercesNetwork/merces/testutil/keeper"
	"github.com/MercesNetwork/merces/x/planet/keeper"
	"github.com/MercesNetwork/merces/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PlanetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
