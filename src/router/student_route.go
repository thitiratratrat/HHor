package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const studentBasePath = "student"

func SetStudentRoutes(router *gin.Engine, studentController controller.StudentController) {
	router.GET(fmt.Sprintf("%s/:email", studentBasePath), studentController.GetStudent)
	router.PATCH(fmt.Sprintf("%s", studentBasePath), studentController.UpdateStudent)
}
