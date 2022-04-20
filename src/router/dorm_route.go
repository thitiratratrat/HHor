package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
	"github.com/thitiratratrat/hhor/src/middleware"
	"github.com/thitiratratrat/hhor/src/service"
)

const dormBasePath = "dorm"

func SetDormRoutes(router *gin.Engine, dormController controller.DormController) {
	dormGroup := router.Group(fmt.Sprintf("/%s", dormBasePath)).Use(middleware.AuthorizeJWT(service.DormOwner))

	router.GET(dormBasePath, dormController.GetDorms)
	router.GET(fmt.Sprintf("%s/suggest/:letter", dormBasePath), dormController.GetDormSuggestions)
	router.GET(fmt.Sprintf("%s/facility", dormBasePath), dormController.GetAllDormFacilities)
	router.GET(fmt.Sprintf("%s/zone", dormBasePath), dormController.GetDormZones)
	router.GET(fmt.Sprintf("%s/:id", dormBasePath), dormController.GetDorm)

	dormGroup.POST("", dormController.CreateDorm)
	dormGroup.PUT("/:id/picture", dormController.UpdateDormPictures)
	dormGroup.PUT("/:id", dormController.UpdateDorm)
	dormGroup.DELETE("/:id", dormController.DeleteDorm)
}
