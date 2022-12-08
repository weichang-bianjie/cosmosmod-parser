package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
	govv1 "github.com/kaifei-bianjie/cosmosmod-parser/modules/gov/v1"
	govv1beta1 "github.com/kaifei-bianjie/cosmosmod-parser/modules/gov/v1beta1"
)

type GovClient struct {
}

func NewClient() GovClient {
	return GovClient{}
}

func (gov GovClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgSubmitProposal:
		docMsg := govv1beta1.DocTxMsgSubmitProposal{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgVote:
		docMsg := govv1beta1.DocTxMsgVote{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgDeposit:
		docMsg := govv1beta1.DocTxMsgDeposit{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgVoteWeighted:
		docMsg := govv1beta1.DocTxMsgVoteWeighted{}
		return docMsg.HandleTxMsg(msg), true

	case *MsgExecLegacyContent:
		docMsg := govv1.DocTxMsgExecLegacyContent{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgSubmitProposalV1:
		docMsg := govv1.DocTxMsgSubmitProposalV1{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgVoteV1:
		docMsg := govv1.DocTxMsgVoteV1{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgDepositV1:
		docMsg := govv1.DocTxMsgDepositV1{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgVoteWeightedV1:
		docMsg := govv1.DocTxMsgVoteWeightedV1{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
