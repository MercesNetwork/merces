package keeper_test

import (
	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (suite *IntegrationTestSuite) TestTwitterToTwitter() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.TransferFromTwitterToTwitterByMerces(ctx, &types.MsgTransferFromTwitterToTwitterByMerces{
		FromUsername: bobUsername,
		ToUsername:   aliceUsername,
		Creator:      keeper.RequiredAddress,
		Denom:        keeper.DenomTrial,
		Amount:       1,
	})
	suite.Require().Nil(err)

	bobTwitterCoins, found := suite.app.PlanetKeeper.GetTwitterCoins(suite.ctx, bobUsername, keeper.DenomTrial)
	suite.Require().True(found)
	suite.Require().Equal(int64(4), bobTwitterCoins.Amount)

	aliceTwitterCoins, found := suite.app.PlanetKeeper.GetTwitterCoins(suite.ctx, aliceUsername, keeper.DenomTrial)
	suite.Require().True(found)
	suite.Require().Equal(int64(1), aliceTwitterCoins.Amount)
}

func (suite *IntegrationTestSuite) TestTwitterEmptyingToken() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.TransferFromTwitterToTwitterByMerces(ctx, &types.MsgTransferFromTwitterToTwitterByMerces{
		FromUsername: bobUsername,
		ToUsername:   aliceUsername,
		Creator:      keeper.RequiredAddress,
		Denom:        keeper.DenomTrial,
		Amount:       1,
	})
	suite.Require().Nil(err)

	_, err = suite.msgServer.TransferFromTwitterToTwitterByMerces(ctx, &types.MsgTransferFromTwitterToTwitterByMerces{
		FromUsername: bobUsername,
		ToUsername:   carolUsername,
		Creator:      keeper.RequiredAddress,
		Denom:        keeper.DenomTrial,
		Amount:       4,
	})
	suite.Require().Nil(err)

	bobTwitterCoins, _ := suite.app.PlanetKeeper.GetTwitterCoins(suite.ctx, bobUsername, keeper.DenomTrial)
	aliceTwitterCoins, _ := suite.app.PlanetKeeper.GetTwitterCoins(suite.ctx, aliceUsername, keeper.DenomTrial)
	carolTwitterCoins, _ := suite.app.PlanetKeeper.GetTwitterCoins(suite.ctx, carolUsername, keeper.DenomTrial)
	suite.Require().Equal(int64(0), bobTwitterCoins.Amount)
	suite.Require().Equal(int64(1), aliceTwitterCoins.Amount)
	suite.Require().Equal(int64(4), carolTwitterCoins.Amount)

	_, err = suite.msgServer.TransferFromTwitterToTwitterByMerces(ctx, &types.MsgTransferFromTwitterToTwitterByMerces{
		FromUsername: bobUsername,
		ToUsername:   carolUsername,
		Creator:      keeper.RequiredAddress,
		Denom:        keeper.DenomTrial,
		Amount:       4,
	})

	_, code, _ := sdkerrors.ABCIInfo(err, false)
	suite.Require().Equal(uint32(5), code)
}

func (suite *IntegrationTestSuite) TestTwitterToTwitterNoBalance() {
	ctx := sdk.WrapSDKContext(suite.ctx)
	_, err := suite.msgServer.TransferFromTwitterToTwitterByMerces(ctx, &types.MsgTransferFromTwitterToTwitterByMerces{
		Creator:      keeper.RequiredAddress,
		FromUsername: bobUsername,
		ToUsername:   aliceUsername,
		Denom:        "notoken",
		Amount:       1,
	})
	_, code, _ := sdkerrors.ABCIInfo(err, false)
	suite.Require().Equal(uint32(5), code)
}
