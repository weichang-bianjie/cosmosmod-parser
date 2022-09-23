package distribution

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// msg struct for validator withdraw
type DocTxMsgWithdrawValidatorCommission struct {
	ValidatorAddress string `bson:"validator_address"`
}

func (m *DocTxMsgWithdrawValidatorCommission) GetType() string {
	return MsgTypeMsgWithdrawValidatorCommission
}

func (m *DocTxMsgWithdrawValidatorCommission) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgWithdrawValidatorCommission)
	m.ValidatorAddress = msg.ValidatorAddress
}

func (m *DocTxMsgWithdrawValidatorCommission) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgWithdrawValidatorCommission)
	addrs = append(addrs, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
