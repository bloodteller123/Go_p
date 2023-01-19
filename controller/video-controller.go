package controller

import (
	"main/entity"
	service "main/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func (vc *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	vc.service.Save(video)
	return video
}
