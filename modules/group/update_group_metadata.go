package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupMetadata struct {
		Admin    string `bson:"admin"`
		GroupId  uint64 `bson:"group_id"`
		Metadata string `bson:"metadata"`
	}
)

func (m *DocTxMsgUpdateGroupMetadata) GetType() string {
	return MsgTypeUpdateGroupMetadata
}

func (m *DocTxMsgUpdateGroupMetadata) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupMetadata)
	m.Admin = msg.Admin
	m.GroupId = msg.GroupId
	m.Metadata = msg.Metadata
}

func (m *DocTxMsgUpdateGroupMetadata) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupMetadata)

	addrs = append(addrs, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
