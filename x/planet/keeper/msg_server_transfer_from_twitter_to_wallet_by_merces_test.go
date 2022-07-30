package keeper_test

import (
	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestTwotterToWallet() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.TransferFromTwitterToWalletByMerces(ctx, &types.MsgTransferFromTwitterToWalletByMerces{
		Creator:  keeper.RequiredAddress,
		Address:  bob,
		Username: bobUsername,
		Denom:    keeper.DenomTrial,
		Amount:   1,
	})
	suite.Require().Nil(err)

	bobTwitterCoins, found := suite.app.PlanetKeeper.GetTwitterCoins(suite.ctx, bobUsername, keeper.DenomTrial)
	suite.Require().True(found)
	suite.Require().Equal(int64(4), bobTwitterCoins.Amount)

	bobAddress, err := sdk.AccAddressFromBech32(bob)
	suite.Require().Nil(err)
	bobCoin := suite.app.BankKeeper.GetBalance(suite.ctx, bobAddress, keeper.DenomTrial)
	suite.Require().Equal(int64(1), bobCoin.Amount.Int64())
}

//func (suite *IntegrationTestSuite) TestTwitterToTwitter() {
//	ctx := sdk.WrapSDKContext(suite.ctx)
//}
