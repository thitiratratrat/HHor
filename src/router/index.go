package router

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	SetHelloRoutes(router)
}
