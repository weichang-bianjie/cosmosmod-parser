package params

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/kaifei-bianjie/cosmosmod-parser/codec"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
