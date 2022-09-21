package staking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type StakingClient struct {
}

func NewClient() StakingClient {
	return StakingClient{}
}

func (staking StakingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgBeginRedelegate:
		docMsg := DocTxMsgBeginRedelegate{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeBeginUnbonding:
		docMsg := DocTxMsgBeginUnbonding{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeDelegate:
		docMsg := DocTxMsgDelegate{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeEdit:
		docMsg := DocMsgEditValidator{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeCreate:
		docMsg := DocTxMsgCreateValidator{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
