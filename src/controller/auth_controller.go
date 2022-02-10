package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type AuthController interface {
	RegisterStudent(context *gin.Context)
	LoginStudent(context *gin.Context)
}

func AuthControllerHandler(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

type authController struct {
	authService service.AuthService
}

// @Summary register student account
// @Description register student account
// @Tags auth
// @Produce json
// @Param data body dto.RegisterStudentDTO true "student registration"
// @Success 200 {object} model.Student "OK"
// @Failure 400,409  {object} dto.ErrorResponse
// @Router /auth/student/register [post]
func (authController *authController) RegisterStudent(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

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
		panic(bindErr)
	}

	validateError := validate.Struct(registerStudentDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdStudent, createError := authController.authService.RegisterStudent(registerStudentDTO)

	if createError != nil {
		panic(createError)
	}

	context.IndentedJSON(http.StatusOK, createdStudent)
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
	defer utils.RecoverInvalidInput(context)
	validate := validator.New()

	var loginCredentialsDTO dto.LoginCredentialsDTO

	bindErr := context.ShouldBind(&loginCredentialsDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(loginCredentialsDTO)

	if validateError != nil {
		panic(validateError)
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
