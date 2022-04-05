package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/service"
)

type DormOwnerController interface {
	GetDormOwner(*gin.Context)
}

func DormOwnerControllerHandler(dormOwnerService service.DormOwnerService) DormOwnerController {
	return &dormOwnerController{
		dormOwnerService: dormOwnerService,
	}
}

type dormOwnerController struct {
	dormOwnerService service.DormOwnerService
}

// @Summary get dorm owner profile
// @Tags dorm-owner
// @Produce json
// @Param id path int true "Dorm Owner ID"
// @Success 200 {object} model.DormOwner "OK"
// @Router /dorm-owner/{id} [get]
func (dormOwnerController *dormOwnerController) GetDormOwner(context *gin.Context) {
	dormOwnerID := context.Param("id")

	dormOwner := dormOwnerController.dormOwnerService.GetDormOwner(dormOwnerID)

	context.IndentedJSON(http.StatusOK, dormOwner)
}
