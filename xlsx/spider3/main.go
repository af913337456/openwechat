package main

// 库不全的话 自己补充 go get
import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"log"
	"regexp"
	"strings"
	"time"
)

var webs = `
https://curve.fi/usecrv
https://debank.com/
https://app.zerion.io/
https://www.tokensets.com/
https://www.rekt.news/
https://app.jarvis.exchange/
https://tinyman.org/
https://app.subsocial.network/claim/dotsama
https://www.angle.money/
https://realt.co/
https://phantom.app/
https://www.dtravel.com/
https://solrazr.com/
https://goldfinch.finance/
https://aztec.network/
https://betafinance.org/
https://beta.horizon.finance/
https://www.gro.xyz/
https://www.unagii.com/
https://ethblock.art/
https://zeta.markets/
https://www.drift.trade/
https://01protocol.com/
https://yield.is/
https://www.teller.finance/
https://www.partybid.app/
https://www.fujidao.org/#/
https://saddle.finance/
https://backd.fund/
https://www.socean.fi/
https://symphony.finance/
https://sablier.finance/
https://www.ante.finance/
https://primitive.finance/
https://qilin.fi/
https://mean.finance/
https://volmex.finance/
https://www.cozy.finance/
https://charm.fi/
https://metapool.app/
https://clipper.exchange/
https://ondo.finance/
https://sherlock.xyz/
https://tinlake.centrifuge.io/
https://www.reaper.farm/
https://vires.finance/
https://parsiq.net/en/
https://boredapeyachtclub.com/#/
https://infinity.xyz/
https://app.ref.finance/
https://www.synfutures.com/
https://www.hashflow.com/
https://shadeprotocol.io/
https://www.rubicon.finance/
https://app.umbra.cash/
https://beta.prysm.xyz/
https://apricot.one/#/
https://pulsemarket.eth.link/#!/
https://tesr.finance/#/
https://jup.ag/
https://stargaze.zone/
https://comdex.one/home
https://fndz.io/
https://app.hawksight.co/
https://kanpeki.finance/
https://airdrop.desmos.network/#roadmap
https://monox.finance/home
https://blizz.finance/
https://zksync.io/
https://growthdefi.com/
https://idex.io/
https://vexchange.io/swap
https://across.to/
https://app.sacred.finance/
https://unstoppabledomains.com/?ref=068e27d9fb8a45b
https://www.sologenic.com/
https://www.praxissociety.com/
https://lum.network/
https://soy.finance/
https://soy.finance/
https://soy.finance/
https://soy.finance/
https://www.instrumental.finance/
https://rhino.fi/
https://evmos.org/
https://prism.ag/
https://sentre.io/#/home
https://tectonic.finance/
https://zks.org/en
https://netswap.io/
https://aladdin.club/
https://www.afterorder.io/
https://curvance.com/
https://www.atrix.finance/
https://www.hydraswap.io/
https://saros.finance/
https://www.jetprotocol.io/
https://sodaprotocol.com/
https://solarisprotocol.com/
https://streamflow.finance/
https://alpha.art/
https://poap.xyz/
https://li.finance/
https://assembly.sc/
https://atlendis.io/
https://www.voltz.xyz/
https://www.earnity.com/
https://bridge.roninchain.com/
https://katana.roninchain.com/#/farm
https://www.neworder.network/
https://www.meanfi.com/
https://prism.ag/
https://prism.ag/
https://www.swaap.finance/
https://www.sandbox.game/
https://giveth.io/
https://oBNBchainurefinance.app/
https://www.spacefi.io/
https://www.optimism.io/
https://www.orbiter.finance/?source=Optimism
https://karma.brahma.fi/
https://nested.fi/
https://starknet.io/what-is-starknet/
https://zora.co/
https://genie.xyz/
https://tezos.domains/
https://www.layerswap.io/ 
https://testnet.aspect.co/ 
https://briq.construction/builder  
https://mintsquare.io/  
https://testnet-app.xbank.finance/ 
https://app.testnet.jediswap.xyz/#/swap
https://www.myswap.xyz/#/ 
https://www.starkswap.co/
https://www.ironfleet.xyz/ 
https://sithswap.com/ 
https://www.brine.finance
https://www.magnety.finance/ 
https://yagi.fi/automation
https://philand.xyz
https://goerli.starkgate.starknet.io/
https://starkgate.starknet.io/
https://twitter.com/ZKEX_Official
https://zksync.trustless.fi/#/ 
https://www.syncswap.xyz/ 
https://zkpad.io/
https://zklend.com/ 
https://www.phezzan.xyz/
https://testnet.app.alpharoad.fi/fr
https://zkx.fi/
https://www.nomad.xyz/
https://debridge.finance/
https://www.bungee.exchange/
https://sturdy.finance/
https://timeswap.io/
https://01.xyz/
https://www.putty.finance/
https://www.cally.finance/
https://avvy.domains/
https://www.topaz.so/explore-collections
https://www.vial.fi/app
https://www.aptosnames.com/
https://liquidswap.com/#/
https://conic.finance/`

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

// 获取网站上爬取的数据
// htmlContent 是上面的 html 页面信息，selector 是我们第一步获取的 selector
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	defer cancel()
	// 执行一个空task, 用提前创建Chrome实例
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为40s  此时间可做更改  调整等待页面加载时间
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 120*time.Second)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitReady(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		//log.Fatal("Run err : %v\n", err)
		return "", err
	}
	//log.Println(htmlContent)

	return htmlContent, nil
}

// 得到具体的数据
// 就是对上面的 html 进行解析，提取我们想要的数据
// 这个 seletor 是在解析后的页面上，定位要哪部分数据
// 要中文 selector 就可以传入 “.chinese” 对应 上面class chinese 部分
// 中英都要，就传入“.item-bottom"
func GetSpecialData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	var str string
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str = selection.Text()
	})
	return str, nil
}
func main() {
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
	for key, value := range tgLinkMap {
		fmt.Println(key, value)
	}
}

func getTgLink() {
	webUrlList := strings.Split(webs, "\n")
	start := false
	fmt.Println("total:", len(webUrlList))
	for i := 0; i < len(webUrlList); i++ {
		if !strings.Contains(webUrlList[i], "http") {
			continue
		}
		if strings.Contains(webUrlList[i], "www.nomad.xyz") {
			start = true
			continue
		}
		if !start {
			fmt.Println("没开始:", webUrlList[i])
			continue
		}
		selector := "body"
		// html 知识 可参考上面链接
		param := `document.querySelector("body")`
		html, err := GetHttpHtmlContent(webUrlList[i], selector, param)
		if err != nil {
			fmt.Println(webUrlList[i] + " err: " + err.Error())
			continue
		}
		reg := regexp.MustCompile(`"https://t.me/[+\da-zA-Z_]*"`)
		tgLink := reg.FindString(html)
		if tgLink == "" {
			secondPar := "https://t.me/joinchat/[\\da-zA-Z_]*"
			reg := regexp.MustCompile(secondPar)
			tgLink = reg.FindString(html)
			if tgLink == "" {
				fmt.Println(webUrlList[i] + ": missing")
				continue
			}
		}
		// 中英都要，就传入“.item-bottom"
		// res, _ := GetSpecialData(html, ".item-bottom")
		fmt.Println(webUrlList[i] + "-->" + tgLink)
	}
	return
}
