package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const authBasePath = "auth"

func SetAuthRoutes(router *gin.Engine, authController controller.AuthController) {
	router.POST(fmt.Sprintf("%s/student/register", authBasePath), authController.RegisterStudent)
	router.POST(fmt.Sprintf("%s/student/login", authBasePath), authController.LoginStudent)
}
