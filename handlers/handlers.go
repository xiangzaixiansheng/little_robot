package handlers

import (
	"github.com/eatmoreapple/openwechat"
)

/**
// 注册消息处理函数
func (m *MessageMatchDispatcher) RegisterHandler(matchFunc matchFunc, handlers ...MessageContextHandler)

// 消息匹配函数
type matchFunc func(*Message) bool

// 消息处理函数
type MessageContextHandler func(ctx *MessageContext)
matchFunc：接受当前收到的消息对象，并返回bool值，返回true则表示处理当前的消息

RegisterHandler：接受一个matchFunc和不定长的消息处理函数，如果matchFunc返回为true，则表示运行对应的处理函数组。

*/

func HandleMessage(bot *openwechat.Bot) {
	//定义一个消息多样化处理器
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// 设置为异步处理
	dispatcher.SetAsync(true)
	// 注册消息处理函数
	// 处理消息为已读
	dispatcher.RegisterHandler(checkIsCanRead, setTheMessageAsRead)

	// 处理文本信息
	dispatcher.OnText(textHandler)

	// 注册消息回调函数
	bot.MessageHandler = dispatcher.AsMessageHandler()
}
