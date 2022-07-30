package keeper

import (
	"context"
	"fmt"

	"github.com/MercesToken/planet/x/planet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	RequiredAddress = "cosmos1fwa553eydyanguhzaycc653j426gzynjsrauxn"
)

func (k msgServer) TransferFromTwitterToWalletByMerces(goCtx context.Context, msg *types.MsgTransferFromTwitterToWalletByMerces) (*types.MsgTransferFromTwitterToWalletByMercesResponse, error) {
	fmt.Println("enter")
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	//signers := msg.GetSigners()
	//if len(signers) != 1 || signers[0].String() != RequiredAddress {
	//	return nil, sdkerrors.ErrorInvalidSigner
	//}
	err := k.Keeper.TransferBetweenFromTwitterToWallet(ctx, msg.Username, msg.Address, msg.Denom, msg.Amount)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &types.MsgTransferFromTwitterToWalletByMercesResponse{}, nil
}
