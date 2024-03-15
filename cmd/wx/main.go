package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/manifoldco/promptui"
	"os"
)

type CmdFunc map[string]func()

var cmdFunc CmdFunc

func (c CmdFunc) GetCmdNames() []string {
	names := []string{}
	for name, _ := range c {
		names = append(names, name)
	}
	return names
}

func firstSelector() {
	prompts := promptui.Select{
		HideHelp: true,
		Size:     len(cmdFunc),
		Label:    "微信黑客工具: 请使用 ↓ ↑ → ← 在下面选择功能👇",
		Items:    cmdFunc.GetCmdNames(),
	}
	_, result, err := prompts.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("您选择了: %q\n", result)
	cmdFunc[result]()
	secondSelector()
}

var (
	xbotAccount = ""
)

func initCmdFunc() {
	if cmdFunc == nil {
		cmdFunc = CmdFunc{}
	}
	cmdFunc["登录微信"] = func() {
		bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
		// 注册消息处理函数
		bot.MessageHandler = func(msg *openwechat.Message) {
			if msg.IsText() {
				if msg.Content == "ping" {
					_, _ = msg.ReplyText("pong")
					return
				}
				fmt.Println("用户发来文字消息内容:", msg.Content)
				if msg.IsSendBySelf() {
					fmt.Println("自己发的消息不回复")
					return
				}
				if !msg.IsSendByFriend() && !msg.IsSendByGroup() {
					fmt.Println("非好友消息 也 非群组消息，不回复")
					return
				}
				fmt.Println("准备请求回复内容...")
				if msg.IsSendByFriend() {
					fmt.Println("私聊信息")
				} else if msg.IsSendByGroup() {
					fmt.Println("群聊信息")
				}
			} else {
				fmt.Println("不支持的消息类型:", msg.String())
			}
		}
		// 注册登陆二维码回调
		bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

		// 登陆
		if err := bot.Login(); err != nil {
			fmt.Println(err)
			return
		}

		// 获取登陆的用户
		self, err := bot.GetCurrentUser()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 获取所有的好友
		friends, err := self.Friends()
		fmt.Println(friends, err)

		// 获取所有的群组
		groups, err := self.Groups()
		fmt.Println(groups, err)

		// 阻塞主goroutine, 直到发生异常或者用户主动退出
		bot.Block()
	}
	cmdFunc["退出"] = func() {
		os.Exit(1)
	}
}

func secondSelector() {
	prompts := promptui.Select{
		HideHelp: true,
		Size:     len(cmdFunc),
		Label:    "请继续选择功能👇",
		Items:    cmdFunc.GetCmdNames(),
	}
	_, result, err := prompts.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("您选择了: %q\n", result)
	cmdFunc[result]()
}

func main() {
	initCmdFunc()
	firstSelector()
}
