package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type RoomController interface {
	GetAllRoomFacilities(context *gin.Context)
	GetRoom(context *gin.Context)
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
// @Tags room
// @Produce json
// @Success 200 {array} string "OK"
// @Router /room/facility [get]
func (roomController *roomController) GetAllRoomFacilities(context *gin.Context) {
	facilities := roomController.roomService.GetAllRoomFacilities()

	context.IndentedJSON(http.StatusOK, facilities)
}

// @Summary get room by id
// @Tags room
// @Produce json
// @Success 200 {object} model.Room "OK"
// @Param id path int true "Room ID"
// @Router /room/{id} [get]
func (roomController *roomController) GetRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	id := context.Param("id")
	room := roomController.roomService.GetRoom(id)

	context.IndentedJSON(http.StatusOK, room)
}
