package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/kaifei-bianjie/common-parser/utils"
)

func (s IntegrationTestSuite) TestAuthz() {
	cases := []SubTest{
		{
			"Grant",
			Grant,
		},
		{
			"Exec",
			Exec,
		},
		{
			"Revoke",
			Revoke,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func Grant(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0AD7010AD4010A1D2F636F736D6F732E617574687A2E763162657461312E4D73674578656312B2010A2A69616131337239346B336477786E76356A617839377676337435703830327365666735617738786D61761283010A2E2F697269736D6F642E636F696E737761702E4D736752656D6F7665556E696C61746572616C4C697175696469747912510A0968746C746263626E62120E0A0968746C746263626E621201311A02313020C2A8AFB4062A2A6961613164656578657863367461746C653579786D77616177663766346737796C32723339716477307A12670A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2102B50E62CA4BA4815E47520DF4CCAF39B5BD6792FB52C7E682F6917B9B7836A42F12040A020801180212130A0D0A05756E79616E12043230303010C0843D1A40C1E4A68AC488D65CF149B9E43A42E0548093116A3B6032CEA19B449D41626A4F268C389F6A64EE04AC7D75E269034F5B316FECB2DA1B026BA8C111E3D86DD95A")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if authzDoc, ok := s.Authz.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(authzDoc))
		}
	}
}

func Exec(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0AD4010AD1010A1D2F636F736D6F732E617574687A2E763162657461312E4D73674578656312AF010A2A69616131337239346B336477786E76356A617839377676337435703830327365666735617738786D61761280010A2B2F697269736D6F642E636F696E737761702E4D7367416464556E696C61746572616C4C697175696469747912510A0968746C746263626E62120F0A0968746C746263626E62120231301A013020C2EA86E4062A2A6961613164656578657863367461746C653579786D77616177663766346737796C32723339716477307A12670A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2102B50E62CA4BA4815E47520DF4CCAF39B5BD6792FB52C7E682F6917B9B7836A42F12040A020801180112130A0D0A05756E79616E12043230303010C0843D1A4051F8821ECB270A6B4D766AFAA58FA39DF68572D0CE7D0AD1AC88BF40921022B26200949E940D58379E11DA1913F22108BA8E700861925B8EDA7CC2B7671ED362")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if authzDoc, ok := s.Authz.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(authzDoc))
		}
	}
}

func Revoke(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)

	txBytes, err := hex.DecodeString("0AAF010AAC010A1F2F636F736D6F732E617574687A2E763162657461312E4D73675265766F6B651288010A2A6961613164656578657863367461746C653579786D77616177663766346737796C32723339716477307A122A69616131337239346B336477786E76356A617839377676337435703830327365666735617738786D61761A2E2F697269736D6F642E636F696E737761702E4D736752656D6F7665556E696C61746572616C4C697175696469747912670A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210342FEEC935706ACFA26BC0B39CD05BC36FBB4309B3BD082CCF99891BB2F027B2412040A020801180412130A0D0A05756E79616E12043230303010C0843D1A405FAD6EF375963E4B2BCCB075B250691C0F885F9155CD23CFEC5397809165AD5E5DEF818B03819D9EA7ECA3C4029D5D52A591E6FA5763BB43B2AE9FD7AF76F86A")
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		s.T().Log(err.Error())
		panic(err)
	}

	for _, msg := range authTx.GetMsgs() {
		if authzDoc, ok := s.Authz.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(authzDoc))
		}
	}
}
