package xrobot

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type XRobotMsgType string

const (
	XRobotMsgType_Text  XRobotMsgType = "text"
	XRobotMsgType_Image XRobotMsgType = "image"
	XRobotMsgType_Voice XRobotMsgType = "voice"
	XRobotMsgType_Video XRobotMsgType = "video"
)

type XRobot struct {
	LoginRetMap sync.Map
}

func NewXRobot() *XRobot {
	return &XRobot{
		LoginRetMap: sync.Map{},
	}
}

func (x *XRobot) Login(account string) error {
	ret, err := DoLogin(account)
	if err != nil {
		return fmt.Errorf("login failed: %s", err.Error())
	}
	x.LoginRetMap.Store(account, ret)
	return nil
}

type RequestMsgParam struct {
	BotAccount string        `json:"bot_account"` // 账号
	SenderId   string        `json:"sender_id"`   // 发送者id
	ReceiverId string        `json:"receiver_id"` // 接收消息者id
	ContentTxt string        `json:"content_txt"` // 发送的内容
	MsgType    XRobotMsgType `json:"msg_type"`    // 消息类型
}

func (x *XRobot) RequestMsg(param RequestMsgParam, platform string, logFunc func(data string), HandleMsg func(msg *XRobotMsg)) error {
	ret, ok := x.LoginRetMap.Load(param.BotAccount)
	if !ok || ret == nil {
		return fmt.Errorf("请先调用 Login 接口登录: [%s] 账号", param.BotAccount)
	}
	resp, err := CallReplyServer(CallReplyServerParam{
		BotLoginToken:  ret.(*LoginRet).Token,
		BotLoginUserId: ret.(*LoginRet).Data.UserId,
		SenderId:       platform + param.SenderId,
		ReceiverId:     platform + param.ReceiverId,
		ContentTxt:     param.ContentTxt,
		MsgType:        string(param.MsgType),
	}, platform, logFunc)
	if err != nil {
		return fmt.Errorf("获取回复内容失败, ，内容：[%s]\nerr: %s", param.ContentTxt, err.Error())
	} else {
		go ClearOlderQueue()
		queueKey := mapKey(param.SenderId, param.ReceiverId)
		existQueue := GetQueue(queueKey)
		if len(resp.MsgList) > 0 && resp.Msg.MsgId == 0 {
			// 批量回复
			logFunc(fmt.Sprintf("批量回复"))
			for i := 0; i < len(resp.MsgList); i++ {
				if delaySec := resp.Msg.Delay; delaySec > 0 {
					logFunc(fmt.Sprintf("当前回复有秒数延迟，延迟 %d 秒回复", delaySec))
				}
				existQueue.Enqueue(&QueueItem{ReplyMsg: resp.MsgList[i]})
			}
		} else {
			// 单条回复
			b, _ := json.Marshal(resp.Msg)
			logFunc(fmt.Sprintf("单条回复，机器人服务器消息：%s", string(b)))
			if delaySec := resp.Msg.Delay; delaySec > 0 {
				logFunc(fmt.Sprintf("当前回复有秒数延迟，延迟 %d 秒回复", delaySec))
			}
			// http://vjs.zencdn.net/v/oceans.mp4
			existQueue.Enqueue(&QueueItem{ReplyMsg: resp.Msg})
		}
		// 启动当前用户的回复队列
		GetOnce(queueKey, true).Do(func() {
			go func() {
				for {
					time.Sleep(time.Millisecond * 30)
					if replyItem := existQueue.Dequeue(); replyItem != nil {
						if delaySec := replyItem.ReplyMsg.Delay; delaySec > 0 {
							time.Sleep(time.Second * time.Duration(delaySec)) // 延迟
						}
						if HandleMsg != nil {
							HandleMsg(&replyItem.ReplyMsg)
						}
						// startReply(replyItem.DXToken, replyItem.Btn, replyItem.BotMsg, replyItem.DXMsg)
					}
				}
			}()
		})
	}
	return nil
}
