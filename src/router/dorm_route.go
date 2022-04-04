package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const dormBasePath = "dorm"

func SetDormRoutes(router *gin.Engine, dormController controller.DormController) {
	router.GET(dormBasePath, dormController.GetDorms)
	router.POST(dormBasePath, dormController.CreateDorm)
	router.GET(fmt.Sprintf("%s/suggest/:letter", dormBasePath), dormController.GetDormSuggestions)
	router.GET(fmt.Sprintf("%s/facility", dormBasePath), dormController.GetAllDormFacilities)
	router.GET(fmt.Sprintf("%s/zone", dormBasePath), dormController.GetDormZones)
	router.PUT(fmt.Sprintf("%s/:id/picture", dormBasePath), dormController.UpdateDormPictures)
	router.GET(fmt.Sprintf("%s/:id", dormBasePath), dormController.GetDorm)
	router.PUT(fmt.Sprintf("%s/:id", dormBasePath), dormController.UpdateDorm)
	router.DELETE(fmt.Sprintf("%s/:id", dormBasePath), dormController.DeleteDorm)
}
