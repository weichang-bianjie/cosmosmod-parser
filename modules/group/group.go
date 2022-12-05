package group

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type GroupClient struct {
}

func NewClient() GroupClient {
	return GroupClient{}
}

func (params GroupClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)

	switch msg := v.(type) {
	case *MsgCreateGroup:
		docMsg := DocTxMsgCreateGroup{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateGroupMembers:
		docMsg := DocTxMsgUpdateGroupMembers{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateGroupAdmin:
		docMsg := DocTxMsgUpdateGroupAdmin{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateGroupMetadata:
		docMsg := DocTxMsgUpdateGroupMetadata{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgCreateGroupPolicy:
		docMsg := DocTxMsgCreateGroupPolicy{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgCreateGroupWithPolicy:
		docMsg := DocTxMsgCreateGroupWithPolicy{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateGroupPolicyAdmin:
		docMsg := DocTxMsgUpdateGroupPolicyAdmin{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateGroupPolicyDecisionPolicy:
		docMsg := DocTxMsgUpdateGroupPolicyDecisionPolicy{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgUpdateGroupPolicyMetadata:
		docMsg := DocTxMsgUpdateGroupPolicyMetadata{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgGroupSubmitProposal:
		docMsg := DocTxMsgGroupSubmitProposal{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgWithdrawProposal:
		docMsg := DocTxMsgWithdrawProposal{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgGroupVote:
		docMsg := DocTxMsgGroupVote{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgGroupExec:
		docMsg := DocTxMsgGroupExec{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
	case *MsgLeaveGroup:
		docMsg := DocTxMsgLeaveGroup{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		return msgDocInfo, false
	}
	return MsgDocInfo{}, false
}
