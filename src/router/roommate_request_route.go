package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const roommateRequestRoute = "roommate-request"

func SetRoommateRequestRoutes(router *gin.Engine, roommateRequestController controller.RoommateRequestController) {
	router.GET(fmt.Sprintf("%s/room", roommateRequestRoute), roommateRequestController.GetRoommateRequestsWithRoom)
	router.GET(fmt.Sprintf("%s/no-room", roommateRequestRoute), roommateRequestController.GetRoommateRequestsNoRoom)
	router.POST(fmt.Sprintf("%s/no-room", roommateRequestRoute), roommateRequestController.CreateRoommateRequestNoRoom)
	router.POST(fmt.Sprintf("%s/registered-dorm", roommateRequestRoute), roommateRequestController.CreateRoommateRequestRegDorm)
	router.POST(fmt.Sprintf("%s/unregistered-dorm", roommateRequestRoute), roommateRequestController.CreateRoommateRequestUnregDorm)
	router.PUT(fmt.Sprintf("%s/unregistered-dorm/picture", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestUnregDormPictures)
	router.PUT(fmt.Sprintf("%s/registered-dorm/picture", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestRegDormPictures)
	router.GET(fmt.Sprintf("%s/:id", roommateRequestRoute), roommateRequestController.GetRoommateRequest)
	router.PUT(fmt.Sprintf("%s/registered-dorm", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestRegDorm)
	router.PUT(fmt.Sprintf("%s/unregistered-dorm", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestUnregDorm)
	router.PUT(fmt.Sprintf("%s/no-room", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestNoRoom)
}
