package global

import (
	"errors"

	"github.com/eatmoreapple/openwechat"
)

//缓存 用户userKey 和 bot 对象的关系
var (
	wechatBotsMap map[string]*openwechat.Bot
)

func InitGlobalWechatBotMap() {
	wechatBotsMap = make(map[string]*openwechat.Bot)
}

func GetBot(userKey string) *openwechat.Bot {
	return wechatBotsMap[userKey]
}

func SetBot(userKey string, bot *openwechat.Bot) {
	wechatBotsMap[userKey] = bot
}

func CheckBot(userKey string) error {
	bot := GetBot(userKey)
	if nil == bot {
		return errors.New("未获取到登录记录")
	}
	// 判断在线状态是否正常
	if !bot.Alive() {
		return errors.New("微信在线状态异常，请重新登录")
	}
	return nil
}
