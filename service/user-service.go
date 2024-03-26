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
	LoginUser(c *gin.Context)
	CurrentUser(c *gin.Context)
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

func (u *UserInfo) LoginUser(c *gin.Context) {
	userBodyParser := &models.User{}

	if err := c.ShouldBind(&userBodyParser); err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	db := LoadConnect()

	if db == nil {
		fmt.Println("Error: db is nil")
		c.Status(http.StatusInternalServerError)
		return
	}

	selectedUser := &models.User{}

	result := db.Where("username = ?", userBodyParser.Username).First(selectedUser)

	if result.Error != nil {
		var msg string

		if result.Error.Error() == "record not found" {
			msg = "username or password is wrong"
		} else {
			msg = result.Error.Error()
		}

		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": msg})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(userBodyParser.Password)); err != nil {

		var msg string

		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			msg = "Username or Password Incorrect !"
		}

		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": msg})

		return
	}

	token := signJWTUserLogin(userBodyParser)

	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": "Server Is Die Ahhhhhhhhhhhhh"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"status": http.StatusOK,
		"msg":    "Login Success",
	})

	return

}

func (u *UserInfo) CurrentUser(c *gin.Context) {
	username, ok := c.Get("user")

	if !ok {
		c.Status(http.StatusInternalServerError)
	}

	db := LoadConnect()

	if db == nil {
		fmt.Println("Error: db is nil")
		c.Status(http.StatusInternalServerError)
		return
	}

	selectedUser := &models.User{}

	result := db.Where("username = ?", username).First(selectedUser)

	if result.Error != nil {
		var msg string

		if result.Error.Error() == "record not found" {
			msg = "username or password is wrong"
		} else {
			msg = result.Error.Error()
		}

		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": msg})

		return
	}

	usernameString := fmt.Sprintf("%v", username)

	token := signJWTUserLogin(&models.User{Username: usernameString})

	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "msg": "Server Is Die Ahhhhhhhhhhhhh"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"status": http.StatusOK,
		"msg":    "Verify Success",
	})

}
