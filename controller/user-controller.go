package controller

import (
	"Gin_Remake/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Current(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(uservice *service.UserService) UserController {
	return &userController{
		userService: *uservice,
	}
}

func (s *userController) Register(c *gin.Context) {
	s.userService.RegisterUser(c)
}

func (s *userController) Login(c *gin.Context) {
	s.userService.LoginUser(c)
}

func (s *userController) Current(c *gin.Context) {
	s.userService.CurrentUser(c)
}
