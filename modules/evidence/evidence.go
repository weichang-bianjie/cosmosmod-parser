package evidence

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type EvidenceClient struct {
}

func NewClient() EvidenceClient {
	return EvidenceClient{}
}

func (evidence EvidenceClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgSubmitEvidence:
		docMsg := DocMsgSubmitEvidence{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
