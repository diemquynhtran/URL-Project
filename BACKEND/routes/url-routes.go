package routes

import (
	"learning-go/controllers"
	"learning-go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsHandler())
	urlRepo := controllers.New()
	r.GET("/:short", urlRepo.GetURL)
	r.POST("/url/", urlRepo.CreateUrl)   // r.POST("/", urlRepo.CreateUrl)

	return r
}

