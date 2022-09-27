package ibc

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocMsgChannelCloseConfirm struct {
	PortId      string `bson:"port_id"`
	ChannelId   string `bson:"channel_id"`
	ProofInit   string `bson:"proof_init"`
	ProofHeight Height `bson:"proof_height"`
	Signer      string `bson:"signer"`
}

func (m *DocMsgChannelCloseConfirm) GetType() string {
	return MsgTypeChannelCloseConfirm
}

func (m *DocMsgChannelCloseConfirm) BuildMsg(v interface{}) {
	msg := v.(*MsgChannelCloseConfirm)
	m.Signer = msg.Signer
	m.PortId = msg.PortId
	m.ChannelId = msg.ChannelId
	m.ProofInit = utils.MarshalJsonIgnoreErr(msg.ProofInit)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgChannelCloseConfirm) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgChannelCloseConfirm)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
