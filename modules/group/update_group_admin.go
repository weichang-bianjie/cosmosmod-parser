package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupAdmin struct {
		Admin    string `bson:"admin"`
		GroupId  uint64 `bson:"group_id"`
		NewAdmin string `bson:"new_admin"`
	}
)

func (m *DocTxMsgUpdateGroupAdmin) GetType() string {
	return MsgTypeUpdateGroupAdmin
}

func (m *DocTxMsgUpdateGroupAdmin) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupAdmin)
	m.Admin = msg.Admin
	m.GroupId = msg.GroupId
	m.NewAdmin = msg.NewAdmin
}
func (m *DocTxMsgUpdateGroupAdmin) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupAdmin)

	addrs = append(addrs, msg.Admin, msg.NewAdmin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
