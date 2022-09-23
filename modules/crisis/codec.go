package crisis

import (
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(crisis.AppModuleBasic{})
}
