package service

import (
	"little_robot/pkg/e"
	"little_robot/serializer"
)

type PictureService struct {
	UserName string `json:"user_name"`
	Password int    `json:"password"`
}

//测试get方法
func (service PictureService) Get() serializer.Response {
	code := e.SUCCESS
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 测试post方法
func (service PictureService) Post() serializer.Response {
	code := e.SUCCESS
	user_name := service.UserName
	passowrd := service.Password

	return serializer.Response{
		Status: code,
		Data:   &PictureService{UserName: user_name, Password: passowrd},
		Msg:    e.GetMsg(code),
	}
}
