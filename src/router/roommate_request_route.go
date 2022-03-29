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
	router.POST(fmt.Sprintf("%s/no-room/:id", roommateRequestRoute), roommateRequestController.CreateRoommateRequestNoRoom)
	router.POST(fmt.Sprintf("%s/registered-dorm/:id", roommateRequestRoute), roommateRequestController.CreateRoommateRequestRegDorm)
	router.POST(fmt.Sprintf("%s/unregistered-dorm/:id", roommateRequestRoute), roommateRequestController.CreateRoommateRequestUnregDorm)
	router.PUT(fmt.Sprintf("%s/unregistered-dorm/:id/picture", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestUnregDormPictures)
	router.PUT(fmt.Sprintf("%s/registered-dorm/:id/picture", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestRegDormPictures)
	router.PUT(fmt.Sprintf("%s/registered-dorm/:id", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestRegDorm)
	router.PUT(fmt.Sprintf("%s/unregistered-dorm/:id", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestUnregDorm)
	router.PUT(fmt.Sprintf("%s/no-room/:id", roommateRequestRoute), roommateRequestController.UpdateRoommateRequestNoRoom)
	router.GET(fmt.Sprintf("%s/:id", roommateRequestRoute), roommateRequestController.GetRoommateRequest)
	router.DELETE(fmt.Sprintf("%s/:id", roommateRequestRoute), roommateRequestController.DeleteRoommateRequest)
}
