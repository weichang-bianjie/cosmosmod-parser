package staking

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

// MsgEditValidator - struct for editing a validator
type DocMsgEditValidator struct {
	Description       Description `bson:"description"`
	ValidatorAddress  string      `bson:"validator_address"`
	CommissionRate    string      `bson:"commission_rate"`
	MinSelfDelegation string      `bson:"min_self_delegation"`
}

func (m *DocMsgEditValidator) GetType() string {
	return MsgTypeStakeEditValidator
}

func (m *DocMsgEditValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeEdit)
	m.ValidatorAddress = msg.ValidatorAddress
	commissionRate := msg.CommissionRate
	if commissionRate == nil {
		m.CommissionRate = ""
	} else {
		m.CommissionRate = commissionRate.String()
	}
	m.Description = loadDescription(msg.Description)
	if msg.MinSelfDelegation != nil {
		m.MinSelfDelegation = msg.MinSelfDelegation.String()
	}
}
func (m *DocMsgEditValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgStakeEdit)
	addrs = append(addrs, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
