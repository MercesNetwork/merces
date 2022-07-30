package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferFromWalletToTwitter = "transfer_from_wallet_to_twitter"

var _ sdk.Msg = &MsgTransferFromWalletToTwitter{}

func NewMsgTransferFromWalletToTwitter(creator string, username string, coin sdk.Coin) *MsgTransferFromWalletToTwitter {
	return &MsgTransferFromWalletToTwitter{
		Creator:  creator,
		Username: username,
		Coin:     coin,
	}
}

func (msg *MsgTransferFromWalletToTwitter) Route() string {
	return RouterKey
}

func (msg *MsgTransferFromWalletToTwitter) Type() string {
	return TypeMsgTransferFromWalletToTwitter
}

func (msg *MsgTransferFromWalletToTwitter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferFromWalletToTwitter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferFromWalletToTwitter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
