package group

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocTxMsgCreateGroupWithPolicy struct {
		Admin               string          `bson:"admin"`
		Members             []MemberRequest `bson:"members"`
		GroupMetadata       string          `bson:"group_metadata"`
		GroupPolicyMetadata string          `bson:"group_policy_metadata"`
		GroupPolicyAsAdmin  bool            `bson:"group_policy_as_admin"`
		DecisionPolicy      interface{}     `bson:"decision_policy"`
	}
)

func (m *DocTxMsgCreateGroupWithPolicy) GetType() string {
	return MsgTypeCreateGroupWithPolicy
}

func (m *DocTxMsgCreateGroupWithPolicy) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgCreateGroupWithPolicy)
	m.Admin = msg.Admin
	members := make([]MemberRequest, 0, len(msg.Members))
	for _, member := range msg.Members {
		members = append(members, MemberRequest{
			member.Address, member.Weight, member.Metadata,
		})
	}
	m.GroupMetadata = msg.GroupMetadata
	m.GroupPolicyMetadata = msg.GroupPolicyMetadata
	m.GroupPolicyAsAdmin = msg.GroupPolicyAsAdmin
	m.DecisionPolicy = msg.DecisionPolicy
}

func (m *DocTxMsgCreateGroupWithPolicy) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgCreateGroupWithPolicy)
	for _, member := range msg.Members {
		addrs = append(addrs, member.Address)
	}
	addrs = append(addrs, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
