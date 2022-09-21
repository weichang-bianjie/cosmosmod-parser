package staking

import (
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/kaifei-bianjie/cosmosmod-parser/codec"
)

func init() {
	codec.RegisterAppModules(staking.AppModuleBasic{})
}
