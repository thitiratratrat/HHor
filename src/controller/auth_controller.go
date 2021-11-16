package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type AuthController interface {
	RegisterStudent(context *gin.Context)
	LoginStudent(context *gin.Context)
}

func AuthControllerHandler(authService service.AuthService, studentService service.StudentService) AuthController {
	return &authController{
		authService:    authService,
		studentService: studentService,
	}
}

type authController struct {
	authService    service.AuthService
	studentService service.StudentService
}

// @Summary register student account
// @Description register student account
// @Tags auth
// @Accept  multipart/form-data
// @Produce json
// @Param data formData dto.RegisterStudentDTO true "student registration"
// @Param profile_picture formData file false "profile picture"
// @Success 200 {array} string "OK"
// @Failure 400,409  {object} dto.ErrorResponse
// @Router /auth/student/register [post]
func (authController *authController) RegisterStudent(context *gin.Context) {
	validate := validator.New()

	_ = validate.RegisterValidation("faculty", func(fl validator.FieldLevel) bool {
		faculties := authController.authService.GetFaculties()

		for _, faculty := range faculties {
			if faculty == fl.Field().String() {
				return true
			}
		}

		return false
	})

	var registerStudentDTO dto.RegisterStudentDTO

	bindErr := context.ShouldBind(&registerStudentDTO)

	if bindErr != nil {
		context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: errortype.ErrInvalidInput.Error(),
		})

		return
	}

	validateError := validate.Struct(registerStudentDTO)

	if validateError != nil {
		context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: validateError.Error(),
		})

		return
	}

	createError := authController.authService.RegisterStudent(registerStudentDTO)

	if createError != nil {
		context.IndentedJSON(http.StatusConflict, dto.ErrorResponse{
			Message: createError.Error(),
		})

		return
	}

	if registerStudentDTO.ProfilePicture != nil {
		filename := registerStudentDTO.StudentID + ".png"
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

		authController.studentService.UpdateStudent(registerStudentDTO.Email, map[string]interface{}{"picture_url": pictureUrl})
	}

	context.IndentedJSON(http.StatusOK, "")
}

// @Summary login
// @Description login
// @Tags auth
// @Produce json
// @Param data body dto.LoginCredentialsDTO true "login credentials"
// @Success 200 {array} string "OK"
// @Failure 400,404,401,500  {object} dto.ErrorResponse
// @Router /auth/student/login [post]
func (authController *authController) LoginStudent(context *gin.Context) {
	validate := validator.New()

	var loginCredentialsDTO dto.LoginCredentialsDTO

	bindErr := context.ShouldBind(&loginCredentialsDTO)

	if bindErr != nil {
		context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: bindErr.Error(),
		})

		return
	}

	validateError := validate.Struct(loginCredentialsDTO)

	if validateError != nil {
		context.IndentedJSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: validateError.Error(),
		})

		return
	}

	loginError := authController.authService.LoginStudent(loginCredentialsDTO)

	if errors.Is(loginError, errortype.ErrUnauthorized) {
		context.IndentedJSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: loginError.Error(),
		})

		return

	} else if errors.Is(loginError, errortype.ErrUserNotFound) {
		context.IndentedJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: loginError.Error(),
		})

		return
	} else if loginError != nil {
		context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: loginError.Error(),
		})

		return
	}

	context.IndentedJSON(http.StatusOK, "")
}
