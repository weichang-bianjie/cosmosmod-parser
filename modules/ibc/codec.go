package ibc

import (
	ibctransfer "github.com/cosmos/ibc-go/v5/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v5/modules/core"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
	)
}
