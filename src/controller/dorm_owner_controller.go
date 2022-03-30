package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/service"
)

type DormOwnerController interface {
	GetDorms(*gin.Context)
}

func DormOwnerControllerHandler(dormOwnerService service.DormOwnerService) DormOwnerController {
	return &dormOwnerController{
		dormOwnerService: dormOwnerService,
	}
}

type dormOwnerController struct {
	dormOwnerService service.DormOwnerService
}

// @Summary get dorm owner's dorms
// @Tags dorm-owner
// @Produce json
// @Param id path int true "Dorm Owner ID"
// @Success 200 {array} string "OK"
// @Router /dorm-owner/{id}/dorm [get]
func (dormOwnerController *dormOwnerController) GetDorms(context *gin.Context) {
	dormOwnerID := context.Param("id")

	dorms := dormOwnerController.dormOwnerService.GetDorms(dormOwnerID)

	context.IndentedJSON(http.StatusOK, dorms)
}
