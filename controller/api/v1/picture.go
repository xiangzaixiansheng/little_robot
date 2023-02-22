package v1

import (
	util "little_robot/pkg/utils"
	"little_robot/service"

	"github.com/gin-gonic/gin"
)

func GetPicture(c *gin.Context) {
	var pictureService service.PictureService
	res := pictureService.Get()
	c.JSON(200, res)
}

func PostPicture(c *gin.Context) {
	var picturePostService service.PictureService
	if err := c.ShouldBind(&picturePostService); err == nil {
		res := picturePostService.Post()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
