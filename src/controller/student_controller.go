package controller

import (
	"errors"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
	"gorm.io/gorm"
)

type StudentController interface {
	GetStudent(context *gin.Context)
	UpdateStudent(context *gin.Context)
}

func StudentControllerHandler(studentService service.StudentService) StudentController {
	return &studentController{
		studentService: studentService,
	}
}

type studentController struct {
	studentService service.StudentService
}

// @Summary update student detail
// @Description update student detail
// @Tags student
// @Produce json
// @Accept  multipart/form-data
// @Param   data formData dto.StudentUpdateDTO false  "student update"
// @Param   profile_picture formData file false  "profile picture"
// @Success 200 {array} string "OK"
// @Router /student [patch]
func (studentController *studentController) UpdateStudent(context *gin.Context) {
	validate := validator.New()

	var studentUpdateDTO dto.StudentUpdateDTO

	bindErr := context.ShouldBind(&studentUpdateDTO)

	if bindErr != nil {
		context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: bindErr.Error(),
		})

		return
	}

	validateError := validate.Struct(studentUpdateDTO)

	if validateError != nil {
		context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: validateError.Error(),
		})

		return
	}

	student, err := studentController.studentService.GetStudent(studentUpdateDTO.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.IndentedJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: errortype.ErrResourceNotFound.Error(),
		})

		return
	} else if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	studentUpdateMap := structs.Map(studentUpdateDTO)

	if studentUpdateDTO.ProfilePicture != nil {
		filename := student.StudentID + ".png"
		file, _, err := context.Request.FormFile("profile_picture")

		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		pictureUrl, err := utils.UploadPicture(file, constant.StudentProfilePictureFolder, filename, context.Request)

		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Message: err.Error(),
			})

			return
		}

		delete(studentUpdateMap, "ProfilePicture")

		studentUpdateMap["picture_url"] = pictureUrl
	}

	updateError := studentController.studentService.UpdateStudent(studentUpdateDTO.Email, studentUpdateMap)

	if updateError != nil {
		context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: updateError.Error(),
		})

		return
	}

	context.IndentedJSON(http.StatusOK, "")
}

// @Summary get student profile
// @Description returns student profile
// @Tags student
// @Produce json
// @Success 200 {object} model.Student "OK"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Param email path string true "Student Email"
// @Router /student/{email} [get]
func (studentController *studentController) GetStudent(context *gin.Context) {
	email := context.Param("email")

	student, err := studentController.studentService.GetStudent(email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.IndentedJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: errortype.ErrResourceNotFound.Error(),
		})

		return
	} else if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	context.IndentedJSON(http.StatusOK, student)
}
