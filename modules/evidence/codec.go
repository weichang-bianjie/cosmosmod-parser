package evidence

import (
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(evidence.AppModuleBasic{})
}
