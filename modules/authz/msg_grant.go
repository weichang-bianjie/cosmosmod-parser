package authz

import (
	"encoding/json"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	DocMsgGrant struct {
		Granter string `bson:"granter"`
		Grantee string `bson:"grantee"`
		Grant   Grant  `bson:"grant"`
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

	marshalJSON, err := commoncodec.GetMarshaler().MarshalJSON(msg.Grant.Authorization)
	if err != nil {
		logrus.Errorf("DocMsgGrant commoncodec.GetMarshaler().MarshalJSON err:%v", err)
	}

	var authorization interface{}
	if err = json.Unmarshal(marshalJSON, &authorization); err != nil {
		logrus.Errorf("DocMsgGrant json.Unmarshal err:%v", err)
	}
	m.Grant.Authorization = authorization
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
