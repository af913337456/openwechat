package main

import (
	"errors"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/eatmoreapple/openwechat/xrobot"
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
	xbot        *xrobot.XRobot
	xbotAccount = ""
)

func initCmdFunc() {
	if cmdFunc == nil {
		cmdFunc = CmdFunc{}
	}
	cmdFunc["登录机器人账号"] = func() {
		if account := inputXBotAccount("输入机器人账号"); account == "" {
			return
		} else {
			xbotAccount = account
			if password := inputXBotPassword("输入机器人密码"); password == "" {
				return
			}
			// login
			xbot = xrobot.NewXRobot()
			if err := xbot.Login(account); err != nil {
				fmt.Println("机器人登录失败:", err.Error())
			} else if loginRet, ok := xbot.LoginRetMap.Load(account); loginRet != nil && ok {
				fmt.Println("机器人登录成功")
			}
		}
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
				fmt.Println("文字消息内容:", msg.Content)
				if msg.IsSendBySelf() {
					fmt.Println("自己发的消息不回复")
					return
				}
				if !msg.IsSendByFriend() && !msg.IsSendByGroup() {
					fmt.Println("非好友消息 也 非群组消息，不回复")
					return
				}
				if xbot == nil {
					fmt.Println("没登录机器人，不自动回复")
					return
				}
				fmt.Println("准备请求回复内容...")
				if err := xbot.RequestMsg(xrobot.RequestMsgParam{
					BotAccount: xbotAccount,
					SenderId:   msg.Owner().UserName, // 这里是发送者
					ReceiverId: msg.FromUserName,
					ContentTxt: msg.Content,
					MsgType:    "text",
				}, "desktop-wx", func(data string) {
					fmt.Println("日志:", data)
				}, func(xbotMsg *xrobot.XRobotMsg) {
					fmt.Println("准备微信自动回复...")
					if xbotMsg.MsgType == "text" {
						if _, err := msg.ReplyText(xbotMsg.Content); err != nil {
							fmt.Println("msg.ReplyText 失败:", err.Error())
						} else {
							fmt.Println("微信自动回复成功")
						}
						return
					}
					fmt.Println("不回复非文本内容!", xbotMsg.MsgType)
				}); err != nil {
					fmt.Println("自动回复失败:", err.Error())
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

func inputXBotAccount(tips string) string {
	var result = ""
	validate := func(input string) error {
		if input == "" {
			return errors.New(fmt.Sprintf("机器人账号不能为空"))
		}
		result = input
		return nil
	}
	prompt := promptui.Prompt{
		Label:    tips,
		Validate: validate,
	}
	if _, err := prompt.Run(); err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	return result
}

func inputXBotPassword(tips string) string {
	var result = ""
	validate := func(input string) error {
		if input == "" {
			return errors.New(fmt.Sprintf("机器人密码不能为空"))
		}
		result = input
		return nil
	}
	prompt := promptui.Prompt{
		Label:    tips,
		Validate: validate,
	}
	if _, err := prompt.Run(); err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	return result
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
