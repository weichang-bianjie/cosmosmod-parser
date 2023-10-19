package authz

import (
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(authz.AppModuleBasic{})
}
