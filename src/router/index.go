package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

func InitRoutes(router *gin.Engine, controllers Controllers, cacheClient *redis.Client) {
	SetDormRoutes(router, controllers.DormController, cacheClient)
	SetRoomRoutes(router, controllers.RoomController, cacheClient)
	SetAuthRoutes(router, controllers.AuthController)
	SetStudentRoutes(router, controllers.StudentController, cacheClient)
	SetRoommateRequestRoutes(router, controllers.RoommateRequestController, cacheClient)
	SetDormOwnerRoutes(router, controllers.DormOwnerController, cacheClient)
}
