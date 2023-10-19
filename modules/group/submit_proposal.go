package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgGroupSubmitProposal struct {
		GroupPolicyAddress string      `bson:"group_policy_address"`
		Proposers          []string    `bson:"proposers"`
		Metadata           string      `bson:"metadata"`
		Messages           interface{} `bson:"messages"`
		Exec               int32       `bson:"exec"`
	}
)

func (m *DocTxMsgGroupSubmitProposal) GetType() string {
	return MsgTypeGroupSubmitProposal
}

func (m *DocTxMsgGroupSubmitProposal) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgGroupSubmitProposal)
	m.GroupPolicyAddress = msg.GroupPolicyAddress
	m.Proposers = msg.Proposers
	m.Metadata = msg.Metadata
	m.Messages = msg.Messages
	m.Exec = int32(msg.Exec)
}

func (m *DocTxMsgGroupSubmitProposal) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgGroupSubmitProposal)
	addrs = append(addrs, msg.GroupPolicyAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
