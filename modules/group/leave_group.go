package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgLeaveGroup struct {
		Address string `bson:"address"`
		GroupId uint64 `bson:"group_id"`
	}
)

func (m *DocTxMsgLeaveGroup) GetType() string {
	return MsgTypeLeaveGroup
}

func (m *DocTxMsgLeaveGroup) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgLeaveGroup)
	m.Address = msg.Address
	m.GroupId = msg.GroupId
}

func (m *DocTxMsgLeaveGroup) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgLeaveGroup)
	addrs = append(addrs, msg.Address)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
