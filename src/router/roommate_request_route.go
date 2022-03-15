package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const roommateRequestRoute = "roommate-request"

func SetRoommateRequestRoutes(router *gin.Engine, roommateRequestController controller.RoommateRequestController) {
	router.GET(fmt.Sprintf("%s/room", roommateRequestRoute), roommateRequestController.GetRoommateRequestsWithRoom)
	router.GET(fmt.Sprintf("%s/no-room", roommateRequestRoute), roommateRequestController.GetRoommateRequestsWithNoRoom)
	router.POST(fmt.Sprintf("%s/no-room", roommateRequestRoute), roommateRequestController.CreateRoommateRequestWithNoRoom)
	router.POST(fmt.Sprintf("%s/registered-dorm", roommateRequestRoute), roommateRequestController.CreateRoommateRequestWithRegisteredDorm)
	router.POST(fmt.Sprintf("%s/unregistered-dorm", roommateRequestRoute), roommateRequestController.CreateRoommateRequestWithUnregisteredDorm)
	router.PUT(fmt.Sprintf("%s/unregistered-dorm/picture", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestWithUnregisteredDormPictures)
	router.PUT(fmt.Sprintf("%s/registered-dorm/picture", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestWithRegisteredDormPictures)
	router.GET(fmt.Sprintf("%s/:id", roommateRequestRoute), roommateRequestController.GetRoommateRequest)
}
