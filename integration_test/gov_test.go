package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/common-parser/codec"
	"github.com/kaifei-bianjie/common-parser/utils"
)

func (s IntegrationTestSuite) TestGov() {
	cases := []SubTest{
		{
			"submitProposal",
			submitProposal,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func submitProposal(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AB6050ACA040A202F636F736D6F732E676F762E76312E4D73675375626D697450726F706F73616C12A5040A85010A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E6412650A2A696161313064303779323635676D6D757674347A30773961773838306A6E73723730306A30716E35357A122A6961613163377933373073766D6D7539796877673766333063666D3730676471756874336C37793277671A0B0A05756E79616E120231300AD0020A172F697269736D6F642E6E66742E4D7367456469744E465412B4020A1F77776E667469646E667469646E667469643131313031323037323030343530121877776E667464656E6F6D69643132303731393439333832391A217777776E6674326E616D656E616D656E616D656E616D6531323038313131353333223B77776269616E6A696261696E6A697878787878787878787878787878787878757269757269757269757269757269757269313230383131313533332A3177776269616E6A696269616E6A697878787878787878787878787878646174617373737373737331323038313131353333322A696161313064303779323635676D6D757674347A30773961773838306A6E73723730306A30716E35357A3A3877776269616E6A696269616E696A78787878787878787572695F686173687572695F686173687572695F6861736831323038313131353333120B0A05756E79616E120231301A2A696161316D756D7672333865647378753471306A7A7436777A61667A6A7A79683239777766396A326134220F3470494D4F6749477831765A47553D12676265697A687561644D656D6F4D656D6F4D656D6F4D656D6F4D656D6F4D656D6F6E6F74656E6F74656E6F74656E6F74656E6F74656E6F74656E6F74656E6F74656466617364667366617364666173646661736466323032322D31322D30385F31313A34373A303012660A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103A2436FB2F383FFA6AEE9D6CB5D73E35F72FD0FCB0802AD3DA398A7F724F68F8C12040A020801183012120A0C0A05756E79616E120338303010C09A0C1A40A99785A6D8CDE2C30A6B1A3CE4EEEBBBAB2E090BEA6B52E0A7D3934DB64FDF1F4285D7814FABA79524B2A4888E94BE724B27D894551C59A33EC7A1511078010D")
	//txBytes, err := hex.DecodeString("0AEB020AA6020A252F636F736D6F732E676F762E763162657461312E4D73675375626D697450726F706F73616C12FC010ABF010A2E2F636F736D6F732E706172616D732E763162657461312E506172616D657465724368616E676550726F706F73616C128C010A1977777777775365727669636520506172616D204368616E676512217777777777557064617465206D6178204D61785265717565737454696D656F75741A4C0A03676F76120D6465706F736974706172616D731A367B226D696E5F6465706F736974223A205B7B22616D6F756E74223A20223130222C2264656E6F6D223A2022756E79616E22207D5D207D120C0A05756E79616E12033130301A2A696161316D756D7672333865647378753471306A7A7436777A61667A6A7A79683239777766396A32613412406265697A68756164666173646661736466617364666161736466617364666173646661736466617364666173646661736466736661736466617364666173646612660A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2103A2436FB2F383FFA6AEE9D6CB5D73E35F72FD0FCB0802AD3DA398A7F724F68F8C12040A020801180812120A0C0A05756E79616E12033830301080B5181A40F5FD45CBF3CBB36664F53C7D64550037920E523F59B0C98BE8FCDE308BBFA776361518BB567DA3E8F42ACAF153C436919EECED9399D22CB9B98F006E537673D2")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Gov.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
