package upgrade

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type UpgradeClient struct {
}

func NewClient() UpgradeClient {
	return UpgradeClient{}
}

func (upgrade UpgradeClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := false

	return msgDocInfo, ok
}
