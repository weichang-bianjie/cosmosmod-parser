package authz

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type AuthzClient struct {
}

func NewClient() AuthzClient {
	return AuthzClient{}
}

func (authz AuthzClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgExec:
		docMsg := DocMsgExec{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgRevoke:
		docMsg := DocMsgRevoke{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgGrant:
		docMsg := DocMsgGrant{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
