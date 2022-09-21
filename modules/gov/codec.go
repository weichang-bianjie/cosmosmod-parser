package gov

import (
	spartangov "github.com/bianjieai/spartan-cosmos/module/gov/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/kaifei-bianjie/cosmosmod-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		gov.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		spartangov.AppModuleBasic{},
	)
}
