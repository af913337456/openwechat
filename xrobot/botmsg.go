package xrobot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const BotHost = "54-169-199-196"

func BotTokenKey(account string) string {
	return fmt.Sprintf("bot_token_%s", account)
}

type User struct {
	UserId int64 `xorm:"pk autoincr" json:"user_id"`
}
type LoginRet struct {
	Data  User   `json:"data"`
	Token string `json:"token"`
}

func DoLogin(account string) (*LoginRet, error) {
	url := "http://ec2-" + BotHost + ".ap-southeast-1.compute.amazonaws.com:8887/InternalLogin"
	// fmt.Println("url:>", url)

	reqParam := struct {
		Account string
	}{
		Account: account,
	}
	bys, err := json.Marshal(struct {
		Data string `json:"data"`
	}{
		Data: AES.AesEncryptBean(reqParam),
	})
	if err != nil {
		return nil, errors.New("marshal reqParam err: " + err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bys))
	if err != nil {
		return nil, err
	}

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
	fmt.Println("response Body:", string(body))

	respJson := LoginRet{}
	respBean := &LghOutput{}
	if err := json.Unmarshal(body, respBean); err != nil {
		return nil, errors.New("Unmarshal resp err:" + err.Error() + "\n" + string(body))
	}
	if respBean.IsSuccess() {
		dataStr := (*respBean)["data"].(string)
		bys := []byte(AES.AesDecryptStr(dataStr))
		if err := json.Unmarshal(bys, &respJson); err != nil {
			return nil, errors.New("Unmarshal replyRet err:" + err.Error() + "\n" + string(bys))
		}
		return &respJson, nil
	} else {
		return nil, errors.New("request failed:" + respBean.String())
	}
}

type CallReplyServerParam struct {
	BotLoginToken  string `json:"bot_login_token"`
	BotLoginUserId int64  `json:"bot_login_user_id"`
	SenderId       string `json:"sender_id"`
	ReceiverId     string `json:"receiver_id"`
	ContentTxt     string `json:"content_txt"`
	MsgType        string `json:"msg_type"`
}

type XRobotMsg struct {
	MsgId      int64  `xorm:"pk autoincr" json:"msg_id"`
	RuleId     int64  `xorm:"index" json:"rule_id"`
	UserId     int64  `xorm:"default 0" json:"user_id"`
	Score      int64  `json:"score"` // 分数
	Title      string `json:"title"`
	LogoUrl    string `json:"logo_url"`
	Content    string `json:"content"`
	MsgType    string `json:"msg_type"`
	MediaUrl   string `json:"media_url"` // 多媒体消息，这个是链接
	Delay      int64  `json:"delay"`     // 延迟回复的秒数
	Backup     string `json:"backup"`
	CreateType string `xorm:"default ''" json:"create_type"`
	IsDelete   bool   `xorm:"default 0" json:"is_delete"`
	Addtime    int64  `json:"addtime"`
	Attach     string `json:"attach"`
}

type OutputMsg struct {
	Msg                   XRobotMsg   `json:"msg"`
	MsgList               []XRobotMsg `json:"msg_list"`
	RuleType              string      `json:"rule_type"`
	IsTop                 bool        `json:"is_top"`
	IsNeedReply           bool        `json:"is_need_reply"`
	Index                 int64       `json:"index"`
	KeepIndex             bool        `json:"keep_index"`
	ParentRuleId          int64       `json:"parent_rule_id"`
	KeyWordReply          bool        `json:"key_word_reply"`
	IsSpecialModelKeyWord bool        `json:"is_special_model_key_word"`
	VirtualFileName       string      `json:"virtual_file_name"`
	VirtualVideoResp      bool        `json:"virtual_video_resp"`
	NeedTranslate         bool        `json:"need_translate"`
}

type InputMsgObj struct {
	UserId  int64  `json:"user_id"`
	AppWxId string `json:"app_wx_id"` // 发消息方的软件 id
	WxId    string `json:"wx_id"`     // 自己的软件 id
	// Index            int64  `json:"index"`
	Content          string `json:"content"`
	IsVoiceRecognize bool   `json:"is_voice_recognize"` // 是否是语音识别
	ParentRuleId     int64  `json:"parent_rule_id"`
	RuleType         string `json:"rule_type"`
	MsgType          string `json:"msg_type"`
	AddTime          int64  `json:"add_time"`
	IsBlock          bool   `json:"is_block"`  // 当关键词被挡住的时候，且服务端也有消息的，也回复
	UserType         int    `json:"user_type"` // wx 自动回复的拓展，分 platform 来源，共有：1 2 3
	Platform         string `json:"platform"`
	PhoneNumber      string `json:"phone_number"`
	Country          string `json:"country"`
	CheckRepeat      bool   `json:"check_repeat"`
	AlreadyVideo     bool   `json:"already_video"`
	OpenVideoFunc    bool   `json:"open_video_func"`
	VideoIng         bool   `json:"video_ing"`
	OnlyType         int64  `json:"only_type"`
}

