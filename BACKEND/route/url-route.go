package routes

import (
	"learning-go/config"
	"learning-go/controller"
	"learning-go/middleware"
	"learning-go/repository"
	"learning-go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db            *gorm.DB                 = config.InitDb()
	urlRepository repository.UrlRepository = repository.NewUrlRepository(db)
	urlService    service.UrlService       = service.NewUrlService(urlRepository)
	urlController controller.UrlController = controller.NewUrlController(urlService)
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsHandler())
	urlRoutes := r.Group("url")
	{
		urlRoutes.POST("/", urlController.CreateUrl)
	}
	r.GET("/:short", urlController.GetUrlByName)

	return r
}
