package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateStoredMiner{}, "cosmosminer/CreateStoredMiner", nil)
	cdc.RegisterConcrete(&MsgUpdateStoredMiner{}, "cosmosminer/UpdateStoredMiner", nil)
	cdc.RegisterConcrete(&MsgDeleteStoredMiner{}, "cosmosminer/DeleteStoredMiner", nil)
	cdc.RegisterConcrete(&MsgCreateMinerTemplate{}, "cosmosminer/CreateMinerTemplate", nil)
	cdc.RegisterConcrete(&MsgCreateMiner{}, "cosmosminer/CreateMiner", nil)
	cdc.RegisterConcrete(&MsgMine{}, "cosmosminer/Mine", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStoredMiner{},
		&MsgUpdateStoredMiner{},
		&MsgDeleteStoredMiner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMinerTemplate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMiner{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMine{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
