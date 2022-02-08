package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const roommateRequestRoute = "roommate-request"

func SetRoommateRequestRoutes(router *gin.Engine, roommateRequestController controller.RoommateRequestController) {
	router.POST(fmt.Sprintf("%s/no-room", roommateRequestRoute), roommateRequestController.CreateRoomRequestWithNoRoom)
	router.POST(fmt.Sprintf("%s/registered-dorm", roommateRequestRoute), roommateRequestController.CreateRoommateRequestWithRegisteredDorm)
}
