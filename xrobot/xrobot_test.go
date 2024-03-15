package xrobot

import (
	"fmt"
	"testing"
	"time"
)

func Test_XRobot(t *testing.T) {
	account := "tg_jujing"
	xbot := NewXRobot()
	if err := xbot.Login(account); err != nil {
		panic(err)
	}
	if err := xbot.RequestMsg(RequestMsgParam{
		BotAccount: account,
		SenderId:   "1",
		ReceiverId: "2",
		ContentTxt: "11",
		MsgType:    "text",
	}, "ins", func(data string) {
		fmt.Println(data)
	}, func(msg *XRobotMsg) {
		fmt.Println(msg)
	}); err != nil {
		fmt.Println(err.Error())
	}
	time.Sleep(time.Minute)
}
