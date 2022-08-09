package routes

import (
	"learning-go/cache"
	"learning-go/config"
	"learning-go/controller"
	"learning-go/middleware"
	"learning-go/repository"
	"learning-go/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.InitDb()

	urlRepository  repository.UrlRepository  = repository.NewUrlRepository(db)
	userRepository repository.UserRepository = repository.NewUSerRepository(db)

	urlCache cache.UrlCache = cache.NewUrlCache("localhost:6379", 1, 100)

	authService service.AuthService = service.NewAuthService(userRepository)
	jwtService  service.JWTService  = service.NewJWTService()
	urlService  service.UrlService  = service.NewUrlService(urlRepository)
	userService service.UserService = service.NewUserService(userRepository)

	urlController  controller.UrlController  = controller.NewUrlController(urlService, jwtService, urlCache)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsHandler())
	
	r.GET("/:short", urlController.GetUrlByName)
	r.POST("/url/free", urlController.CreateFreeUrl)
	//r.GET("/:cache", urlController.GetUrlCache)
	//r.GET("/:short", urlController.GetFreeUrl)


	urlRoutes := r.Group("url", middleware.AuthorizeJWT(jwtService))
	{
		urlRoutes.POST("/", urlController.CreateUrl)
		urlRoutes.GET("/urls", urlController.GetUrlByUser)
		urlRoutes.PUT("edit/:urlid", urlController.EditUrl)
		urlRoutes.DELETE("delete/:urlid", urlController.DeleteUrl)

	}
	authRoutes := r.Group("auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/update", userController.Update)
	}

	return r
}