type InputMsg struct {
	Current InputMsgObj `json:"current"`
	//LastMsg InputMsgObj `json:"last_msg"`
}

type ReqReplyMsg struct {
	Msg InputMsg `json:"msg"`
}

func CallReplyServer(param CallReplyServerParam, platform string, logCallback func(data string)) (*OutputMsg, error) {
	//if needTranslate {
	//	if err := util.NetErrRetry(func() error {
	//		// 翻译转中文去请求
	//		resp, err := translate.GoogleTranslateTo_Cn(param.ContentTxt)
	//		if err != nil {
	//			return errors.New("请求->翻译失败:" + err.Error())
	//		} else if resp != nil && len(resp.Data.Text) > 0 {
	//			newContent := resp.Data.Text[0]
	//			fmt.Println(fmt.Sprintf("请求->翻译: %s --> %s", param.ContentTxt, newContent))
	//			param.ContentTxt = newContent
	//		} else {
	//			return fmt.Errorf("请求->翻译失败: 没有数据 resp.Data.Text==0, %s", param.ContentTxt)
	//		}
	//		return nil
	//	}, 3); err != nil {
	//		return nil, err
	//	}
	//}
	replyServerUrl := "http://ec2-" + BotHost + ".ap-southeast-1.compute.amazonaws.com:8887/ReplyMsg"
	if param.MsgType == "" {
		param.MsgType = "text"
	}
	reqParam := ReqReplyMsg{
		Msg: InputMsg{Current: InputMsgObj{
			UserId:   param.BotLoginUserId,
			AppWxId:  param.ReceiverId,
			WxId:     param.SenderId,
			Content:  param.ContentTxt,
			MsgType:  param.MsgType,
			AddTime:  time.Now().Unix(),
			IsBlock:  false,
			Platform: platform,
		}},
	}
	bys, err := json.Marshal(struct {
		Data string `json:"data"`
	}{
		Data: AES.AesEncryptBean(reqParam),
	})
	req, err := http.NewRequest("POST", replyServerUrl, bytes.NewReader(bys))
	if err != nil {
		return nil, errors.New("new request err: " + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", param.BotLoginToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("read body err:" + err.Error())
	}
	respBean := &LghOutput{}
	if err := json.Unmarshal(body, respBean); err != nil {
		return nil, errors.New("Unmarshal resp err:" + err.Error() + "\n" + string(body))
	}
	if respBean.IsSuccess() {
		dataStr := (*respBean)["data"].(string)
		bys := []byte(AES.AesDecryptStr(dataStr))
		replyRet := &OutputMsg{}
		if err := json.Unmarshal(bys, replyRet); err != nil {
			return nil, errors.New("Unmarshal replyRet err:" + err.Error() + "\n" + string(bys))
		}
		// 处理翻译转英文返回
		//if size := len(replyRet.MsgList); size > 0 && replyRet.Msg.MsgId == 0 {
		//	for i := 0; i < size; i++ {
		//		enResp, enErr := translate.GoogleTranslateTo_En(replyRet.MsgList[i].Content)
		//		if enErr != nil {
		//			logStr := fmt.Sprintf("返回->翻译错误，批量: %s", enErr.Error())
		//			fmt.Println(logStr)
		//			logCallback(logStr)
		//		} else {
		//			replyRet.MsgList[i].Content = enResp.Data.Text[0]
		//		}
		//	}
		//} else if replyRet.Msg.MsgId > 0 {
		//	enResp, enErr := translate.GoogleTranslateTo_En(replyRet.Msg.Content)
		//	if enErr != nil {
		//		logStr := fmt.Sprintf("返回->翻译错误，单条: %s", enErr.Error())
		//		fmt.Println(logStr)
		//		logCallback(logStr)
		//	} else {
		//		replyRet.Msg.Content = enResp.Data.Text[0]
		//	}
		//}
		return replyRet, nil
	} else {
		return nil, errors.New("request failed:" + respBean.String())
	}
}

type LghOutput map[string]interface{}

func (l LghOutput) IsSuccess() bool {
	if l["ret"] == "success" {
		return true
	}
	return false
}
func (l LghOutput) String() string {
	bys, _ := json.Marshal(l)
	return string(bys)
}
func (l LghOutput) ToError() error {
	bys, _ := json.Marshal(l)
	return errors.New(string(bys))
}

type KeyId struct {
	Id string `xorm:"pk" json:"id,omitempty"`
}
