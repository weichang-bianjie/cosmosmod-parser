package params

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type ParamsClient struct {
}

func NewClient() ParamsClient {
	return ParamsClient{}
}

func (params ParamsClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := false

	return msgDocInfo, ok
}
