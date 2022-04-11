package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/fieldvalidator"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type DormOwnerController interface {
	GetDormOwner(*gin.Context)
	UpdateDormOwner(*gin.Context)
	UploadPicture(context *gin.Context)
	UpdateBankAccount(context *gin.Context)
	DeleteBankAccount(context *gin.Context)
}

func DormOwnerControllerHandler(dormOwnerService service.DormOwnerService, fieldValidator fieldvalidator.FieldValidator) DormOwnerController {
	return &dormOwnerController{
		dormOwnerService: dormOwnerService,
		fieldValidator:   fieldValidator,
	}
}

type dormOwnerController struct {
	dormOwnerService service.DormOwnerService
	fieldValidator   fieldvalidator.FieldValidator
}

// @Summary get dorm owner profile
// @Tags dorm-owner
// @Produce json
// @Param id path int true "Dorm Owner ID"
// @Success 200 {object} model.DormOwner "OK"
// @Router /dorm-owner/{id} [get]
func (dormOwnerController *dormOwnerController) GetDormOwner(context *gin.Context) {
	dormOwnerID := context.Param("id")

	dormOwner := dormOwnerController.dormOwnerService.GetDormOwner(dormOwnerID)

	context.IndentedJSON(http.StatusOK, dormOwner)
}

// @Summary update dorm owner
// @Tags dorm-owner
// @Produce json
// @Param id path int true "Dorm Owner ID"
// @Param data body dto.UpdateDormOwnerDTO false "dorm owner update"
// @Success 200 {object} model.DormOwner "OK"
// @Router /dorm-owner/{id} [put]
func (dormOwnerController *dormOwnerController) UpdateDormOwner(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormOwnerID := context.Param("id")
	var updateDormOwnerDTO dto.UpdateDormOwnerDTO
	validate := validator.New()
	_ = validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return dormOwnerController.fieldValidator.ValidPhoneNumber(fl.Field().String())
	})
	bindErr := context.ShouldBind(&updateDormOwnerDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(updateDormOwnerDTO)

	if validateError != nil {
		panic(validateError)
	}

	dormOwner := dormOwnerController.dormOwnerService.UpdateDormOwner(dormOwnerID, updateDormOwnerDTO)

	context.IndentedJSON(http.StatusOK, dormOwner)
}

// @Summary update bank account
// @Tags dorm-owner
// @Produce json
// @Param id path int true "Dorm Owner ID"
// @Param data body dto.UpdateBankAccountDTO false "bank account update"
// @Success 200 {object} model.DormOwner "OK"
// @Router /dorm-owner/{id}/bank-account [put]
func (dormOwnerController *dormOwnerController) UpdateBankAccount(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormOwnerID := context.Param("id")
	var updateBankAccountDTO dto.UpdateBankAccountDTO
	validate := validator.New()
	bindErr := context.ShouldBind(&updateBankAccountDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(updateBankAccountDTO)

	if validateError != nil {
		panic(validateError)
	}

	dormOwner := dormOwnerController.dormOwnerService.UpdateBankAccount(dormOwnerID, updateBankAccountDTO)

	context.IndentedJSON(http.StatusOK, dormOwner)
}

// @Summary delete bank account
// @Tags dorm-owner
// @Produce json
// @Param id path int true "Dorm Owner ID"
// @Success 200 {object} model.DormOwner "OK"
// @Router /dorm-owner/{id}/bank-account [delete]
func (dormOwnerController *dormOwnerController) DeleteBankAccount(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormOwnerID := context.Param("id")
	dormOwner := dormOwnerController.dormOwnerService.DeleteBankAccount(dormOwnerID)

	context.IndentedJSON(http.StatusOK, dormOwner)
}

// @Summary upload picture
// @Tags dorm-owner
// @Accept  multipart/form-data
// @Produce json
// @Success 200 {object} model.DormOwner "OK"
// @Param id path string true "Dorm Owner ID"
// @Param profile_picture formData file false  "profile picture"
// @Param bank_qr formData file false  "bank qr"
// @Param data formData dto.DormOwnerPictureDTO true  "profile picture"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /dorm-owner/{id}/picture [put]
func (dormOwnerController *dormOwnerController) UploadPicture(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormOwnerID := context.Param("id")
	var dormOwnerPictureDTO dto.DormOwnerPictureDTO
	validate := validator.New()
	bindErr := context.ShouldBind(&dormOwnerPictureDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(dormOwnerPictureDTO)

	if validateError != nil {
		panic(validateError)
	}

	var profilePictureUrl string
	var bankQrUrl string

	if dormOwnerPictureDTO.ProfilePicture != nil {
		filename := dormOwnerID + ".png"
		file, _, err := context.Request.FormFile("profile_picture")

		if err != nil {
			panic(err)
		}

		profilePicture := utils.UploadPicture(file, constant.DormOwnerProfilePictureFolder, filename, context.Request)
		profilePictureUrl = profilePicture
	}

	//TODO: delete old files from bucket storage too
	if dormOwnerPictureDTO.BankQR != nil {
		filename := dormOwnerID + ".png"
		file, _, err := context.Request.FormFile("bank_qr")

		if err != nil {
			panic(err)
		}

		bankQr := utils.UploadPicture(file, constant.BankQRPictureFolder, filename, context.Request)
		bankQrUrl = bankQr
	}

	updateDormOwner := dormOwnerController.dormOwnerService.UpdateDormOwnerPictures(dormOwnerID, profilePictureUrl, bankQrUrl)

	context.IndentedJSON(http.StatusOK, updateDormOwner)
}
