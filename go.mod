module github.com/kaifei-bianjie/cosmosmod-parser

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.47.4
	github.com/cosmos/ibc-go/v7 v7.3.0
	github.com/gogo/protobuf v1.3.3
	github.com/kaifei-bianjie/common-parser v0.0.0-20220923023138-65dfc81a8ff5
	github.com/stretchr/testify v1.8.4
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/kaifei-bianjie/common-parser => github.com/weichang-bianjie/common-parser v0.0.0-20231019073333-dec71bc4a8e5
)
