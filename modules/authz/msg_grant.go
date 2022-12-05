package authz

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
	"time"
)

type (
	DocMsgGrant struct {
		Granter string `bson:"granter"`
		Grantee string `bson:"grantee"`
		Grant   Grant  `bson:"msgs"`
	}

	Grant struct {
		Authorization interface{} `bson:"authorization"`
		Expiration    *time.Time  `bson:"expiration"`
	}
)

func (m *DocMsgGrant) GetType() string {
	return MsgTypeAuthzGrant
}

func (m *DocMsgGrant) BuildMsg(v interface{}) {
	msg := v.(*MsgGrant)
	m.Granter = msg.Granter
	m.Grantee = msg.Grantee
	m.Grant.Authorization = msg.Grant.Authorization.GetCachedValue().(*GenericAuthorization)
	m.Grant.Expiration = msg.Grant.Expiration
}

func (m *DocMsgGrant) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgGrant)
	addrs = append(addrs, msg.Granter, msg.Grantee)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
