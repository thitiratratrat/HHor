package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

type Controllers struct {
	DormController controller.DormController
	RoomController controller.RoomController
}

func InitRoutes(router *gin.Engine, controllers Controllers) {
	SetDormRoutes(router, controllers.DormController)
	SetRoomRoutes(router, controllers.RoomController)
}
