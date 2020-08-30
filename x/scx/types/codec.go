package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAppendOrganization{}, "scx/MsgAppendOrganization", nil)
	cdc.RegisterConcrete(MsgChangeOrganizationApproval{}, "scx/MsgChangeOrganizationApproval", nil)
	cdc.RegisterConcrete(MsgCreateProduct{}, "scx/MsgCreateProduct", nil)
	cdc.RegisterConcrete(MsgCreateUnit{}, "scx/MsgCreateUnit", nil)
	cdc.RegisterConcrete(MsgTransferUnit{}, "scx/MsgTransferUnit", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
