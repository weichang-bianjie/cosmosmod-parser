package gov

import (
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		gov.AppModuleBasic{},
	)
}
