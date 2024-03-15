package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/robfig/cron"
	"github.com/tealeg/xlsx"
)

var replaceNameMap = map[string]string{}

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

//type objItem struct {
//	Protocol   string `json:"protocol"`
//	Statut     string `json:"statut"`
//	Blockchain string `json:"blockchain"`
//	Action     string `json:"action"`
//	Date       string `json:"date"`
//	Site       string `json:"site"`
//	Twitter    string `json:"twitter"`
//	Note       string `json:"note"`
//}

//	func tmpFile(bys []byte)  {
//		// 假设这是你要写入文件的字节数组
//
//		// 打开文件，如果不存在则创建它，使用写模式（O_WRONLY）
//		// 并且设置权限为0644（即文件所有者可以读写，组用户和其他用户可以读取）
//		file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE, 0644)
//		if err != nil {
//			panic(err) // 处理错误，这里简单地使用了panic
//		}
//		defer file.Close() // 确保在函数结束时关闭文件
//
//		// 将字节数组写入文件
//		_, err = file.Write(data)
//		if err != nil {
//			panic(err) // 处理错误
//		}
//	}
const _35dayRelease = 35
const (
	maxNum = 15
	minNum = 8

	oFilePath = "config/o.txt"
	configMp  = "config/config.mp"
)

var airdropToken = ""

func main() {
	go func() {
		fmt.Println("===== 币圈撸羊毛信息爬取程序: v.1.1 ===== ")
		fmt.Println("===== 目标网站: X.com/Twitter、youtube.com、google搜索 =====")
		fmt.Println("")
	RETRY:
		fmt.Println("请输入账号:")
		var account string
		_, err := fmt.Scanf("%s", &account)
		if err != nil {
			fmt.Println("输入错误")
			goto RETRY
		}
		fmt.Println("正在登录..." + account)
		token, err := login(account)
		if err != nil {
			fmt.Println("登录失败:" + err.Error())
			goto RETRY
		}
		airdropToken = token
		fmt.Println("===== 程序运行中 =====")
		go func() {
			list := []AirdropInfo{}
			for {
				time.Sleep(time.Second * 3)
				info, err := LoopAirdropInfo(token)
				if err != nil {
					if err.Error() == "-1" {
						fmt.Printf("*")
					} else {
						fmt.Println("Err:" + err.Error())
					}
				} else {
					if info != nil {
						fmt.Println("")
						fmt.Println(*info)
						list = append(list, *info)
					}
				}
				if size := len(list); size > 10 {
					objList := [][]string{}
					/**
					replaceNameMap["Protocol"] = "项目名称"
					replaceNameMap["Statut"] = "项目现在状态"
					replaceNameMap["Blockchain"] = "项目技术涉及的链"
					replaceNameMap["Action Required"] = "参与要做的动作"
					replaceNameMap["Date"] = "项目现在状态"
					replaceNameMap["Site"] = "项目官网"
					replaceNameMap["Twitter"] = "Twitter"
					*/
					objList = append(objList, []string{
						"项目名称", "项目现在状态", "项目技术涉及的链",
						"参与要做的动作", "项目现在状态", "项目官网", "Twitter", "Telegram"})
					for i := 0; i < size; i++ {
						strList := []string{}
						strList = append(strList, list[i].Protocol, list[i].Statut)
						strList = append(strList, list[i].Blockchain, list[i].Action)
						strList = append(strList, list[i].Date, list[i].Site)
						strList = append(strList, list[i].Twitter, list[i].Telegram)
						objList = append(objList, strList)
					}
					fileName := time.Now().Format("2006_01_02_15_04_05")
					fmt.Println("存储一次:" + fileName)
					createXlsx(fmt.Sprintf("%s.xlsx", fileName), objList)
					list = []AirdropInfo{}
				}
			}
		}()
		/**
		秒      分     小时    日     月      星期     命令
		0-60   0-59   0-23   1-31   1-12     0-6    command     (取值范围,0表示周日一般一行对应一个任务)
		*/
		//NewTimingMission().StartCornMission("0 16 13 * * *", func() {
		//	limitModel()
		//})
	}()
	select {}
}
func apiModel() {

}

