package handlers

import (
	"fmt"
	util "little_robot/pkg/utils"

	"github.com/eatmoreapple/openwechat"
)

// 处理文本消息
func textHandler(ctx *openwechat.MessageContext) {
	sender, _ := ctx.Sender()
	senderUser := sender.NickName
	if ctx.IsSendByGroup() {
		//获取群组里面的发送人信息
		senderInGroup, _ := ctx.SenderInGroup()
		senderUser = fmt.Sprintf("%v[%v]", senderInGroup.NickName, senderUser)
	}

	util.LogrusObj.Infoln("[收到新文字消息] == 发信人：%v ==> 内容：%v", senderUser, ctx.Content)
	ctx.Next()

}
