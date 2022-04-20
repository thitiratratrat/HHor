package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/fieldvalidator"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type AuthController interface {
	RegisterStudent(context *gin.Context)
	LoginStudent(context *gin.Context)
	RegisterDormOwner(context *gin.Context)
	LoginDormOwner(context *gin.Context)
}

func AuthControllerHandler(authService service.AuthService, jwtService service.JWTService, fieldValidator fieldvalidator.FieldValidator) AuthController {
	return &authController{
		authService:    authService,
		jwtService:     jwtService,
		fieldValidator: fieldValidator,
	}
}

type authController struct {
	authService    service.AuthService
	jwtService     service.JWTService
	fieldValidator fieldvalidator.FieldValidator
}

// @Summary register student account
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
		return authController.fieldValidator.ValidFaculty([]string{fl.Field().String()})
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

	createdStudent := authController.authService.RegisterStudent(registerStudentDTO)

	context.IndentedJSON(http.StatusOK, createdStudent)
}

// @Summary register dorm owner account
// @Tags auth
// @Produce json
// @Param data body dto.RegisterDormOwnerDTO true "dorm owner registration"
// @Success 200 {object} model.DormOwner "OK"
// @Failure 400,409  {object} dto.ErrorResponse
// @Router /auth/dorm-owner/register [post]
func (authController *authController) RegisterDormOwner(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()
	_ = validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return authController.fieldValidator.ValidPhoneNumber(fl.Field().String())
	})
	var registerDormOwnerDTO dto.RegisterDormOwnerDTO
	bindErr := context.ShouldBind(&registerDormOwnerDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(registerDormOwnerDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdDormOwner := authController.authService.RegisterDormOwner(registerDormOwnerDTO)

	context.IndentedJSON(http.StatusOK, createdDormOwner)
}

// @Summary login student
// @Tags auth
// @Produce json
// @Param data body dto.LoginCredentialsDTO true "login credentials"
// @Success 200 {object} dto.LoginDTO "OK"
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

	student := authController.authService.LoginStudent(loginCredentialsDTO)
	token := authController.jwtService.GenerateToken(student.ID, service.Student)

	context.IndentedJSON(http.StatusOK, dto.LoginDTO{
		ID:    student.ID,
		Token: token,
	})
}

// @Summary login dorm owner
// @Tags auth
// @Produce json
// @Param data body dto.LoginCredentialsDTO true "login credentials"
// @Success 200 {object} dto.LoginDTO "OK"
// @Failure 400,404,401,500  {object} dto.ErrorResponse
// @Router /auth/dorm-owner/login [post]
func (authController *authController) LoginDormOwner(context *gin.Context) {
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

	dormOwner := authController.authService.LoginDormOwner(loginCredentialsDTO)
	dormOwnerID := strconv.FormatUint(uint64(dormOwner.ID), 10)
	token := authController.jwtService.GenerateToken(dormOwnerID, service.DormOwner)

	context.IndentedJSON(http.StatusOK, dto.LoginDTO{
		ID:    dormOwnerID,
		Token: token,
	})
}
