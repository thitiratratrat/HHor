package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
	"github.com/thitiratratrat/hhor/src/middleware"
	"github.com/thitiratratrat/hhor/src/service"
)

const roomBasePath = "room"

func SetRoomRoutes(router *gin.Engine, roomController controller.RoomController) {
	roomGroup := router.Group(fmt.Sprintf("/%s", roomBasePath)).Use(middleware.AuthorizeJWT(service.DormOwner))

	router.GET(fmt.Sprintf("%s/facility", roomBasePath), roomController.GetAllRoomFacilities)
	router.GET(fmt.Sprintf("%s/:id", roomBasePath), roomController.GetRoom)

	roomGroup.PUT("/:id/picture", roomController.UpdateRoomPictures)
	roomGroup.PUT("/:id", roomController.UpdateRoom)
	roomGroup.DELETE("/:id", roomController.DeleteRoom)
	roomGroup.POST("", roomController.CreateRoom)
}
