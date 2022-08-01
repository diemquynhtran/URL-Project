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
	db *gorm.DB = config.InitDb()

	urlRepository  repository.UrlRepository  = repository.NewUrlRepository(db)
	userRepository repository.UserRepository = repository.NewUSerRepository(db)

	authService service.AuthService = service.NewAuthService(userRepository)
	jwtService  service.JWTService  = service.NewJWTService()
	urlService  service.UrlService  = service.NewUrlService(urlRepository)
	userSerivce service.UserService = service.NewUserService(userRepository)

	urlController  controller.UrlController  = controller.NewUrlController(urlService)
	userController controller.UserController = controller.NewUserController(userSerivce, jwtService)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsHandler())
	
	r.GET("/:short", urlController.GetUrlByName)

	urlRoutes := r.Group("url")
	{
		urlRoutes.POST("/", urlController.CreateUrl)
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
