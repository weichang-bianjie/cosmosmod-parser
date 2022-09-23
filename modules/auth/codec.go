package auth

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(auth.AppModuleBasic{})
}
