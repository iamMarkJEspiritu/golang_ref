package main

import (
	"io"
	"os"
	
	"yt/controller"
	"yt/middleware"
	"yt/service"

	"github.com/gin-gonic/gin"
)

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func setupLogOutput() {
	log_file, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(log_file, os.Stdout)
}
func main() {
	setupLogOutput()
	server := gin.New()
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		middleware.BasicAuth(),
		//gindump.Dump(),
	)
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, gin.H{"message": "Video added"})
		}
	})
	server.Run(":8080")
}
