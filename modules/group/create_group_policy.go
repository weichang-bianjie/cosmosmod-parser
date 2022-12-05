package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgCreateGroupPolicy struct {
		Admin          string      `bson:"admin"`
		GroupId        uint64      `bson:"group_id"`
		Metadata       string      `bson:"metadata"`
		DecisionPolicy interface{} `bson:"decision_policy"`
	}
)

func (m *DocTxMsgCreateGroupPolicy) GetType() string {
	return MsgTypeCreateGroupPolicy
}

func (m *DocTxMsgCreateGroupPolicy) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgCreateGroupPolicy)
	m.Admin = msg.Admin
	m.GroupId = msg.GroupId
	m.Metadata = msg.Metadata
	m.DecisionPolicy = msg.DecisionPolicy
}

func (m *DocTxMsgCreateGroupPolicy) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgCreateGroupPolicy)

	addrs = append(addrs, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
