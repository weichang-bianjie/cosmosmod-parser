package codec

import (
	"github.com/bianjieai/spartan-cosmos/module/node"
	"github.com/cosmos/cosmos-sdk/codec"
	amino "github.com/cosmos/cosmos-sdk/codec"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/tendermint/tendermint/types"
	enccodec "github.com/tharsis/ethermint/encoding/codec"
)

var (
	appModules []module.AppModuleBasic
	encodecfg  params.EncodingConfig
)

// 初始化账户地址前缀
func MakeEncodingConfig() {
	moduleBasics := module.NewBasicManager(appModules...)
	encodingConfig := MakeConfig(moduleBasics)

	encodecfg = encodingConfig
}

// MakeConfig creates an EncodingConfig for testing
func MakeConfig(mb module.BasicManager) params.EncodingConfig {
	cdc := amino.NewLegacyAmino()
	interfaceRegistry := types2.NewInterfaceRegistry()
	marshaler := amino.NewProtoCodec(interfaceRegistry)

	node.RegisterInterfaces(interfaceRegistry)
	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             cdc,
	}

	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	mb.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	mb.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

func GetTxDecoder() sdk.TxDecoder {
	return encodecfg.TxConfig.TxDecoder()
}

func GetMarshaler() codec.Codec {
	return encodecfg.Marshaler
}

func GetSigningTx(txBytes types.Tx) (signing.Tx, error) {
	Tx, err := GetTxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}
	signingTx := Tx.(signing.Tx)
	return signingTx, nil
}

func RegisterAppModules(module ...module.AppModuleBasic) {
	appModules = append(appModules, module...)
}
