package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
)

func RecoverInvalidInput(context *gin.Context) {
	if err := recover(); err != nil {
		switch errType := err.(type) {
		case errortype.ErrorMessage:
			context.IndentedJSON(errType.StatusCode, dto.ErrorResponse{
				Message: errType.Error(),
			})

		default:
			context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: fmt.Sprint(errType),
			})

		}
	}
}
