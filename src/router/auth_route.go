package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const authBasePath = "auth"

func SetAuthRoutes(router *gin.Engine, authController controller.AuthController) {
	router.POST(fmt.Sprintf("%s/student/register", authBasePath), authController.RegisterStudent)
	router.POST(fmt.Sprintf("%s/dorm-owner/register", authBasePath), authController.RegisterDormOwner)
	router.POST(fmt.Sprintf("%s/student/login", authBasePath), authController.LoginStudent)
	router.POST(fmt.Sprintf("%s/dorm-owner/login", authBasePath), authController.LoginDormOwner)
	router.POST(fmt.Sprintf("%s/student/verify-code", authBasePath), authController.VerifyCodeStudent)
	router.POST(fmt.Sprintf("%s/dorm-owner/verify-code", authBasePath), authController.VerifyCodeDormOwner)
	router.POST(fmt.Sprintf("%s/student/resend-code", authBasePath), authController.ResendCodeStudent)
	router.POST(fmt.Sprintf("%s/dorm-owner/resend-code", authBasePath), authController.ResendCodeDormOwner)
}
