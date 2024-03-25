package route

import (
	"Gin_Remake/service"

	"github.com/gin-gonic/gin"
)

func setupUserRoute(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	var usrservice service.UserService = service.UserServiceFunc()

	users.POST("/register", func(c *gin.Context) { usrservice.RegisterUser(c) })
}
