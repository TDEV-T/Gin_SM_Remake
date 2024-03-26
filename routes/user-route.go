package route

import (
	"Gin_Remake/controller"
	"Gin_Remake/middleware"
	"Gin_Remake/service"

	"github.com/gin-gonic/gin"
)

func setupUserRoute(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	var usrservice service.UserService = service.UserServiceFunc()
	var userController = controller.NewUserController(&usrservice)

	users.POST("/register", func(c *gin.Context) { userController.Register(c) })
	users.POST("/login", func(c *gin.Context) { userController.Login(c) })
	users.POST("/current", middleware.AuthUserRequire, func(c *gin.Context) { userController.Current(c) })
}
