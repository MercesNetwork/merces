package keeper

import (
	"fmt"

	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DenomTrial         = "trialtoken"
	DefaultAmountTrial = int64(5)
)

func (k Keeper) TransferBetweenTwitterUsers(ctx sdk.Context, from, to string, denom string, amount int64) error {
	fromCoin, found := k.GetTwitterCoins(ctx, from, denom)
	if !found {
		if denom != DenomTrial {
			return sdkerrors.ErrInsufficientFunds
		}
		trialCoin, err := k.initTrialTokenForUsername(ctx, from)
		if err != nil {
			return err
		}
		fromCoin = types.TwitterCoins{Index: denom, Amount: trialCoin.Amount.Int64()}
	}
	fromCoin.Amount = sdk.NewCoin(fromCoin.Index, sdk.NewInt(fromCoin.Amount)).
		SubAmount(sdk.NewInt(amount)).
		Amount.Int64()

	toCoin, found := k.GetTwitterCoins(ctx, to, denom)
	if !found {
		toCoin = types.TwitterCoins{
			Index:  denom,
			Amount: 0,
		}
	}
	toCoin.Amount = sdk.NewCoin(toCoin.Index, sdk.NewInt(toCoin.Amount)).
		AddAmount(sdk.NewInt(amount)).
		Amount.Int64()

	k.SetTwitterCoins(ctx, from, fromCoin)
	k.SetTwitterCoins(ctx, to, toCoin)
	return nil
}

func (k Keeper) TransferFromWalletToTwitter(ctx sdk.Context, address, username string, coin sdk.Coin) error {
	walletAccount, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	twitterCoins, found := k.GetTwitterCoins(ctx, username, coin.Denom)
	if !found {
		twitterCoins.Index = coin.Denom
		twitterCoins.Amount = 0
	}
	twitterCoins.Amount = sdk.NewCoin(twitterCoins.Index, sdk.NewInt(twitterCoins.Amount)).
		AddAmount(coin.Amount).Amount.Int64()
	k.SetTwitterCoins(ctx, username, twitterCoins)
	err = k.bank.SendCoinsFromAccountToModule(ctx, walletAccount, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) TransferBetweenFromTwitterToWallet(ctx sdk.Context, username, address string, denom string, amount int64) error {
	walletAccount, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("wallet addr %s from %s", walletAccount.String(), address))
	fromCoin, found := k.GetTwitterCoins(ctx, username, denom)
	if !found {
		if denom != DenomTrial {
			return sdkerrors.ErrInsufficientFunds
		}
		trialCoin, err := k.initTrialTokenForUsername(ctx, username)
		if err != nil {
			return err
		}
		fromCoin = types.TwitterCoins{Index: denom, Amount: trialCoin.Amount.Int64()}
	}
	fromCoin.Amount = sdk.NewCoin(fromCoin.Index, sdk.NewInt(fromCoin.Amount)).
		SubAmount(sdk.NewInt(amount)).
		Amount.Int64()
	k.SetTwitterCoins(ctx, username, fromCoin)

	coins := sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(amount)))

	if err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, walletAccount, coins); err != nil {
		return err
	}

	return nil
}

func (k Keeper) initTrialTokenForUsername(ctx sdk.Context, username string) (sdk.Coin, error) {
	trialCoin := sdk.NewCoin(DenomTrial, sdk.NewInt(DefaultAmountTrial))
	if err := k.bank.MintCoins(ctx, types.ModuleName, sdk.NewCoins(trialCoin)); err != nil {
		return sdk.Coin{}, err
	}
	return trialCoin, nil
}
