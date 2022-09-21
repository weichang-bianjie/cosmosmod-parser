package distribution

import (
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
	models "github.com/kaifei-bianjie/cosmosmod-parser/types"
)

// msg struct for delegation withdraw for all of the delegator's delegations
type DocTxMsgFundCommunityPool struct {
	Amount    []models.Coin `bson:"amount"`
	Depositor string        `bson:"depositor"`
}

func (m *DocTxMsgFundCommunityPool) GetType() string {
	return MsgTypeMsgFundCommunityPool
}

func (m *DocTxMsgFundCommunityPool) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgFundCommunityPool)
	m.Depositor = msg.Depositor
	m.Amount = models.BuildDocCoins(msg.Amount)
}
func (m *DocTxMsgFundCommunityPool) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgFundCommunityPool)
	addrs = append(addrs, msg.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
