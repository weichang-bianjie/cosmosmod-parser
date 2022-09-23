package msgs

import (
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stake "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	MsgTypeSend      = "send"
	MsgTypeMultiSend = "multisend"

	MsgTypeGrantAllowance  = "grant_allowance"
	MsgTypeRevokeAllowance = "revoke_allowance"

	MsgTypeStakeCreateValidator = "create_validator"
	MsgTypeStakeEditValidator   = "edit_validator"
	MsgTypeStakeDelegate        = "delegate"
	MsgTypeStakeBeginUnbonding  = "begin_unbonding"
	MsgTypeBeginRedelegate      = "begin_redelegate"
	MsgTypeUnjail               = "unjail"

	MsgTypeSetWithdrawAddress             = "set_withdraw_address"
	MsgTypeWithdrawDelegatorReward        = "withdraw_delegator_reward"
	MsgTypeMsgFundCommunityPool           = "fund_community_pool"
	MsgTypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
	MsgTypeSubmitProposal                 = "submit_proposal"
	MsgTypeDeposit                        = "deposit"
	MsgTypeVote                           = "vote"
	MsgTypeVoteWeighted                   = "vote_weighted"

	MsgTypeSubmitEvidence  = "submit_evidence"
	MsgTypeVerifyInvariant = "verify_invariant"
)

type (
	MsgSend      = bank.MsgSend
	MsgMultiSend = bank.MsgMultiSend

	MsgStakeCreate         = stake.MsgCreateValidator
	MsgStakeEdit           = stake.MsgEditValidator
	MsgStakeDelegate       = stake.MsgDelegate
	MsgStakeBeginUnbonding = stake.MsgUndelegate
	MsgBeginRedelegate     = stake.MsgBeginRedelegate
	StakeValidator         = stake.Validator
	Delegation             = stake.Delegation
	UnbondingDelegation    = stake.UnbondingDelegation

	MsgUnjail                      = slashing.MsgUnjail
	MsgStakeSetWithdrawAddress     = distribution.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward     = distribution.MsgWithdrawDelegatorReward
	MsgFundCommunityPool           = distribution.MsgFundCommunityPool
	MsgWithdrawValidatorCommission = distribution.MsgWithdrawValidatorCommission

	MsgGrantAllowance  = feegrant.MsgGrantAllowance
	MsgRevokeAllowance = feegrant.MsgRevokeAllowance

	MsgDeposit        = gov.MsgDeposit
	MsgSubmitProposal = gov.MsgSubmitProposal
	TextProposal      = gov.TextProposal
	MsgVote           = gov.MsgVote
	Proposal          = gov.Proposal
	SdkVote           = gov.Vote
	GovContent        = gov.Content
	MsgVoteWeighted   = gov.MsgVoteWeighted

	MsgSubmitEvidence  = evidence.MsgSubmitEvidence
	MsgVerifyInvariant = crisis.MsgVerifyInvariant
)
