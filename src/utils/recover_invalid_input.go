package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/dto"
)

//TODO: add status code as input?
func RecoverInvalidInput(context *gin.Context) {
	if err := recover(); err != nil {
		switch errType := err.(type) {

		case error:
			context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: errType.Error(),
			})

		default:
			context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Message: "bad request",
			})

		}
		return
	}
}
