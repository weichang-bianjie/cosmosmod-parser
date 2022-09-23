package staking

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	models "github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// MsgDelegate - struct for bonding transactions
type DocTxMsgDelegate struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
	Amount           Coin   `bson:"amount"`
}

func (m *DocTxMsgDelegate) GetType() string {
	return MsgTypeStakeDelegate
}

func (m *DocTxMsgDelegate) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeDelegate)
	m.ValidatorAddress = msg.ValidatorAddress
	m.DelegatorAddress = msg.DelegatorAddress
	m.Amount = Coin(models.BuildDocCoin(msg.Amount))
}
func (m *DocTxMsgDelegate) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgStakeDelegate)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
