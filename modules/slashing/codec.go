package slashing

import (
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/kaifei-bianjie/cosmosmod-parser/codec"
)

func init() {
	codec.RegisterAppModules(slashing.AppModuleBasic{})
}
