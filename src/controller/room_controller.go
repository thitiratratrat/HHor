package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/service"
)

type RoomController interface {
	GetAllRoomFacilities(context *gin.Context)
}

func RoomControllerHandler(roomService service.RoomService) RoomController {
	return &roomController{
		roomService: roomService,
	}
}

type roomController struct {
	roomService service.RoomService
}

// @Summary get room facilities
// @Description returns list of room facilities
// @Tags room
// @Produce json
// @Success 200 {array} string "OK"
// @Router /room/facility [get]
func (roomController *roomController) GetAllRoomFacilities(context *gin.Context) {
	facilities := roomController.roomService.GetAllRoomFacilities()

	context.IndentedJSON(http.StatusOK, facilities)
}
