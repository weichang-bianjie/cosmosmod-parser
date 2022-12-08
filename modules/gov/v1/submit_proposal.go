package v1

import (
	"encoding/json"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
	. "github.com/kaifei-bianjie/common-parser/modules"
	models "github.com/kaifei-bianjie/common-parser/types"
	"github.com/kaifei-bianjie/common-parser/utils"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
	"github.com/sirupsen/logrus"
)

type DocTxMsgSubmitProposalV1 struct {
	Messages       interface{}   `bson:"messages"`        //  Address of the proposer
	InitialDeposit []models.Coin `bson:"initial_deposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Proposer       string        `bson:"proposer"`        //  Address of the proposer
	Metadata       string        `bson:"metadata"`
}

func (m *DocTxMsgSubmitProposalV1) GetType() string {
	return MsgTypeSubmitProposal
}

func (m *DocTxMsgSubmitProposalV1) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgSubmitProposalV1)
	messages := make([]interface{}, 0, len(msg.Messages))
	for _, message := range msg.Messages {
		marshalJSON, err := commoncodec.GetMarshaler().MarshalJSON(message)
		if err != nil {
			logrus.Errorf("DocTxMsgSubmitProposalV1 commoncodec.GetMarshaler().MarshalJSON err:%v", err)
		}

		var msgInterface interface{}
		if err = json.Unmarshal(marshalJSON, &msgInterface); err != nil {
			logrus.Errorf("DocTxMsgSubmitProposalV1 json.Unmarshal err:%v", err)
		}

		messages = append(messages, msgInterface)
	}
	m.Messages = messages
	m.InitialDeposit = models.BuildDocCoins(msg.InitialDeposit)
	m.Proposer = msg.Proposer
	m.Metadata = msg.Metadata
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

func (m *DocTxMsgSubmitProposalV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgSubmitProposalV1
	)

	addrs = append(addrs, msg.Proposer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
