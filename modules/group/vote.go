package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgGroupVote struct {
		ProposalId uint64 `bson:"proposal_id"`
		Voter      string `bson:"voter"`
		Option     int32  `bson:"option"`
		Metadata   string `bson:"metadata"`
		Exec       int32  `bson:"exec"`
	}
)

func (m *DocTxMsgGroupVote) GetType() string {
	return MsgTypeGroupVote
}

func (m *DocTxMsgGroupVote) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgGroupVote)
	m.ProposalId = msg.ProposalId
	m.Voter = msg.Voter
	m.Option = int32(msg.Option)
	m.Metadata = msg.Metadata
	m.Exec = int32(msg.Exec)
}

func (m *DocTxMsgGroupVote) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgGroupVote)
	addrs = append(addrs, msg.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
