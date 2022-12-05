package v1beta1

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// MsgVote
type DocTxMsgVote struct {
	ProposalID int64  `bson:"proposal_id"` // ID of the proposal
	Voter      string `bson:"voter"`       //  address of the voter
	Option     int32  `bson:"option"`      //  option from OptionSet chosen by the voter
}

func (m *DocTxMsgVote) GetType() string {
	return MsgTypeVote
}

func (m *DocTxMsgVote) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgVote)
	m.Voter = msg.Voter
	m.Option = int32(msg.Option)
	m.ProposalID = int64(msg.ProposalId)
}

func (m *DocTxMsgVote) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgVote)
	addrs = append(addrs, msg.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
