package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const dormBasePath = "dorm"

func SetDormRoutes(router *gin.Engine, dormController controller.DormController) {
	router.GET(dormBasePath, dormController.GetDorms)
	router.GET(fmt.Sprintf("%s/name/:letter", dormBasePath), dormController.GetDormNames)
	router.GET(fmt.Sprintf("%s/facility", dormBasePath), dormController.GetAllDormFacilities)
	router.GET(fmt.Sprintf("%s/zone", dormBasePath), dormController.GetDormZones)
	router.GET(fmt.Sprintf("%s/:id", dormBasePath), dormController.GetDormDetail)
}
