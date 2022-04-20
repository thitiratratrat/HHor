package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thitiratratrat/hhor/src/controller"
	"github.com/thitiratratrat/hhor/src/middleware"
	"github.com/thitiratratrat/hhor/src/service"
)

const studentBasePath = "student"

func SetStudentRoutes(router *gin.Engine, studentController controller.StudentController) {
	studentGroup := router.Group(fmt.Sprintf("/%s", studentBasePath)).Use(middleware.AuthorizeJWT(service.Student))

	router.GET(fmt.Sprintf("/%s/faculty", studentBasePath), studentController.GetFaculties)
	router.GET(fmt.Sprintf("/%s/habit", studentBasePath), studentController.GetHabits)

	studentGroup.GET("/:userid", studentController.GetStudent)
	studentGroup.PUT("/:userid/picture", studentController.UploadPicture)
	studentGroup.PATCH("/:userid", studentController.UpdateStudent)
	studentGroup.PATCH("/:userid/habit", studentController.UpdateHabit)
	studentGroup.PATCH("/:userid/preference", studentController.UpdatePreference)
}
