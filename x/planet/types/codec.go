package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgTransferFromTwitterToWalletByMerces{}, "planet/TransferFromTwitterToWalletByMerces", nil)
	cdc.RegisterConcrete(&MsgTransferFromTwitterToTwitterByMerces{}, "planet/TransferFromTwitterToTwitterByMerces", nil)
	cdc.RegisterConcrete(&MsgTransferFromWalletToTwitter{}, "planet/TransferFromWalletToTwitter", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferFromTwitterToWalletByMerces{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferFromTwitterToTwitterByMerces{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferFromWalletToTwitter{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
