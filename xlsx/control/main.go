package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var replaceNameMap = map[string]string{}

var data2 = `https://curve.fi/usecrv-->"https://t.me/curvefi"
https://debank.com/: missing
https://app.zerion.io/: missing
https://www.tokensets.com/: missing
https://www.rekt.news/: missing
https://app.jarvis.exchange/ err: page load error net::ERR_CONNECTION_CLOSED
https://tinyman.org/-->"https://t.me/tinymanofficial"
https://app.subsocial.network/claim/dotsama: missing
https://www.angle.money/: missing
https://realt.co/-->"https://t.me/Realtoken_welcome"
https://phantom.app/: missing
https://www.dtravel.com/-->"https://t.me/dtravelcommunity"
https://solrazr.com/-->"https://t.me/solrazr_community"
https://goldfinch.finance/-->"https://t.me/goldfinch_finance"
https://aztec.network/: missing
https://betafinance.org/: missing
https://beta.horizon.finance/ err: page load error net::ERR_CONNECTION_CLOSED
https://www.gro.xyz/: missing
https://www.unagii.com/-->"https://t.me/unagiidotcom"
https://ethblock.art/: missing
https://zeta.markets/: missing
https://www.drift.trade/: missing
https://01protocol.com/: missing
https://yield.is/: missing
https://www.teller.finance/-->"https://t.me/tellerofficial"
https://www.partybid.app/: missing
https://www.fujidao.org/#/: missing
https://saddle.finance/-->"https://t.me/saddle_announcements"
https://backd.fund/-->"https://t.me/meroannouncements"
https://www.socean.fi/-->"https://t.me/soceaneers"
https://symphony.finance/: missing
https://sablier.finance/: missing
https://www.ante.finance/: missing
https://primitive.finance/: missing
https://qilin.fi/-->"https://t.me/QilinOfficial"
https://mean.finance/: missing
https://volmex.finance/-->"https://t.me/VolmexUpdates"
https://www.cozy.finance/: missing
https://charm.fi/-->"https://t.me/charmfinance"
https://metapool.app/-->"https://t.me/MetaPoolOfficialGroup"
https://clipper.exchange/: missing
https://ondo.finance/-->"https://t.me/ondofinance"
https://sherlock.xyz/: missing
https://tinlake.centrifuge.io/: missing
https://www.reaper.farm/-->"https://t.me/reaperfarm"
https://vires.finance/-->"https://t.me/vf_announcements"
https://parsiq.net/en/-->"https://t.me/parsiq_group"
https://boredapeyachtclub.com/#/: missing
https://infinity.xyz/ err: page load error net::ERR_CONNECTION_CLOSED
https://app.ref.finance/-->"https://t.me/ref_finance"
https://www.synfutures.com/: missing
https://www.hashflow.com/-->"https://t.me/hashflowdex"
https://shadeprotocol.io/-->"https://t.me/ShadeProtocol"
https://www.rubicon.finance/: missing
https://app.umbra.cash/: missing
https://beta.prysm.xyz/ err: page load error net::ERR_CONNECTION_CLOSED
https://apricot.one/#/-->"https://t.me/ApricotOfficial"
https://pulsemarket.eth.link/#!/-->"https://t.me/pulsemarkets"
https://tesr.finance/#/-->"https://t.me/tesseractFinanceANN"
https://jup.ag/: missing
https://stargaze.zone/: missing
https://comdex.one/home-->"https://t.me/ComdexChat"
https://fndz.io/: missing
https://app.hawksight.co/ err: page load error net::ERR_CONNECTION_CLOSED
https://kanpeki.finance/: missing
https://airdrop.desmos.network/#roadmap err: page load error net::ERR_CONNECTION_CLOSED
https://monox.finance/home: missing
https://blizz.finance/: missing
https://zksync.io/: missing
https://growthdefi.com/-->"https://t.me/growthdefi"
https://idex.io/-->"https://t.me/IDEXChat"
https://vexchange.io/swap: missing
https://across.to/: missing
https://app.sacred.finance/ err: page load error net::ERR_CONNECTION_CLOSED
https://unstoppabledomains.com/?ref=068e27d9fb8a45b: missing
https://www.sologenic.com/-->"https://t.me/SOLOGENICxGoSOLO"
https://www.praxissociety.com/-->"https://t.me/+1llT8zSSTWAzNzJh"
https://lum.network/-->"https://t.me/lum_network"
https://soy.finance/-->"https://t.me/Soy_Finance"
https://www.instrumental.finance/-->"https://t.me/instrumentalfinance"
https://rhino.fi/: missing
https://evmos.org/-->"https://t.me/EvmosOrg"
https://prism.ag/: missing
https://sentre.io/#/home-->"https://t.me/Sentre"
https://tectonic.finance/-->"https://t.me/TectonicOfficial"
https://zks.org/en: missing
https://netswap.io/-->"https://t.me/Netswap_Announcement_CH"
https://aladdin.club/-->"https://t.me/aladdin_dao"
https://www.afterorder.io/: missing
https://curvance.com/: missing
https://www.atrix.finance/-->"https://t.me/AtrixProtocol"
https://www.hydraswap.io/ err: page load error net::ERR_CONNECTION_CLOSED
https://saros.finance/-->"https://t.me/saros_finance"
https://www.jetprotocol.io/: missing
https://sodaprotocol.com/-->"https://t.me/SodaProtocol"
https://solarisprotocol.com/ err: page load error net::ERR_TIMED_OUT
https://streamflow.finance/: missing
https://alpha.art/: missing
https://poap.xyz/: missing
https://li.finance/-->"https://t.me/lifinews"
https://assembly.sc/: missing
https://atlendis.io/: missing
https://www.voltz.xyz/: missing
https://www.earnity.com/: missing
https://bridge.roninchain.com/: missing
https://katana.roninchain.com/#/farm: missing
https://www.neworder.network/: missing
https://www.meanfi.com/: missing
https://prism.ag/: missing
https://prism.ag/: missing
https://www.swaap.finance/: missing
https://www.sandbox.game/-->"https://t.me/sandboxgame"
https://giveth.io/: missing
https://oBNBchainurefinance.app/ err: page load error net::ERR_CONNECTION_CLOSED
https://www.spacefi.io/: missing
https://www.optimism.io/: missing
https://www.orbiter.finance/?source=Optimism: missing
https://karma.brahma.fi/: missing
https://nested.fi/-->"https://t.me/NestedOfficialChannel"
https://starknet.io/what-is-starknet/: missing
https://zora.co/: missing
https://genie.xyz/ err: page load error net::ERR_CERT_COMMON_NAME_INVALID
https://tezos.domains/: missing
https://www.layerswap.io/ : missing
https://testnet.aspect.co/ -->"https://t.me/aspectco"
https://briq.construction/builder  : missing
https://mintsquare.io/  : missing
https://testnet-app.xbank.finance/ : missing
https://app.testnet.jediswap.xyz/#/swap: missing
https://www.myswap.xyz/#/-->"https://t.me/mySwapxyz"
https://www.starkswap.co/-->"https://t.me/+lDFExB27O6o0YjU0"
https://www.ironfleet.xyz/  err: page load error net::ERR_CONNECTION_CLOSED
https://sithswap.com/ -->"https://t.me/SithWars"
https://www.brine.finance: missing
https://www.magnety.finance/ : missing
https://yagi.fi/automation: missing
https://philand.xyz: missing
https://goerli.starkgate.starknet.io/: missing
https://starkgate.starknet.io/: missing
https://twitter.com/ZKEX_Official: missing
https://zksync.trustless.fi/#/  err: page load error net::ERR_CONNECTION_RESET
https://www.syncswap.xyz/ : missing
https://zkpad.io/ err: page load error net::ERR_CONNECTION_CLOSED
https://zklend.com/ -->"https://t.me/zkLendAnnouncements"
https://www.phezzan.xyz/ err: page load error net::ERR_CONNECTION_CLOSED
https://testnet.app.alpharoad.fi/fr-->"https://t.me/alpharoad_fi"
https://zkx.fi/-->"https://t.me/zkxcommunity"
https://www.nomad.xyz/ err: page load error net::ERR_CONNECTION_RESET
https://debridge.finance/-->"https://t.me/deBridge_finance"
https://www.bungee.exchange/ err: page load error net::ERR_CONNECTION_RESET
https://sturdy.finance/-->"https://t.me/sturdyfinance"
https://timeswap.io/-->"https://t.me/timeswap"
https://01.xyz/: missing
https://www.putty.finance/: missing
https://www.cally.finance/ err: page load error net::ERR_CONNECTION_RESET
https://avvy.domains/-->"https://t.me/avvydomains"
https://www.topaz.so/explore-collections: missing
https://www.vial.fi/app err: page load error net::ERR_CONNECTION_RESET
https://www.aptosnames.com/ err: page load error net::ERR_CONNECTION_RESET
https://liquidswap.com/#/-->"https://t.me/pontemnetworkchat"
https://conic.finance/: missing`

