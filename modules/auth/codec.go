package auth

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/kaifei-bianjie/cosmosmod-parser/codec"
)

func init() {
	codec.RegisterAppModules(auth.AppModuleBasic{})
}
