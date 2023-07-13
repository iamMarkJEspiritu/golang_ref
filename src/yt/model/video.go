package model

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10"` // validate:"is-cool"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
