package distribution

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocTxMsgSetWithdrawAddress struct {
	DelegatorAddress string `bson:"delegator_address"`
	WithdrawAddress  string `bson:"withdraw_address"`
}

func (m *DocTxMsgSetWithdrawAddress) GetType() string {
	return MsgTypeSetWithdrawAddress
}

func (m *DocTxMsgSetWithdrawAddress) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeSetWithdrawAddress)
	m.DelegatorAddress = msg.DelegatorAddress
	m.WithdrawAddress = msg.WithdrawAddress
}

func (m *DocTxMsgSetWithdrawAddress) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgStakeSetWithdrawAddress)
	addrs = append(addrs, msg.DelegatorAddress, msg.WithdrawAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
