package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateSynchronizationFeedback{}, "trustmesh/CreateSynchronizationFeedback", nil)
	cdc.RegisterConcrete(&MsgUpdateSynchronizationFeedback{}, "trustmesh/UpdateSynchronizationFeedback", nil)
	cdc.RegisterConcrete(&MsgDeleteSynchronizationFeedback{}, "trustmesh/DeleteSynchronizationFeedback", nil)

	cdc.RegisterConcrete(&MsgCreateSynchronizationRequest{}, "trustmesh/CreateSynchronizationRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateSynchronizationRequest{}, "trustmesh/UpdateSynchronizationRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteSynchronizationRequest{}, "trustmesh/DeleteSynchronizationRequest", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSynchronizationFeedback{},
		&MsgUpdateSynchronizationFeedback{},
		&MsgDeleteSynchronizationFeedback{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSynchronizationRequest{},
		&MsgUpdateSynchronizationRequest{},
		&MsgDeleteSynchronizationRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
