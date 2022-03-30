package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const roomBasePath = "room"

func SetRoomRoutes(router *gin.Engine, roomController controller.RoomController) {
	router.GET(fmt.Sprintf("%s/facility", roomBasePath), roomController.GetAllRoomFacilities)
	router.GET(fmt.Sprintf("%s/:id", roomBasePath), roomController.GetRoom)
	router.POST(fmt.Sprintf("%s", roomBasePath))
}
