package main

import (
	route "Gin_Remake/routes"
	"Gin_Remake/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db := service.LoadConnect()

	fmt.Println(db)

	apiGroup := server.Group("/api")

	route.SetUpallRoute(apiGroup)

	server.Run(":8080")
}
