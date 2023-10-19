package authz

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocMsgRevoke struct {
		Granter    string `bson:"granter"`
		Grantee    string `bson:"grantee"`
		MsgTypeUrl string `bson:"msg_type_url"`
	}
)

func (m *DocMsgRevoke) GetType() string {
	return MsgTypeAuthzRevoke
}

func (m *DocMsgRevoke) BuildMsg(v interface{}) {
	msg := v.(*MsgRevoke)
	m.Granter = msg.Granter
	m.Grantee = msg.Grantee
	m.MsgTypeUrl = msg.MsgTypeUrl
}

func (m *DocMsgRevoke) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgRevoke)
	addrs = append(addrs, msg.Granter, msg.Grantee)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
