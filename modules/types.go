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
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcconnect "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibcchannel "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
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

	MsgTypeRecvPacket  = "recv_packet"
	MsgTypeIBCTransfer = "transfer"

	//ibc client
	MsgTypeCreateClient             = "create_client"
	MsgTypeUpdateClient             = "update_client"
	MsgTypeUpgradeClient            = "upgrade_client"
	MsgTypeSubmitMisbehaviourClient = "submit_misbehaviour"

	//ibc connect
	MsgTypeConnectionOpenInit    = "connection_open_init"
	MsgTypeConnectionOpenTry     = "connection_open_try"
	MsgTypeConnectionOpenAck     = "connection_open_ack"
	MsgTypeConnectionOpenConfirm = "connection_open_confirm"

	//ibc channel
	MsgTypeChannelOpenInit     = "channel_open_init"
	MsgTypeChannelOpenTry      = "channel_open_try"
	MsgTypeChannelOpenAck      = "channel_open_ack"
	MsgTypeChannelOpenConfirm  = "channel_open_confirm"
	MsgTypeChannelCloseInit    = "channel_close_init"
	MsgTypeChannelCloseConfirm = "channel_close_confirm"
	MsgTypeTimeout             = "timeout_packet"
	MsgTypeTimeoutOnClose      = "timeout_on_close_packet"
	MsgTypeAcknowledgement     = "acknowledge_packet"
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

	FungibleTokenPacketData = ibctransfer.FungibleTokenPacketData
	MsgRecvPacket           = ibc.MsgRecvPacket
	MsgTransfer             = ibctransfer.MsgTransfer

	MsgCreateClient       = ibcclient.MsgCreateClient
	MsgUpdateClient       = ibcclient.MsgUpdateClient
	MsgSubmitMisbehaviour = ibcclient.MsgSubmitMisbehaviour
	MsgUpgradeClient      = ibcclient.MsgUpgradeClient

	MsgConnectionOpenInit    = ibcconnect.MsgConnectionOpenInit
	MsgConnectionOpenAck     = ibcconnect.MsgConnectionOpenAck
	MsgConnectionOpenConfirm = ibcconnect.MsgConnectionOpenConfirm
	MsgConnectionOpenTry     = ibcconnect.MsgConnectionOpenTry

	Acknowledgement        = ibc.Acknowledgement
	MsgChannelOpenInit     = ibcchannel.MsgChannelOpenInit
	MsgChannelOpenTry      = ibcchannel.MsgChannelOpenTry
	MsgChannelOpenAck      = ibcchannel.MsgChannelOpenAck
	MsgChannelOpenConfirm  = ibcchannel.MsgChannelOpenConfirm
	MsgChannelCloseConfirm = ibcchannel.MsgChannelCloseConfirm
	MsgChannelCloseInit    = ibcchannel.MsgChannelCloseInit
	MsgAcknowledgement     = ibcchannel.MsgAcknowledgement
	MsgTimeout             = ibcchannel.MsgTimeout
	MsgTimeoutOnClose      = ibcchannel.MsgTimeoutOnClose
)
