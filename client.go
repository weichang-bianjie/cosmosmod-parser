package cosmosmod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/cosmosmod-parser/codec"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/auth"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/authz"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/bank"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/crisis"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/distribution"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/evidence"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/feegrant"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/gov"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/group"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/ibc"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/params"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/slashing"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/staking"
	"github.com/kaifei-bianjie/cosmosmod-parser/modules/upgrade"
)

type MsgClient struct {
	Auth         Client
	Bank         Client
	Crisis       Client
	Distribution Client
	Evidence     Client
	Feegrant     Client
	Gov          Client
	Params       Client
	Slashing     Client
	Staking      Client
	Upgrade      Client
	Ibc          Client
	Authz        Client
	Group        Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Auth:         auth.NewClient(),
		Bank:         bank.NewClient(),
		Crisis:       crisis.NewClient(),
		Distribution: distribution.NewClient(),
		Evidence:     evidence.NewClient(),
		Feegrant:     feegrant.NewClient(),
		Gov:          gov.NewClient(),
		Params:       params.NewClient(),
		Slashing:     slashing.NewClient(),
		Upgrade:      upgrade.NewClient(),
		Staking:      staking.NewClient(),
		Ibc:          ibc.NewClient(),
		Authz:        authz.NewClient(),
		Group:        group.NewClient(),
	}
}
