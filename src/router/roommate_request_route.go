package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
	"github.com/thitiratratrat/hhor/src/middleware"
	"github.com/thitiratratrat/hhor/src/service"
)

const roommateReqBasePath = "roommate-request"

func SetRoommateRequestRoutes(router *gin.Engine, roommateRequestController controller.RoommateRequestController) {
	roommateReqGroup := router.Group(fmt.Sprintf("/%s", roommateReqBasePath)).Use(middleware.AuthorizeJWT(service.Student))

	router.GET(fmt.Sprintf("/%s/room", roommateReqBasePath), roommateRequestController.GetRoommateRequestsWithRoom)
	router.GET(fmt.Sprintf("/%s/no-room", roommateReqBasePath), roommateRequestController.GetRoommateRequestsNoRoom)
	router.GET(fmt.Sprintf("/%s/:id", roommateReqBasePath), roommateRequestController.GetRoommateRequest)

	roommateReqGroup.POST("/no-room/:userid", roommateRequestController.CreateRoommateRequestNoRoom)
	roommateReqGroup.POST("/registered-dorm/:userid", roommateRequestController.CreateRoommateRequestRegDorm)
	roommateReqGroup.POST("/unregistered-dorm/:userid", roommateRequestController.CreateRoommateRequestUnregDorm)
	roommateReqGroup.PUT("/unregistered-dorm/:userid/picture", roommateRequestController.UpdateRoommateRequestUnregDormPictures)
	roommateReqGroup.PUT("/registered-dorm/:userid/picture", roommateRequestController.UpdateRoommateRequestRegDormPictures)
	roommateReqGroup.PUT("/registered-dorm/:userid", roommateRequestController.UpdateRoommateRequestRegDorm)
	roommateReqGroup.PUT("/unregistered-dorm/:userid", roommateRequestController.UpdateRoommateRequestUnregDorm)
	roommateReqGroup.PUT("/no-room/:userid", roommateRequestController.UpdateRoommateRequestNoRoom)
	roommateReqGroup.DELETE("/:userid", roommateRequestController.DeleteRoommateRequest)
}
