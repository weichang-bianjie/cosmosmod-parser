package v1

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// MsgDeposit
type DocTxMsgDepositV1 struct {
	ProposalID int64        `bson:"proposal_id"` // ID of the proposal
	Depositor  string       `bson:"depositor"`   // Address of the depositor
	Amount     []types.Coin `bson:"amount"`      // Coins to add to the proposal's deposit
}

func (m *DocTxMsgDepositV1) GetType() string {
	return MsgTypeDeposit
}

func (m *DocTxMsgDepositV1) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgDepositV1)
	m.Depositor = msg.Depositor
	m.Amount = types.BuildDocCoins(msg.Amount)
	m.ProposalID = int64(msg.ProposalId)
}

func (m *DocTxMsgDepositV1) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgDepositV1)
	addrs = append(addrs, msg.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
