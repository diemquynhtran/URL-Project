package routes

import (
	"learning-go/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	urlRepo := controllers.New()

	group := r.Group("/url")
	{
		group.POST("/", urlRepo.CreateUrl)
	}
	
	r.GET("/:short", urlRepo.GetURL)

	return r
}