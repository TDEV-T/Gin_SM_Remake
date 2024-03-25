package route

import "github.com/gin-gonic/gin"

func SetUpallRoute(rg *gin.RouterGroup) {
	setupUserRoute(rg)
}
