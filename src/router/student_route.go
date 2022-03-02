package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
)

const studentBasePath = "student"

func SetStudentRoutes(router *gin.Engine, studentController controller.StudentController) {
	router.GET(fmt.Sprintf("%s/faculty", studentBasePath), studentController.GetFaculties)
	router.GET(fmt.Sprintf("%s/habit", studentBasePath), studentController.GetHabits)
	router.GET(fmt.Sprintf("%s/:id", studentBasePath), studentController.GetStudent)
	router.PATCH(fmt.Sprintf("%s/:id", studentBasePath), studentController.UpdateStudent)
	router.POST(fmt.Sprintf("%s/picture/:id", studentBasePath), studentController.UploadPicture)
}
