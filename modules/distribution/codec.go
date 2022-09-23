package distribution

import (
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(distribution.AppModuleBasic{})
}
