package staking

import (
	stake "github.com/cosmos/cosmos-sdk/x/staking/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
	"github.com/kaifei-bianjie/cosmosmod-parser/utils"
)

// MsgCreateValidator defines an SDK message for creating a new validator.
type DocTxMsgCreateValidator struct {
	Description       Description     `bson:"description"`
	Commission        CommissionRates `bson:"commission"`
	MinSelfDelegation string          `bson:"min_self_delegation"`
	DelegatorAddress  string          `bson:"delegator_address"`
	ValidatorAddress  string          `bson:"validator_address"`
	Pubkey            string          `bson:"pubkey"`
	Value             Coin            `bson:"value"`
}

func (m *DocTxMsgCreateValidator) GetType() string {
	return MsgTypeStakeCreateValidator
}

func (m *DocTxMsgCreateValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeCreate)
	if msg.Pubkey != nil {
		m.Pubkey = utils.MarshalJsonIgnoreErr(msg.Pubkey)
	}

	m.ValidatorAddress = msg.ValidatorAddress
	m.DelegatorAddress = msg.DelegatorAddress
	m.MinSelfDelegation = msg.MinSelfDelegation.String()
	m.Commission = CommissionRates{
		Rate:          msg.Commission.Rate.String(),
		MaxChangeRate: msg.Commission.MaxChangeRate.String(),
		MaxRate:       msg.Commission.MaxRate.String(),
	}
	m.Description = loadDescription(msg.Description)
	m.Value = Coin{Denom: msg.Value.Denom, Amount: msg.Value.Amount.String()}
}
func (m *DocTxMsgCreateValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgStakeCreate)

	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

func loadDescription(description stake.Description) Description {
	return Description{
		Moniker:         description.Moniker,
		Details:         description.Details,
		Identity:        description.Identity,
		Website:         description.Website,
		SecurityContact: description.SecurityContact,
	}
}
