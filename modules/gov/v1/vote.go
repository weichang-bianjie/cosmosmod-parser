package v1

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// MsgVote
type DocTxMsgVoteV1 struct {
	ProposalID int64  `bson:"proposal_id"` // ID of the proposal
	Voter      string `bson:"voter"`       //  address of the voter
	Option     int32  `bson:"option"`      //  option from OptionSet chosen by the voter
	Metadata   string `bson:"metadata"`
}

func (m *DocTxMsgVoteV1) GetType() string {
	return MsgTypeVote
}

func (m *DocTxMsgVoteV1) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgVoteV1)
	m.Voter = msg.Voter
	m.Option = int32(msg.Option)
	m.ProposalID = int64(msg.ProposalId)
	m.Metadata = msg.Metadata
}

func (m *DocTxMsgVoteV1) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgVoteV1)
	addrs = append(addrs, msg.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