//func allModel() {
//	list := run(0, 1000)
//	createXlsx(fmt.Sprintf("%s.xlsx", "1"), list)
//}
//func limitModel() {
//	start, err := validateTimeString("2024-02-09 13:00:00")
//	if err != nil {
//		panic(err)
//	}
//	nowTime := time.Now()
//	leftDay := nowTime.Day() - start.Day()
//	if _35dayRelease < leftDay {
//		// finish
//	} else {
//		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
//		number := rnd.Int63n(maxNum)
//		if number < minNum {
//			number = number + minNum
//		}
//
//		oFile, err := os.OpenFile(oFilePath, os.O_RDWR, 0666)
//		if err != nil {
//			panic(err)
//		}
//		defer oFile.Close()
//
//		offsetLineList := []string{}
//		scanner := bufio.NewScanner(oFile)
//		for scanner.Scan() {
//			line := scanner.Text()
//			offsetLineList = append(offsetLineList, line)
//		}
//
//		var offSet = 0
//		if lastLine := offsetLineList[len(offsetLineList)-1]; lastLine != "" {
//			offSet, err = strconv.Atoi(lastLine)
//			if err != nil {
//				panic(err)
//			}
//		} else {
//			panic("1")
//		}
//		fmt.Println("offSet:", offSet)
//		list := run(offSet, number)
//		if total := len(list); total > 0 {
//			// 使用 golang 实现算法，将217个苹果按照每天最少4个最多7个分配到35天里面
//			_, err := oFile.WriteString(fmt.Sprintf("%d\n", offSet+total-1))
//			if err != nil {
//				panic(err)
//			}
//			fileName := time.Now().Format("2006-01-02-15:04:05")
//			createXlsx(fmt.Sprintf("%s.xlsx", fileName), list)
//			fmt.Println("---------> 保存成功:", fileName)
//		}
//	}
//}
//func run(offset int, itemSize int64) [][]string {
//	// 打开XLSX文件
//	file, err := xlsx.OpenFile(configMp)
//	if err != nil {
//		log.Fatal(err)
//	}
//	const startRowIndex = 3
//	objItemList := [][]string{}
//	skipColMap := map[int]bool{}
//	skipColMap[7] = true
//	// 遍历每个工作表
//	for index, sheet := range file.Sheets {
//		if index != 1 {
//			continue
//		}
//		// 遍历每一行数据
//		for index, row := range sheet.Rows {
//			// 遍历每个单元格数据
//			skip := false
//			objList := []string{}
//			for colIndex, cell := range row.Cells {
//				value := cell.String() // 获取单元格值（字符串类型）
//				if index == startRowIndex {
//					value = replaceNameMap[value]
//				}
//				if offset > 0 && index < startRowIndex+offset+1 && index > startRowIndex {
//					skip = true
//					break
//				}
//				if skipColMap[colIndex] {
//					continue
//				}
//				if colIndex == 1 && strings.Contains(value, "ended") {
//					skip = true
//					break
//				}
//				if index != startRowIndex && colIndex == 5 && !strings.Contains(value, "http") {
//					skip = true
//					break
//				}
//				if value == "" && len([]rune(value)) <= 2 {
//					skip = true
//					break
//				}
//				if strings.HasPrefix(value, "Ended ") ||
//					strings.HasPrefix(value, "End") ||
//					strings.HasPrefix(value, "Fin ") ||
//					strings.HasPrefix(value, "Snapshot") {
//					skip = true
//					break
//				}
//				objList = append(objList, value)
//				//switch i {
//				//case 0:
//				//	obj.Protocol = value
//				//case 1:
//				//	obj.Statut = value
//				//case 2:
//				//	obj.Blockchain = value
//				//case 3:
//				//	obj.Action = value
//				//case 4:
//				//	obj.Date = value
//				//case 5:
//				//	obj.Site = value
//				//case 6:
//				//	obj.Twitter = value
//				//case 7:
//				//	obj.Note = value
//				//}
//				//fmt.Println(index, value) // 输出单元格值
//			}
//			if len(objList) > 3 && !skip {
//				//fmt.Println(objList)
//				if len(objItemList) > int(itemSize) {
//					return objItemList
//				}
//				objItemList = append(objItemList, objList)
//			}
//		}
//	}
//	return objItemList
//}
//
//func validateTimeString(str string) (*time.Time, error) {
//	loc, _ := time.LoadLocation("Local")
//	//使用模板在对应时区转化为time.time类型
//	this, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc) // 根据指定格式解析字符串为时间类型
//	if err != nil {
//		return nil, err
//	}
//	return &this, nil
//}

