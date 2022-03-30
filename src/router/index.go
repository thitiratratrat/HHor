package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

type Controllers struct {
	DormController            controller.DormController
	RoomController            controller.RoomController
	AuthController            controller.AuthController
	StudentController         controller.StudentController
	RoommateRequestController controller.RoommateRequestController
	DormOwnerController       controller.DormOwnerController
}

func InitRoutes(router *gin.Engine, controllers Controllers) {
	SetDormRoutes(router, controllers.DormController)
	SetRoomRoutes(router, controllers.RoomController)
	SetAuthRoutes(router, controllers.AuthController)
	SetStudentRoutes(router, controllers.StudentController)
	SetRoommateRequestRoutes(router, controllers.RoommateRequestController)
	SetDormOwnerRoutes(router, controllers.DormOwnerController)
}
