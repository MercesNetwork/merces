package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferFromTwitterToTwitterByMerces = "transfer_from_twitter_to_twitter_by_merces"

var _ sdk.Msg = &MsgTransferFromTwitterToTwitterByMerces{}

func NewMsgTransferFromTwitterToTwitterByMerces(creator string, fromUsername string, toUsername string, denom string, amount int64) *MsgTransferFromTwitterToTwitterByMerces {
	return &MsgTransferFromTwitterToTwitterByMerces{
		Creator:      creator,
		FromUsername: fromUsername,
		ToUsername:   toUsername,
		Denom:        denom,
		Amount:       amount,
	}
}

func (msg *MsgTransferFromTwitterToTwitterByMerces) Route() string {
	return RouterKey
}

func (msg *MsgTransferFromTwitterToTwitterByMerces) Type() string {
	return TypeMsgTransferFromTwitterToTwitterByMerces
}

func (msg *MsgTransferFromTwitterToTwitterByMerces) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferFromTwitterToTwitterByMerces) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferFromTwitterToTwitterByMerces) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
