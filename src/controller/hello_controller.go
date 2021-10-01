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

// @Summary get hello
// @Description returns hello string
// @Tags greeting
// @Produce json
// @Success 200 {string} Hello
// @Router /hello [get]
func (controller *helloController) Hello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hello")
}
