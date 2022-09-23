package gov

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	models "github.com/kaifei-bianjie/common-parser/types"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocTxMsgSubmitProposal struct {
	Proposer       string        `bson:"proposer"`        //  Address of the proposer
	InitialDeposit []models.Coin `bson:"initial_deposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Content        interface{}   `bson:"content"`
}

func (m *DocTxMsgSubmitProposal) GetType() string {
	return MsgTypeSubmitProposal
}

func (m *DocTxMsgSubmitProposal) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgSubmitProposal)
	m.Content = CovertContent(msg.GetContent())
	m.Proposer = msg.Proposer
	m.InitialDeposit = models.BuildDocCoins(msg.InitialDeposit)
}

func CovertContent(content GovContent) interface{} {
	switch content.ProposalType() {
	case ProposalTypeCancelSoftwareUpgrade:
		var data ContentCancelSoftwareUpgradeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeSoftwareUpgrade:
		var data ContentSoftwareUpgradeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeCommunityPoolSpend:
		var data ContentCommunityPoolSpendProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeClientUpdate:
		var data ContentClientUpdateProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeText:
		var data ContentTextProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeParameterChange:
		var data ContentParameterChangeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	}
	return content
}

func (m *DocTxMsgSubmitProposal) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string
	msg := v.(*MsgSubmitProposal)

	content := msg.GetContent()
	if content != nil && ProposalTypeCommunityPoolSpend == content.ProposalType() {
		communityPoolSpend := CovertContent(content).(ContentCommunityPoolSpendProposal)
		addrs = append(addrs, communityPoolSpend.Recipient)
	}
	addrs = append(addrs, msg.Proposer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
