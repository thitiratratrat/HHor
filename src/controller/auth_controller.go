package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
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
	VerifyCodeStudent(context *gin.Context)
	VerifyCodeDormOwner(context *gin.Context)
	ResendCodeStudent(context *gin.Context)
	ResendCodeDormOwner(context *gin.Context)
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
// @Accept  multipart/form-data
// @Param   profile_picture formData file false  "profile picture"
// @Param data formData dto.RegisterStudentDTO true "student registration"
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

	var pictureUrl *string
	if registerStudentDTO.ProfilePicture != nil {
		filename := registerStudentDTO.StudentID + ".png"
		file, _, err := context.Request.FormFile("profile_picture")

		if err != nil {
			panic(err)
		}

		url := utils.UploadPicture(file, constant.StudentProfilePictureFolder, filename)
		pictureUrl = &url
	}

	authController.authService.RegisterStudent(registerStudentDTO, pictureUrl)

	context.IndentedJSON(http.StatusCreated, "")
}

// @Summary register dorm owner account
// @Accept  multipart/form-data
// @Tags auth
// @Produce json
// @Param   profile_picture formData file false  "profile picture"
// @Param data formData dto.RegisterDormOwnerDTO true "dorm owner registration"
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

	var pictureUrl *string
	if registerDormOwnerDTO.ProfilePicture != nil {
		filename := registerDormOwnerDTO.Email + ".png"
		file, _, err := context.Request.FormFile("profile_picture")

		if err != nil {
			panic(err)
		}

		url := utils.UploadPicture(file, constant.StudentProfilePictureFolder, filename)
		pictureUrl = &url
	}

	authController.authService.RegisterDormOwner(registerDormOwnerDTO, pictureUrl)

	context.IndentedJSON(http.StatusCreated, "")
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
	dormOwnerID := fmt.Sprintf("%v", dormOwner.ID)
	token := authController.jwtService.GenerateToken(dormOwnerID, service.DormOwner)

	context.IndentedJSON(http.StatusOK, dto.LoginDTO{
		ID:    dormOwnerID,
		Token: token,
	})
}

// @Summary verify code student
// @Tags auth
// @Produce json
// @Param data body dto.VerifyCodeDTO true "verify code"
// @Success 200 "OK"
// @Failure 400,401,409,500  {object} dto.ErrorResponse
// @Router /auth/student/verify-code [post]
func (authController *authController) VerifyCodeStudent(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)
	validate := validator.New()

	var verifyCodeDTO dto.VerifyCodeDTO

	bindErr := context.ShouldBind(&verifyCodeDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(verifyCodeDTO)

	if validateError != nil {
		panic(validateError)
	}

	authController.authService.VerifyCodeStudent(verifyCodeDTO)

	context.IndentedJSON(http.StatusOK, "")
}

// @Summary verify code dorm owner
// @Tags auth
// @Produce json
// @Param data body dto.VerifyCodeDTO true "verify code"
// @Success 200 "OK"
// @Failure 400,401,409,500  {object} dto.ErrorResponse
// @Router /auth/dorm-owner/verify-code [post]
func (authController *authController) VerifyCodeDormOwner(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)
	validate := validator.New()

	var verifyCodeDTO dto.VerifyCodeDTO

	bindErr := context.ShouldBind(&verifyCodeDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(verifyCodeDTO)

	if validateError != nil {
		panic(validateError)
	}

	authController.authService.VerifyCodeDormOwner(verifyCodeDTO)

	context.IndentedJSON(http.StatusOK, "")
}

// @Summary resend code student
// @Tags auth
// @Produce json
// @Param data body dto.ResendCodeDTO true "resend code"
// @Success 200 "OK"
// @Failure 400,401,409,500  {object} dto.ErrorResponse
// @Router /auth/student/resend-code [post]
func (authController *authController) ResendCodeStudent(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)
	validate := validator.New()

	var resendCodeDTO dto.ResendCodeDTO

	bindErr := context.ShouldBind(&resendCodeDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(resendCodeDTO)

	if validateError != nil {
		panic(validateError)
	}

	authController.authService.ResendCodeStudent(resendCodeDTO)

	context.IndentedJSON(http.StatusOK, "")
}

// @Summary resend code dorm owner
// @Tags auth
// @Produce json
// @Param data body dto.ResendCodeDTO true "resend code"
// @Success 200 "OK"
// @Failure 400,401,409,500  {object} dto.ErrorResponse
// @Router /auth/dorm-owner/resend-code [post]
func (authController *authController) ResendCodeDormOwner(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)
	validate := validator.New()

	var resendCodeDTO dto.ResendCodeDTO

	bindErr := context.ShouldBind(&resendCodeDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(resendCodeDTO)

	if validateError != nil {
		panic(validateError)
	}

	authController.authService.ResendCodeDormOwner(resendCodeDTO)

	context.IndentedJSON(http.StatusOK, "")
}
