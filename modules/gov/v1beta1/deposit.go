package v1beta1

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// MsgDeposit
type DocTxMsgDeposit struct {
	ProposalID int64        `bson:"proposal_id"` // ID of the proposal
	Depositor  string       `bson:"depositor"`   // Address of the depositor
	Amount     []types.Coin `bson:"amount"`      // Coins to add to the proposal's deposit
}

func (m *DocTxMsgDeposit) GetType() string {
	return MsgTypeDeposit
}

func (m *DocTxMsgDeposit) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgDeposit)
	m.Depositor = msg.Depositor
	m.Amount = types.BuildDocCoins(msg.Amount)
	m.ProposalID = int64(msg.ProposalId)
}

func (m *DocTxMsgDeposit) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgDeposit)
	addrs = append(addrs, msg.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
