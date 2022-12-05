package authz

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocMsgExec struct {
		Grantee string        `bson:"grantee"`
		Msgs    []interface{} `bson:"msgs"`
	}
)

func (m *DocMsgExec) GetType() string {
	return MsgTypeAuthzExec
}

func (m *DocMsgExec) BuildMsg(v interface{}) {
	msg := v.(*MsgExec)
	m.Grantee = msg.Grantee
	msgs := make([]interface{}, 0, len(msg.Msgs))
	messages, _ := msg.GetMessages()
	for _, m := range messages {
		msgs = append(msgs, m)
	}
	m.Msgs = msgs
}

func (m *DocMsgExec) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgExec)
	addrs = append(addrs, msg.Grantee)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
