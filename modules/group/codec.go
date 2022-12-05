package group

import (
	group "github.com/cosmos/cosmos-sdk/x/group/module"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(group.AppModuleBasic{})
}
