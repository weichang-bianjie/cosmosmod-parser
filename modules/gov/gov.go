package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type GovClient struct {
}

func NewClient() GovClient {
	return GovClient{}
}

func (gov GovClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgSubmitProposal:
		docMsg := DocTxMsgSubmitProposal{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgVote:
		docMsg := DocTxMsgVote{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgDeposit:
		docMsg := DocTxMsgDeposit{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgVoteWeighted:
		docMsg := DocTxMsgVoteWeighted{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