func init() {
	replaceNameMap["Protocol"] = "项目名称"
	replaceNameMap["Statut"] = "项目现在状态"
	replaceNameMap["Blockchain"] = "项目技术涉及的链"
	replaceNameMap["Action Required"] = "参与要做的动作"
	replaceNameMap["Date"] = "项目现在状态"
	replaceNameMap["Site"] = "项目官网"
	replaceNameMap["Twitter"] = "Twitter"
	replaceNameMap["Telegram"] = "Telegram"
	// replaceNameMap["Note"] = "备注"
}

func getTgLinkMap() map[string]string {
	webUrlList := strings.Split(data2, "\n")
	tgLinkMap := map[string]string{}
	for _, item := range webUrlList {
		if arr := strings.Split(item, "-->"); len(arr) == 2 {
			tgLinkMap[arr[0]] = strings.Replace(arr[1], "\"", "", -1)
			continue
		}
		if strings.Contains(item, "missing") {
			arr := strings.Split(item, ": ")
			tgLinkMap[arr[0]] = "dont have telegram"
			continue
		}
		if strings.Contains(item, "err:") {
			arr := strings.Split(item, "err:")
			tgLinkMap[strings.Replace(arr[0], " ", "", -1)] = "dont have telegram"
			continue
		}
	}
	if len(tgLinkMap) == 0 {
		panic("len(tgLinkMap) == 0")
	}
	return tgLinkMap
}

