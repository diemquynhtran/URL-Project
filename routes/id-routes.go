package routes

import (
	"learning-go/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	urlRepo := controllers.New()

	group := r.Group("/url")
	{
		group.POST("/", urlRepo.CreateUrl)
		group.GET("/", urlRepo.GetURL)
	}

	return r
}