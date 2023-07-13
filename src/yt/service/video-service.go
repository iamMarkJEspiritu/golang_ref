package service

import "yt/model"

type VideoService interface {
	Save(model.Video) model.Video
	FindAll() []model.Video
}

type videoService struct {
	videos []model.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video model.Video) model.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []model.Video {
	return service.videos
}
