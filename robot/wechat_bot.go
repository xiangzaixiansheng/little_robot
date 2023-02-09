package robot

import (
	"little_robot/handlers"

	"github.com/eatmoreapple/openwechat"
)

//初始化wechat登陆
func InitWechatBot() *openwechat.Bot {
	//bot := openwechat.DefaultBot()
	// 桌面模式，上面登录不上的可以尝试切换这种模式
	bot := openwechat.DefaultBot(openwechat.Desktop)

	//注册处理消息的handler
	handlers.HandleMessage(bot)

	return bot
}
