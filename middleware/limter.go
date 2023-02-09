package middleware

import (
	"little_robot/db"
	util "little_robot/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		//超过限制则返回429并抛出error
		util.LogrusObj.Infoln("param", "ClientIP", c.ClientIP(), c.RemoteIP())
		if !db.GetRedisInstance().AllowIp(&db.RateLimit{Key: c.ClientIP(), Expire: 10 * time.Second, Max: 5, Decrease: true}) {
			util.LogrusObj.Errorln("too many requests", c.ClientIP(), c.RemoteIP())

			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}
		c.Next()
	}
}
