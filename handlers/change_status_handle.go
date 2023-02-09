package handlers

import (
	util "little_robot/pkg/utils"

	"github.com/eatmoreapple/openwechat"
)

func checkIsCanRead(message *openwechat.Message) bool {
	// 通知消息和自己发的不处理
	return !message.IsNotify() && !message.IsSendBySelf()
}

// 设置消息为已读
func setTheMessageAsRead(ctx *openwechat.MessageContext) {
	err := ctx.AsRead()
	if err != nil {
		util.LogrusObj.Errorln("设置消息为已读出错: %v", err)
	}
	ctx.Next()
}
