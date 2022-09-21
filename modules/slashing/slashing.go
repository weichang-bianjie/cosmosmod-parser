package slashing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type SlashingClient struct {
}

func NewClient() SlashingClient {
	return SlashingClient{}
}

func (slashing SlashingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgUnjail:
		docMsg := DocTxMsgUnjail{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
