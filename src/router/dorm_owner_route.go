package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
	"github.com/thitiratratrat/hhor/src/middleware"
	"github.com/thitiratratrat/hhor/src/service"
)

const dormOwnerBasePath = "dorm-owner"

func SetDormOwnerRoutes(router *gin.Engine, dormOwnerController controller.DormOwnerController) {
	dormOwnerGroup := router.Group(fmt.Sprintf("/%s", dormOwnerBasePath)).Use(middleware.AuthorizeJWT(service.DormOwner))

	dormOwnerGroup.PUT("/:userid/picture", dormOwnerController.UploadPicture)
	dormOwnerGroup.GET("/:userid", dormOwnerController.GetDormOwner)
	dormOwnerGroup.PUT("/:userid", dormOwnerController.UpdateDormOwner)
	dormOwnerGroup.PUT("/:userid/bank-account", dormOwnerController.UpdateBankAccount)
	dormOwnerGroup.DELETE(":userid/bank-account", dormOwnerController.DeleteBankAccount)
}
