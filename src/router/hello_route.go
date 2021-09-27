package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const basePath = "hello"

func SetHelloRoutes(router *gin.Engine) {
	router.GET(basePath, controller.HelloHandler().Hello)
}
