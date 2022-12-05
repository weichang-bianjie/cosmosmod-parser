package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupPolicyAdmin struct {
		Admin              string `bson:"admin"`
		GroupPolicyAddress string `bson:"group_policy_address"`
		NewAdmin           string `bson:"new_admin"`
	}
)

func (m *DocTxMsgUpdateGroupPolicyAdmin) GetType() string {
	return MsgTypeUpdateGroupPolicyAdmin
}

func (m *DocTxMsgUpdateGroupPolicyAdmin) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupPolicyAdmin)
	m.Admin = msg.Admin
	m.GroupPolicyAddress = msg.GroupPolicyAddress
	m.NewAdmin = msg.NewAdmin
}

func (m *DocTxMsgUpdateGroupPolicyAdmin) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupPolicyAdmin)

	addrs = append(addrs, msg.Admin, msg.GroupPolicyAddress, msg.NewAdmin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
