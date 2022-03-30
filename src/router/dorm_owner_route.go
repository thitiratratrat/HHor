package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const dormOwnerBasePath = "dorm-owner"

func SetDormOwnerRoutes(router *gin.Engine, dormOwnerController controller.DormOwnerController) {
	router.GET(fmt.Sprintf("%s/:id/dorm", dormOwnerBasePath), dormOwnerController.GetDorms)
}
