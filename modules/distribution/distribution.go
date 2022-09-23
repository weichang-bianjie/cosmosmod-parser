package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DistributionClient struct {
}

func NewClient() DistributionClient {
	return DistributionClient{}
}

func (distribution DistributionClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgStakeSetWithdrawAddress:
		docMsg := DocTxMsgSetWithdrawAddress{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgWithdrawDelegatorReward:
		docMsg := DocTxMsgWithdrawDelegatorReward{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgWithdrawValidatorCommission:
		docMsg := DocTxMsgWithdrawValidatorCommission{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgFundCommunityPool:
		docMsg := DocTxMsgFundCommunityPool{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
