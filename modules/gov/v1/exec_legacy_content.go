package v1

import (
	"encoding/json"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
	"github.com/sirupsen/logrus"
)

// MsgExecLegacyContent
type DocTxMsgExecLegacyContent struct {
	Content   interface{} `bson:"content"`
	Authority string      `bson:"authority"`
}

func (m *DocTxMsgExecLegacyContent) GetType() string {
	return MsgTypeExecLegacyContent
}

func (m *DocTxMsgExecLegacyContent) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgExecLegacyContent)

	marshalJSON, err := commoncodec.GetMarshaler().MarshalJSON(msg.Content)
	if err != nil {
		logrus.Errorf("DocTxMsgExecLegacyContent commoncodec.GetMarshaler().MarshalJSON err:%v", err)
	}

	var msgInterface interface{}
	if err = json.Unmarshal(marshalJSON, &msgInterface); err != nil {
		logrus.Errorf("DocTxMsgExecLegacyContent json.Unmarshal err:%v", err)
	}

	m.Content = msgInterface
	m.Authority = msg.Authority
}

func (m *DocTxMsgExecLegacyContent) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgExecLegacyContent)
	addrs = append(addrs, msg.Authority)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
