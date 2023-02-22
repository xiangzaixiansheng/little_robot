package routes

import (
	api "little_robot/controller/api/v1"
	"little_robot/middleware"
	"little_robot/pkg/e"
	"little_robot/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	v1.Use(middleware.Limiter())

	{

		v1.GET("ping", func(c *gin.Context) {
			code := e.SUCCESS
			c.JSON(200, serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			})
		})

		//获取图片
		v1.GET("picture", api.GetPicture)
		v1.POST("postpicture", api.PostPicture)

		//增加jwt验证
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{

		}

	}
	return r
}
