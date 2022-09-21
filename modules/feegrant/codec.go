package feegrant

import (
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/kaifei-bianjie/cosmosmod-parser/codec"
)

func init() {
	codec.RegisterAppModules(feegrant.AppModuleBasic{})
}
