package keeper

import (
	"context"

	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) TransferFromTwitterToTwitterByMerces(goCtx context.Context, msg *types.MsgTransferFromTwitterToTwitterByMerces) (*types.MsgTransferFromTwitterToTwitterByMercesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	signers := msg.GetSigners()
	if len(signers) != 1 || signers[0].String() != RequiredAddress {
		return nil, sdkerrors.ErrorInvalidSigner
	}
	err := k.Keeper.TransferBetweenTwitterUsers(ctx, msg.FromUsername, msg.ToUsername, msg.Denom, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferFromTwitterToTwitterByMercesResponse{}, nil
}
