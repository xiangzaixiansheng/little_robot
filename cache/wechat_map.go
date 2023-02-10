package cache

//缓存wechat的一些信息

import (
	"errors"

	"github.com/eatmoreapple/openwechat"
)

var (
	wechatBotsMap     map[string]*openwechat.Bot //缓存 用户userKey 和 bot 对象的关系
	nickNameFriendMap map[string]*openwechat.Friend
)

func InitGlobalWechatBotMap() {
	wechatBotsMap = make(map[string]*openwechat.Bot)
	nickNameFriendMap = make(map[string]*openwechat.Friend)
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

func GetFriendMap(nickName string) *openwechat.Friend {
	return nickNameFriendMap[nickName]
}

func SetFriendMap(nickName string, friend *openwechat.Friend) {
	nickNameFriendMap[nickName] = friend
}

func ClearFriendMap(nickName string, friend *openwechat.Friend) {
	//直接清空
	nickNameFriendMap = make(map[string]*openwechat.Friend)
}

//批量写入Friends
func SetFriendsMap(friends []*openwechat.Friend) {
	for _, friend := range friends {
		SetFriendMap(friend.NickName, friend)
	}
}
