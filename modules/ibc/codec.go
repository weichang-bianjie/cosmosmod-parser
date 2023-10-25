package ibc

import (
	nfttransfer "github.com/bianjieai/nft-transfer"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclients "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		ibcclients.AppModuleBasic{},
		nfttransfer.AppModuleBasic{},
	)
}