func main() {
	list := run()
	fmt.Println("总条数:", len(list))
	start := false
	tgLinkMap := getTgLinkMap()
	for i := 0; i < len(list); i++ {
		if list[i].Protocol == "项目名称" {
			continue
		}
		if list[i].Protocol == "Conic Finance" {
			start = true
			continue
		}
		if !start {
			fmt.Println("没开始:", list[i].Protocol)
			continue
		}
		list[i].Telegram = tgLinkMap[list[i].Site]
		if list[i].Telegram == "" && !strings.Contains(list[i].Site, "https://chrome.google.com") {
			tmp := strings.Replace(list[i].Site, " ", "", -1)
			list[i].Telegram = tgLinkMap[tmp]
			if list[i].Telegram == "" {
				panic(fmt.Sprintf("empty link: %s [%s]", list[i].Protocol, list[i].Site))
			}
		}
		//number := 3
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		number := rnd.Int63n(60 * 5 * 6) // 40 min
		if number < 60*3*2 {
			number = 60 * 3 * 2
		}
		fmt.Println(fmt.Sprintf("等待: %d 秒", number))
		time.Sleep(time.Second * time.Duration(number))
		fmt.Println(list[i])
		ret, err := PushAirdropInfo("push_qw12iouuynnfs00990009s", list[i])
		fmt.Println("插入:", ret, err)
	}

}

type commonRet struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type AirdropInfo struct {
	Protocol   string `json:"protocol"`
	Statut     string `json:"statut"`
	Blockchain string `json:"blockchain"`
	Action     string `json:"action"`
	Date       string `json:"date"`
	Site       string `json:"site"`
	Twitter    string `json:"twitter"`
	Telegram   string `json:"telegram"`
}

func (a AirdropInfo) Valid() bool {
	if a.Site == "" || a.Date == "" || a.Twitter == "" || a.Action == "" ||
		a.Blockchain == "" || a.Protocol == "" || a.Statut == "" {
		return false
	}
	return true
}

const host = "http://3.12.161.134:8887"

func PushAirdropInfo(token string, info AirdropInfo) (*commonRet, error) {
	url := host + "/PushAirdropInfo"
	bys, _ := json.Marshal(info)
	req, err := http.NewRequest("POST", url, bytes.NewReader(bys))
	if err != nil {
		return nil, err
	}
	req.Header.Set("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}
	com := commonRet{}
	_ = json.Unmarshal(body, &com)
	if com.Code != 0 {
		return nil, fmt.Errorf(com.Msg)
	}
	// fmt.Println("login response Body:", string(body))
	return &com, nil
}

const configMp = "config/config.mp"

func run() []AirdropInfo {
	// 打开XLSX文件
	file, err := xlsx.OpenFile(configMp)
	if err != nil {
		panic(err)
	}
	const startRowIndex = 3
	objItemList := []AirdropInfo{}
	skipColMap := map[int]bool{}
	skipColMap[7] = true
	// 遍历每个工作表
	for index, sheet := range file.Sheets {
		if index != 1 {
			continue
		}
		// 遍历每一行数据
		for index, row := range sheet.Rows {
			// 遍历每个单元格数据
			skip := false
			airdrop := AirdropInfo{}
			for colIndex, cell := range row.Cells {
				value := cell.String() // 获取单元格值（字符串类型）
				if index == startRowIndex {
					value = replaceNameMap[value]
				}
				//if offset > 0 && index < startRowIndex+offset+1 && index > startRowIndex {
				//	skip = true
				//	break
				//}
				if skipColMap[colIndex] {
					continue
				}
				if colIndex == 1 && strings.Contains(value, "ended") {
					skip = true
					break
				}
				if index != startRowIndex && colIndex == 5 && !strings.Contains(value, "http") {
					skip = true
					break
				}
				if value == "" && len([]rune(value)) <= 2 {
					skip = true
					break
				}
				if strings.HasPrefix(value, "Ended ") ||
					strings.HasPrefix(value, "End") ||
					strings.HasPrefix(value, "Fin ") ||
					strings.HasPrefix(value, "Snapshot") {
					skip = true
					break
				}
				switch colIndex {
				case 0:
					airdrop.Protocol = value
				case 1:
					airdrop.Statut = value
				case 2:
					airdrop.Blockchain = value
				case 3:
					airdrop.Action = value
				case 4:
					airdrop.Date = value
				case 5:
					airdrop.Site = value
				case 6:
					airdrop.Twitter = value
				case 7:
					airdrop.Telegram = value
					//case 7:
					//	airdrop.Note = value
				}
				//fmt.Println(index, value) // 输出单元格值
			}
			if !skip && airdrop.Valid() {
				objItemList = append(objItemList, airdrop)
			}
		}
	}
	return objItemList
}
