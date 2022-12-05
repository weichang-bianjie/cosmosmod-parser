package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupPolicyMetadata struct {
		Admin              string `bson:"admin"`
		GroupPolicyAddress string `bson:"group_policy_address"`
		Metadata           string `bson:"metadata"`
	}
)

func (m *DocTxMsgUpdateGroupPolicyMetadata) GetType() string {
	return MsgTypeUpdateGroupPolicyMetadata
}

func (m *DocTxMsgUpdateGroupPolicyMetadata) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupPolicyMetadata)
	m.Admin = msg.Admin
	m.GroupPolicyAddress = msg.GroupPolicyAddress
	m.Metadata = msg.Metadata
}

func (m *DocTxMsgUpdateGroupPolicyMetadata) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupPolicyMetadata)
	addrs = append(addrs, msg.Admin, msg.GroupPolicyAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
