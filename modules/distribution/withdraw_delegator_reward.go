package distribution

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// msg struct for delegation withdraw from a single validator
type DocTxMsgWithdrawDelegatorReward struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
}

func (m *DocTxMsgWithdrawDelegatorReward) GetType() string {
	return MsgTypeWithdrawDelegatorReward
}

func (m *DocTxMsgWithdrawDelegatorReward) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgWithdrawDelegatorReward)
	m.DelegatorAddress = msg.DelegatorAddress
	m.ValidatorAddress = msg.ValidatorAddress
}
func (m *DocTxMsgWithdrawDelegatorReward) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgWithdrawDelegatorReward)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
