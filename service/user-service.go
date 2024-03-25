package service

import (
	"Gin_Remake/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(c *gin.Context)
}

type UserInfo struct {
	models.User
}

func UserServiceFunc() UserService {
	return &UserInfo{}
}

func (u *UserInfo) RegisterUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBind(&user); err != nil {
		c.Status(http.StatusBadGateway)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	fmt.Println(user)

	db := DB_Con

	if db == nil {
		fmt.Println("Error: db is nil")
		c.Status(http.StatusInternalServerError)
		return
	}

	fmt.Println(db)

	if err := db.Save(&user).Error; err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
