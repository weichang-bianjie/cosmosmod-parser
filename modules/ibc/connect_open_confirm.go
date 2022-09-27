package ibc

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocMsgConnectionOpenConfirm struct {
	ConnectionId string `bson:"connection_id"`
	ProofAck     string `bson:"proof_ack"`
	ProofHeight  Height `bson:"proof_height"`
	Signer       string `bson:"signer"`
}

func (m *DocMsgConnectionOpenConfirm) GetType() string {
	return MsgTypeConnectionOpenConfirm
}

func (m *DocMsgConnectionOpenConfirm) BuildMsg(v interface{}) {
	msg := v.(*MsgConnectionOpenConfirm)
	m.Signer = msg.Signer
	m.ConnectionId = msg.ConnectionId
	m.ProofAck = utils.MarshalJsonIgnoreErr(msg.ProofAck)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgConnectionOpenConfirm) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgConnectionOpenConfirm)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
