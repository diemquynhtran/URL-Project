package main

import (
	"learning-go/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/id", controllers.CreateID)

	router.Run(":3000")
}