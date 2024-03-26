package middleware

import (
	"Gin_Remake/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const ContextKey = "user"

func AuthUserRequire(c *gin.Context) {

	configString := service.Config

	jwttoken := c.GetHeader("authtoken")

	if configString.JWT == "" || jwttoken == "" {
		c.Status(http.StatusUnauthorized)
	}

	token, err := jwt.ParseWithClaims(jwttoken, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configString.JWT), nil
	})

	if err != nil || !token.Valid {
		c.Status(http.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)

	username, ok := claim["username"].(string)

	if !ok || username == "" {
		c.Status(http.StatusInternalServerError)
	}

	exp, ok := claim["exp"].(string)

	if !ok || exp == "" {
		c.Status(http.StatusInternalServerError)
	}

	const layout = "2006-01-02"

	dateNow := time.Now()

	dateNow.Format(layout)

	Expires, err := time.Parse(layout, exp)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	if Expires.After(dateNow) {
		c.Status(http.StatusUnauthorized)
	}

	c.Set(ContextKey, username)

	return

}
