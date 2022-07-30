package keeper_test

import (
	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//const (
//	rUsername = "Na_dsfF"
//)

func (suite *IntegrationTestSuite) TestWalletToTwitter() {
	suite.setupSuiteWithBalances()
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.TransferFromWalletToTwitter(ctx, &types.MsgTransferFromWalletToTwitter{
		Creator:  alice,
		Username: rUsername,
		Coin:     sdk.Coin{Denom: sdk.DefaultBondDenom, Amount: sdk.NewInt(4)},
	})
	suite.Require().Nil(err)

	coins, found := suite.app.PlanetKeeper.GetTwitterCoins(suite.ctx, rUsername, sdk.DefaultBondDenom)
	suite.Require().True(found)
	suite.Require().Equal(int64(4), coins.Amount)
}
