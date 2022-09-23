package crisis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type CrisisClient struct {
}

func NewClient() CrisisClient {
	return CrisisClient{}
}

func (crisis CrisisClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgVerifyInvariant:
		docMsg := DocMsgVerifyInvariant{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
