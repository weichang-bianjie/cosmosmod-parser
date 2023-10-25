package msgs

import (
	nfttransfer "github.com/bianjieai/nft-transfer/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/group"
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

	MsgTypeStakeCreateValidator      = "create_validator"
	MsgTypeStakeEditValidator        = "edit_validator"
	MsgTypeStakeDelegate             = "delegate"
	MsgTypeStakeBeginUnbonding       = "begin_unbonding"
	MsgTypeBeginRedelegate           = "begin_redelegate"
	MsgTypeUnjail                    = "unjail"
	MsgTypeCancelUnbondingDelegation = "cancel_unbonding_delegation"

	MsgTypeSetWithdrawAddress             = "set_withdraw_address"
	MsgTypeWithdrawDelegatorReward        = "withdraw_delegator_reward"
	MsgTypeMsgFundCommunityPool           = "fund_community_pool"
	MsgTypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
	MsgTypeSubmitProposal                 = "submit_proposal"
	MsgTypeDeposit                        = "deposit"
	MsgTypeVote                           = "vote"
	MsgTypeVoteWeighted                   = "vote_weighted"
	MsgTypeExecLegacyContent              = "exec_legacy_content"

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

	MsgTypeAuthzRevoke = "revoke"
	MsgTypeAuthzExec   = "exec"
	MsgTypeAuthzGrant  = "grant"

	MsgTypeCreateGroup                     = "create_group"
	MsgTypeUpdateGroupMembers              = "update_group_members"
	MsgTypeUpdateGroupAdmin                = "update_group_admin"
	MsgTypeUpdateGroupMetadata             = "update_group_metadata"
	MsgTypeCreateGroupPolicy               = "create_group_policy"
	MsgTypeCreateGroupWithPolicy           = "create_group_with_policy"
	MsgTypeUpdateGroupPolicyAdmin          = "update_group_policy_admin"
	MsgTypeUpdateGroupPolicyDecisionPolicy = "update_group_policy_decision_policy"
	MsgTypeUpdateGroupPolicyMetadata       = "update_group_policy_metadata"
	MsgTypeGroupSubmitProposal             = "submit_proposal"
	MsgTypeWithdrawProposal                = "withdraw_proposal"
	MsgTypeGroupVote                       = "vote"
	MsgTypeGroupExec                       = "exec"
	MsgTypeLeaveGroup                      = "group"
	//nft transfer
	MsgTypeNftTransfer = "nft_transfer"
)

type (
	MsgSend      = bank.MsgSend
	MsgMultiSend = bank.MsgMultiSend

	MsgStakeCreate               = stake.MsgCreateValidator
	MsgStakeEdit                 = stake.MsgEditValidator
	MsgStakeDelegate             = stake.MsgDelegate
	MsgStakeBeginUnbonding       = stake.MsgUndelegate
	MsgBeginRedelegate           = stake.MsgBeginRedelegate
	MsgCancelUnbondingDelegation = stake.MsgCancelUnbondingDelegation
	StakeValidator               = stake.Validator
	Delegation                   = stake.Delegation
	UnbondingDelegation          = stake.UnbondingDelegation

	MsgUnjail                      = slashing.MsgUnjail
	MsgStakeSetWithdrawAddress     = distribution.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward     = distribution.MsgWithdrawDelegatorReward
	MsgFundCommunityPool           = distribution.MsgFundCommunityPool
	MsgWithdrawValidatorCommission = distribution.MsgWithdrawValidatorCommission

	MsgGrantAllowance  = feegrant.MsgGrantAllowance
	MsgRevokeAllowance = feegrant.MsgRevokeAllowance

	//gov v1bata1
	MsgDeposit        = govv1beta1.MsgDeposit
	MsgSubmitProposal = govv1beta1.MsgSubmitProposal
	TextProposal      = govv1beta1.TextProposal
	MsgVote           = govv1beta1.MsgVote
	Proposal          = govv1beta1.Proposal
	SdkVote           = govv1beta1.Vote
	GovContent        = govv1beta1.Content
	MsgVoteWeighted   = govv1beta1.MsgVoteWeighted

	//gov v1
	MsgSubmitProposalV1  = govv1.MsgSubmitProposal
	MsgExecLegacyContent = govv1.MsgExecLegacyContent
	MsgVoteV1            = govv1.MsgVote
	MsgVoteWeightedV1    = govv1.MsgVoteWeighted
	MsgDepositV1         = govv1.MsgDeposit

	MsgSubmitEvidence  = evidence.MsgSubmitEvidence
	MsgVerifyInvariant = crisis.MsgVerifyInvariant

	NonFungibleTokenIbcPacketData = nfttransfer.NonFungibleTokenPacketData
	NftMsgTransfer                = nfttransfer.MsgTransfer

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

	MsgRevoke            = authz.MsgRevoke
	MsgExec              = authz.MsgExec
	MsgGrant             = authz.MsgGrant
	GenericAuthorization = authz.GenericAuthorization

	MsgCreateGroup                     = group.MsgCreateGroup
	MsgUpdateGroupMembers              = group.MsgUpdateGroupMembers
	MsgUpdateGroupAdmin                = group.MsgUpdateGroupAdmin
	MsgUpdateGroupMetadata             = group.MsgUpdateGroupMetadata
	MsgCreateGroupPolicy               = group.MsgCreateGroupPolicy
	MsgCreateGroupWithPolicy           = group.MsgCreateGroupWithPolicy
	MsgUpdateGroupPolicyAdmin          = group.MsgUpdateGroupPolicyAdmin
	MsgUpdateGroupPolicyDecisionPolicy = group.MsgUpdateGroupPolicyDecisionPolicy
	MsgUpdateGroupPolicyMetadata       = group.MsgUpdateGroupPolicyMetadata
	MsgGroupSubmitProposal             = group.MsgSubmitProposal
	MsgWithdrawProposal                = group.MsgWithdrawProposal
	MsgGroupVote                       = group.MsgVote
	MsgGroupExec                       = group.MsgExec
	MsgLeaveGroup                      = group.MsgLeaveGroup
)
