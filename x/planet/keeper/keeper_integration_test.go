package keeper_test

import (
	"testing"

	planetapp "github.com/MercesToken/planet/app"
	"github.com/MercesToken/planet/x/planet/keeper"
	"github.com/MercesToken/planet/x/planet/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"

	bobUsername   = "bob"
	aliceUsername = "alice"
	carolUsername = "carol"

	balAlice = 50000000
	balBob   = 20000000
	balCarol = 10000000
)

var (
	planetModuleAddress string
)

type IntegrationTestSuite struct {
	suite.Suite

	app         *planetapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestPlanetKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := planetapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	planetModuleAddress = app.AccountKeeper.GetModuleAddress(types.ModuleName).String()

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.PlanetKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.msgServer = keeper.NewMsgServerImpl(app.PlanetKeeper)
	suite.ctx = ctx
	suite.queryClient = queryClient
}

func makeBalance(address string, balance int64) banktypes.Balance {
	return banktypes.Balance{
		Address: address,
		Coins: sdk.Coins{
			sdk.Coin{
				Denom:  sdk.DefaultBondDenom,
				Amount: sdk.NewInt(balance),
			},
		},
	}
}

func getBankGenesis() *banktypes.GenesisState {
	coins := []banktypes.Balance{
		makeBalance(alice, balAlice),
		makeBalance(bob, balBob),
		makeBalance(carol, balCarol),
	}

	supply := sdk.NewCoins(coins[0].Coins.Add(coins[1].Coins...).Add(coins[2].Coins...)...)
	// supply := banktypes.NewSupply(coins[0].Coins.Add(coins[1].Coins...).Add(coins[2].Coins...))

	state := banktypes.NewGenesisState(
		banktypes.DefaultParams(),
		coins,
		supply,
		[]banktypes.Metadata{})

	return state
}

func (suite *IntegrationTestSuite) setupSuiteWithBalances() {
	suite.app.BankKeeper.InitGenesis(suite.ctx, getBankGenesis())
}

func (suite *IntegrationTestSuite) RequireBankBalance(expected int, atAddress string) {
	sdkAdd, err := sdk.AccAddressFromBech32(atAddress)
	suite.Require().Nil(err, "Address %s failed to parse")
	suite.Require().Equal(
		int64(expected),
		suite.app.BankKeeper.GetBalance(suite.ctx, sdkAdd, sdk.DefaultBondDenom).Amount.Int64())
}