func createXlsx(fileName string, objItemList [][]string) {
	// 创建一个新的工作簿
	wb := xlsx.NewFile()

	// 创建一个新的工作表
	sheet, err := wb.AddSheet("项目信息")
	if err != nil {
		fmt.Println("createXlsx err:" + err.Error())
		return
	}

	style := xlsx.NewStyle()
	font := *xlsx.NewFont(12, "Verdana")
	font.Bold = true
	sheet.Col(0).Width = 20
	sheet.Col(1).Width = 10
	sheet.Col(2).Width = 20
	sheet.Col(3).Width = 50 // action
	sheet.Col(4).Width = 10
	sheet.Col(5).Width = 30
	sheet.Col(6).Width = 30
	// sheet.Col(6).Width = 30

	for i := 0; i < len(objItemList); i++ {
		// 创建一个新的行
		row := sheet.AddRow()
		for j := 0; j < len(replaceNameMap); j++ {
			// 创建一个新的单元格并设置其值
			//fmt.Println(i, j, objItemList[i])
			cell := row.AddCell()
			cell.SetStyle(style)
			cell.Value = objItemList[i][j]
		}
	}
	// 保存工作簿到文件
	err = wb.Save(fileName)
	if err != nil {
		fmt.Println("save err:" + err.Error())
		return
	}

	//fmt.Println("XLSX file created successfully!")
}

type TimingMission struct{}

type MissionCallbackFunc func()

func NewTimingMission() *TimingMission {
	return &TimingMission{}
}

// DEMO: 0 11 * * 5 , 每周 5 早上11点执行一次
func (t *TimingMission) StartCornMission(spec string, callback MissionCallbackFunc) error {
	c := cron.New()
	if err := c.AddFunc(spec, func() {
		//
		if callback != nil {
			callback()
		}
	}); err != nil {
		return err
	}
	c.Start()
	return nil
}

const host = "http://3.12.161.134:8887"

type commonRet struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type ReqLogin struct {
	Account string `json:"account"`
}
type LoginRet struct {
	Token string `json:"token"`
}

func login(account string) (string, error) {
	url := host + "/Login"
	bys, _ := json.Marshal(ReqLogin{Account: account})
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bys))
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(string(body))
	}
	com := commonRet{}
	_ = json.Unmarshal(body, &com)
	if com.Code != 0 {
		return "", fmt.Errorf(com.Msg)
	}
	log := com.Data
	b, _ := json.Marshal(log)
	ret := LoginRet{}
	_ = json.Unmarshal(b, &ret)
	return ret.Token, nil
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

func LoopAirdropInfo(token string) (*AirdropInfo, error) {
	url := host + "/LoopAirdropInfo"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("token", token)
	req.Header.Set("version", "1.0")
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
	log := com.Data
	b, _ := json.Marshal(log)
	ret := AirdropInfo{}
	_ = json.Unmarshal(b, &ret)
	fmt.Println("response Body:", string(body))
	return &ret, nil
}

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
