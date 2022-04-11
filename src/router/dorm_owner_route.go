package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const dormOwnerBasePath = "dorm-owner"

func SetDormOwnerRoutes(router *gin.Engine, dormOwnerController controller.DormOwnerController) {
	router.GET(fmt.Sprintf("%s/:id", dormOwnerBasePath), dormOwnerController.GetDormOwner)
	router.PUT(fmt.Sprintf("%s/:id", dormOwnerBasePath), dormOwnerController.UpdateDormOwner)
	router.PUT(fmt.Sprintf("%s/:id/picture", dormOwnerBasePath), dormOwnerController.UploadPicture)
	router.PUT(fmt.Sprintf("%s/:id/bank-account", dormOwnerBasePath), dormOwnerController.UpdateBankAccount)
	router.DELETE(fmt.Sprintf("%s/:id/bank-account", dormOwnerBasePath), dormOwnerController.DeleteBankAccount)
}
