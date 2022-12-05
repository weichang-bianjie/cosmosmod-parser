package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgCreateGroup struct {
		Admin    string          `bson:"admin"`
		Members  []MemberRequest `bson:"members"`
		Metadata string          `bson:"metadata"`
	}

	MemberRequest struct {
		Address  string `bson:"address"`
		Weight   string `bson:"weight"`
		Metadata string `bson:"metadata"`
	}
)

func (m *DocTxMsgCreateGroup) GetType() string {
	return MsgTypeCreateGroup
}

func (m *DocTxMsgCreateGroup) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgCreateGroup)
	m.Admin = msg.Admin
	members := make([]MemberRequest, 0, len(msg.Members))
	for _, member := range msg.Members {
		members = append(members, MemberRequest{
			member.Address, member.Weight, member.Metadata,
		})
	}
	m.Members = members
	m.Metadata = msg.Metadata
}
func (m *DocTxMsgCreateGroup) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgCreateGroup)
	for _, member := range msg.Members {
		addrs = append(addrs, member.Address)
	}

	addrs = append(addrs, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
