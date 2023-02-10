package robot

import (
	"errors"
	"little_robot/cache"
	util "little_robot/pkg/utils"
)

//发送消息
func SendMsgToFriend(nickName, msg string) error {
	friend := cache.GetFriendMap(nickName)
	if friend != nil {
		if _, err := friend.SendText(msg); err != nil {
			util.LogrusObj.Errorf("发送消息失败: %s", err.Error())
			return err
		}
		return nil
	}
	return errors.New("friend can't find")
}
