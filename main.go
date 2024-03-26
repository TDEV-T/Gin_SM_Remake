package main

import (
	route "Gin_Remake/routes"
	"Gin_Remake/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	env := service.LoadConfig()
	db := service.LoadConnect()

	fmt.Println(db, env)

	apiGroup := server.Group("/api")

	route.SetUpallRoute(apiGroup)

	server.Run(":8080")
}
