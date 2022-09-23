package integration

import (
	"encoding/hex"
	"fmt"
	commoncodec "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/kaifei-bianjie/common-parser/utils"
	cosmoscodec "github.com/kaifei-bianjie/cosmosmod-parser/codec"
)

func (s IntegrationTestSuite) TestFeegrant() {
	cases := []SubTest{
		{
			"GrantAllowance",
			GrantAllowance,
		},
		{
			"RevokeAllowance",
			RevokeAllowance,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func GrantAllowance(s IntegrationTestSuite) {
	cosmoscodec.SetBech32Prefix(cosmoscodec.Bech32PrefixAccAddr, cosmoscodec.Bech32PrefixAccPub, cosmoscodec.Bech32PrefixValAddr,
		cosmoscodec.Bech32PrefixValPub, cosmoscodec.Bech32PrefixConsAddr, cosmoscodec.Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0AB5010AB2010A2A2F636F736D6F732E6665656772616E742E763162657461312E4D73674772616E74416C6C6F77616E63651283010A2A696161313836716874633632636636656A6C7433657277367A6B32386D6777386E6537677766686E3566122A69616131796A71306E32657775666C75656E7979766A32793973656164396A667374707877656C746A701A290A272F636F736D6F732E6665656772616E742E763162657461312E4261736963416C6C6F77616E636512670A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103B4A357463AB84D52BE10E822CD3298CBB050D93D4281392E86A2973A7B6DA29E12040A020801180312130A0D0A04756465761205313030303010C09A0C1A4008F037F973C8DDEE1E07446B34BEA228A4FCACF8829708BF4D746BD565EBEA556C0292124A0908F73543757AE08082E76BD82E0D6209A1F12BF9E951B398DD62")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := commoncodec.GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if feegrantDoc, ok := s.Feegrant.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(feegrantDoc))
		}
	}
}

func RevokeAllowance(s IntegrationTestSuite) {
	cosmoscodec.SetBech32Prefix(cosmoscodec.Bech32PrefixAccAddr, cosmoscodec.Bech32PrefixAccPub, cosmoscodec.Bech32PrefixValAddr,
		cosmoscodec.Bech32PrefixValPub, cosmoscodec.Bech32PrefixConsAddr, cosmoscodec.Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0A8A010A87010A2B2F636F736D6F732E6665656772616E742E763162657461312E4D73675265766F6B65416C6C6F77616E636512580A2A696161313836716874633632636636656A6C7433657277367A6B32386D6777386E6537677766686E3566122A69616131796A71306E32657775666C75656E7979766A32793973656164396A667374707877656C746A7012670A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103B4A357463AB84D52BE10E822CD3298CBB050D93D4281392E86A2973A7B6DA29E12040A020801180412130A0D0A04756465761205313030303010C09A0C1A40238582014C9ED183EF66BAFDB2A9D4B49C5E13023038C51DB42840AAAA31C5593A4139CF71C6F6769392DCABCECD57AE76433C50CA49984795D423DC2DD8586C")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := commoncodec.GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if feegrantDoc, ok := s.Feegrant.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(feegrantDoc))
		}
	}
}
