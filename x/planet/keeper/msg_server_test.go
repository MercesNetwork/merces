package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/MercesToken/planet/testutil/keeper"
	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PlanetKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
