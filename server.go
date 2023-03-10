package main

import (
	videoController "main/controller"
	middleware "main/middleware"
	service "main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService            = service.New()
	controller   videoController.VideoController = videoController.New(videoService)
)

func main() {
	// server := gin.Default()
	// =
	// server := gin.New()
	// server.Use(gin.Recovery())
	// server.Use(gin.Logger())

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			//
			ctx.JSON(200, controller.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			//
			err := controller.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", controller.ShowAll)
	}
	server.Run(":8080")
}
