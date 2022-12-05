package staking

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	models "github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocTxMsgCancelUnbondingDelegation struct {
	DelegatorAddress string      `bson:"delegator_address"`
	ValidatorAddress string      `bson:"validator_address"`
	Amount           models.Coin `bson:"amount"`
	CreationHeight   int64       `bson:"creation_height"`
}

func (m *DocTxMsgCancelUnbondingDelegation) GetType() string {
	return MsgTypeCancelUnbondingDelegation
}

func (m *DocTxMsgCancelUnbondingDelegation) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgCancelUnbondingDelegation)
	m.ValidatorAddress = msg.ValidatorAddress
	m.DelegatorAddress = msg.DelegatorAddress
	m.Amount = models.BuildDocCoin(msg.Amount)
	m.CreationHeight = msg.CreationHeight
}
func (m *DocTxMsgCancelUnbondingDelegation) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCancelUnbondingDelegation)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
