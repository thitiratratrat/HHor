package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController interface {
	Hello(context *gin.Context)
}

func HelloHandler() HelloController {
	return &helloController{}
}

type helloController struct{}

func (controller *helloController) Hello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hello")
}
