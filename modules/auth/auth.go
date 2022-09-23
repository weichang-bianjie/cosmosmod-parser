package auth

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
)

type AuthClient struct {
}

func NewClient() AuthClient {
	return AuthClient{}
}

func (auth AuthClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	return msgDocInfo, false
}
