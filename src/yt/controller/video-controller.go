package controller

import (
	"yt/model"
	"yt/service"
	"yt/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type VideoController interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}
