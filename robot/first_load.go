package robot

import (
	"fmt"
	"little_robot/cache"
	. "little_robot/db"
	util "little_robot/pkg/utils"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

func First_load() {
	//当前用户的key 默认只能登陆一个用户
	key := cache.USER_KEY
	//先查询redis中是否有缓存的key
	loadInfo, error := GetRedisInstance().Get(cache.USER_KEY)
	if error != nil {
		util.LogrusObj.Errorln("first_load 获取rediskey失败")
	}
	util.LogrusObj.Infof("[%v]缓存的用户信息是", loadInfo)
	//直接登陆
	bot := InitWechatBot()
	//加载缓存登陆
	if len(loadInfo) > 0 && loadInfo != "" {
		util.LogrusObj.Debugf("当前登陆的用户key: %v", key)
		storage := cache.NewRedisHotReloadStorage(key)
		if err := bot.HotLogin(storage, openwechat.HotLoginWithRetry(false)); err != nil {
			util.LogrusObj.Infof("[%v] 热登录失败，错误信息：%v", key, err.Error())
			if _, err = GetRedisInstance().Del(key); err != nil {
				util.LogrusObj.Errorf("[%v] Redis缓存删除失败，错误信息：%v", key, err.Error())
			}
			return
		}
		loginUser, _ := bot.GetCurrentUser()
		util.LogrusObj.Infof("[%v]初始化自动登录成功，用户名：%v", key, loginUser.NickName)
		cache.SetBot(key, bot)

	} else {
		//直接登陆
		util.LogrusObj.Infof("没有缓存的数据=====第一次登陆")
		bot.ScanCallBack = func(body openwechat.CheckLoginResponse) {
			util.LogrusObj.Infof("已扫码")
		}

		// 设置登录成功回调
		bot.LoginCallBack = func(body openwechat.CheckLoginResponse) {
			util.LogrusObj.Infof("登录成功")
		}
		// 获取登录二维码链接
		bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

		uuid, err := bot.Caller.GetLoginUUID()
		if err != nil {
			util.LogrusObj.Errorf("获取登录二维码失败: %v", err.Error())
			return
		}
		util.LogrusObj.Infof("获取到uuid: %v", uuid)
		// 拼接URL
		url := fmt.Sprintf("https://login.weixin.qq.com/qrcode/%s", uuid)

		util.LogrusObj.Infof("二维码的url: %v", url)

		//打印二维码信息到控制台
		q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
		fmt.Println(q.ToString(true))

		// 设置UUID
		bot.SetUUID(uuid)

		// 定义登录数据缓存
		storage := cache.NewRedisHotReloadStorage(key)

		// 热登录
		var opts []openwechat.BotLoginOption
		opts = append(opts, openwechat.NewRetryLoginOption())                        // 热登录失败使用扫码登录，适配第一次登录的时候无热登录数据
		opts = append(opts, openwechat.NewSyncReloadDataLoginOption(10*time.Minute)) // 十分钟同步一次热登录数据

		// 登录
		if err := bot.HotLogin(storage, opts...); err != nil {
			util.LogrusObj.Errorf("登录失败: %v", err)
			return
		}

		user, err := bot.GetCurrentUser()
		if err != nil {
			util.LogrusObj.Errorf("获取登录用户信息失败: %v", err.Error())
			return
		}
		util.LogrusObj.Infof("当前登录用户：%v", user.NickName)
		cache.SetBot(key, bot)
	}

	//加载用户关系到map中
	self, _ := bot.GetCurrentUser()  // 获取登录用户
	friends, _ := self.Friends(true) // 查找指定的好友
	cache.SetFriendsMap(friends)

}
