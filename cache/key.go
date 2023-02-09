package cache

import (
	"fmt"
	"strconv"
)

const (
	USER_KEY = "wechat:login:xiangzai"
)

//获取rediskey
func GetViewKey(id uint) string {
	return fmt.Sprintf("views:product:%s", strconv.Itoa(int(id)))
}
