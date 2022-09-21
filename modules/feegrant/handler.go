package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type FeegrantClient struct {
}

func NewClient() FeegrantClient {
	return FeegrantClient{}
}

func (feegrant FeegrantClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgGrantAllowance:
		docMsg := DocTxMsgGrantAllowance{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgRevokeAllowance:
		docMsg := DocTxMsgRevokeAllowance{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
