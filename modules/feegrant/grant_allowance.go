package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/gogo/protobuf/proto"
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocTxMsgGrantAllowance struct {
	Granter             string               `bson:"granter"`
	Grantee             string               `bson:"grantee"`
	Allowance           *Allowance           `bson:"allowance"`
	PeriodicAllowance   *PeriodicAllowance   `bson:"periodic_allowance"`
	AllowedMsgAllowance *AllowedMsgAllowance `bson:"allowed_msg_allowance"`
}

type (
	Allowance struct {
		SpendLimit []types.Coin `bson:"spend_limit"`
		Expiration int64        `bson:"expiration"`
	}
	PeriodicAllowance struct {
		Basic            Allowance    `bson:"basic"`
		Period           int64        `bson:"period"`
		PeriodSpendLimit []types.Coin `bson:"period_spend_limit"`
		PeriodCanSpend   []types.Coin `bson:"period_can_spend"`
		PeriodReset      int64        `bson:"period_reset"`
	}

	AllowedMsgAllowance struct {
		Allowance       *Allowance `json:"allowance"`
		AllowedMessages []string   `json:"'allowed_messages'"`
	}
)

func (m *DocTxMsgGrantAllowance) GetType() string {
	return MsgTypeGrantAllowance
}

func (m *DocTxMsgGrantAllowance) BuildMsg(v interface{}) {
	msg := v.(*MsgGrantAllowance)
	m.Granter = msg.Granter
	m.Grantee = msg.Grantee

	if msg.Allowance.Value != nil {
		switch msg.Allowance.GetTypeUrl() {
		case "/" + proto.MessageName(&feegrant.BasicAllowance{}):
			a := getBasicAllowance(msg.Allowance.Value)
			m.Allowance = &a
		case "/" + proto.MessageName(&feegrant.PeriodicAllowance{}):
			var period feegrant.PeriodicAllowance
			_ = proto.Unmarshal(msg.Allowance.Value, &period)

			var expiration int64
			if period.Basic.Expiration != nil {
				expiration = period.Basic.Expiration.Unix()
			}
			m.PeriodicAllowance = &PeriodicAllowance{
				Basic: Allowance{
					SpendLimit: types.BuildDocCoins(period.Basic.SpendLimit),
					Expiration: expiration,
				},
				Period:           int64(period.Period.Seconds()),
				PeriodSpendLimit: types.BuildDocCoins(period.PeriodSpendLimit),
				PeriodCanSpend:   types.BuildDocCoins(period.PeriodCanSpend),
				PeriodReset:      period.PeriodReset.Unix(),
			}
		case "/" + proto.MessageName(&feegrant.AllowedMsgAllowance{}):
			var allowMsg feegrant.AllowedMsgAllowance
			_ = proto.Unmarshal(msg.Allowance.Value, &allowMsg)

			res := AllowedMsgAllowance{
				Allowance:       nil,
				AllowedMessages: allowMsg.AllowedMessages,
			}

			if allowMsg.Allowance.GetTypeUrl() == "/"+proto.MessageName(&feegrant.BasicAllowance{}) {
				a := getBasicAllowance(allowMsg.Allowance.Value)
				res.Allowance = &a
			}
			m.AllowedMsgAllowance = &res
		}
	}
}

func getBasicAllowance(protoValue []byte) Allowance {
	var basic feegrant.BasicAllowance
	_ = proto.Unmarshal(protoValue, &basic)

	var expiration int64
	if basic.Expiration != nil {
		expiration = basic.Expiration.Unix()
	}

	return Allowance{
		SpendLimit: types.BuildDocCoins(basic.SpendLimit),
		Expiration: expiration,
	}
}

func (m *DocTxMsgGrantAllowance) HandleTxMsg(v sdk.Msg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgGrantAllowance)
	addrs = append(addrs, msg.Granter, msg.Grantee)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
