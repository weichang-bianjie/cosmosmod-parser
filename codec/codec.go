package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
)

// MakeEncodingConfig 初始化账户地址前缀
func MakeEncodingConfig() {
	var cdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(cdc)

	interfaceRegistry := ctypes.NewInterfaceRegistry()
	moduleBasics := module.NewBasicManager(commoncodec.AppModules...)
	moduleBasics.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	commoncodec.Encodecfg = commoncodec.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             cdc,
	}
}
