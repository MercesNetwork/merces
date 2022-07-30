package keeper

import (
	"context"
	"fmt"

	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferFromWalletToTwitter(goCtx context.Context, msg *types.MsgTransferFromWalletToTwitter) (*types.MsgTransferFromWalletToTwitterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	ctx.Logger().Error(fmt.Sprintf("Try transfering wallet to twitter, %v", msg))
	ctx.Logger().Error(fmt.Sprintf("username %s, source %s, denom %s, amount %v", msg.Username, msg.Creator, msg.Coin.GetDenom(), msg.Coin.Amount.Int64()))
	// TODO: Handling the message
	_ = ctx
	err := k.Keeper.TransferFromWalletToTwitter(ctx, msg.Creator, msg.Username, msg.Coin)
	if err != nil {
		ctx.Logger().Error("transfer err %v", err)
		return nil, err
	}

	return &types.MsgTransferFromWalletToTwitterResponse{}, nil
}
