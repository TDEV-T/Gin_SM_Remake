package service

import (
	"Gin_Remake/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func signJWTUserLogin(u *models.User) string {
	jwtSecret := Config

	if jwtSecret.JWT == "" {
		return ""
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecret.JWT))

	if err != nil {
		fmt.Println(err.Error())

		return ""
	}

	return t
}
