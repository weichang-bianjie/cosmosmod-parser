package staking

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	models "github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// MsgDelegate - struct for bonding transactions
type DocTxMsgBeginRedelegate struct {
	DelegatorAddress    string      `bson:"delegator_address"`
	ValidatorSrcAddress string      `bson:"validator_src_address"`
	ValidatorDstAddress string      `bson:"validator_dst_address"`
	Amount              models.Coin `bson:"amount"`
}

func (m *DocTxMsgBeginRedelegate) GetType() string {
	return MsgTypeBeginRedelegate
}

func (m *DocTxMsgBeginRedelegate) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgBeginRedelegate)
	m.DelegatorAddress = msg.DelegatorAddress
	m.ValidatorSrcAddress = msg.ValidatorSrcAddress
	m.ValidatorDstAddress = msg.ValidatorDstAddress
	m.Amount = models.BuildDocCoin(msg.Amount)
}
func (m *DocTxMsgBeginRedelegate) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgBeginRedelegate)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorDstAddress, msg.ValidatorSrcAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
