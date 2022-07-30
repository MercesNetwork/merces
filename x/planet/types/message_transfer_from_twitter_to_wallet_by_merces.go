package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferFromTwitterToWalletByMerces = "transfer_from_twitter_to_wallet_by_merces"

var _ sdk.Msg = &MsgTransferFromTwitterToWalletByMerces{}

func NewMsgTransferFromTwitterToWalletByMerces(creator string, username string, address string, denom string, amount int64) *MsgTransferFromTwitterToWalletByMerces {
	return &MsgTransferFromTwitterToWalletByMerces{
		Creator:  creator,
		Username: username,
		Address:  address,
		Denom:    denom,
		Amount:   amount,
	}
}

func (msg *MsgTransferFromTwitterToWalletByMerces) Route() string {
	return RouterKey
}

func (msg *MsgTransferFromTwitterToWalletByMerces) Type() string {
	return TypeMsgTransferFromTwitterToWalletByMerces
}

func (msg *MsgTransferFromTwitterToWalletByMerces) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferFromTwitterToWalletByMerces) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferFromTwitterToWalletByMerces) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
