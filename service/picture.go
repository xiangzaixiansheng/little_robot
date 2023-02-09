package service

import (
	"little_robot/pkg/e"
	"little_robot/serializer"
)

type PictureService struct {
}

func (service PictureService) Get() serializer.Response {
	code := e.SUCCESS
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
