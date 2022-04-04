package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/fieldvalidator"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type DormController interface {
	GetDorms(context *gin.Context)
	GetDorm(context *gin.Context)
	GetDormSuggestions(context *gin.Context)
	GetAllDormFacilities(context *gin.Context)
	GetDormZones(context *gin.Context)
	CreateDorm(context *gin.Context)
	UpdateDorm(context *gin.Context)
	UpdateDormPictures(context *gin.Context)
	DeleteDorm(context *gin.Context)
}

func DormControllerHandler(dormService service.DormService, fieldValidator fieldvalidator.FieldValidator) DormController {
	return &dormController{
		dormService:    dormService,
		fieldValidator: fieldValidator,
	}
}

type dormController struct {
	dormService    service.DormService
	fieldValidator fieldvalidator.FieldValidator
}

//TODO: filter by match score, cosine similarity, similarity measures
// @Summary get list of dorms
// @Tags dorm
// @Produce json
// @Accept json
// @Success 200 {array} dto.DormDTO "OK"
// @Param dorm_filter query dto.DormFilterDTO false "Dorm Filter"
// @Param type query []string false "type" collectionFormat(multi)
// @Param dorm_facilities query []string false "dorm facility" collectionFormat(multi)
// @Param room_facilities query []string false "room facility" collectionFormat(multi)
// @Router /dorm [get]
func (dormController *dormController) GetDorms(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	var dormFilterDTO dto.DormFilterDTO
	validate := validator.New()
	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return dormController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("roomfacilities", func(fl validator.FieldLevel) bool {
		return dormController.fieldValidator.ValidRoomFacility(fl.Field().Interface().([]string))
	})
	_ = validate.RegisterValidation("dormfacilities", func(fl validator.FieldLevel) bool {
		return dormController.fieldValidator.ValidDormFacility(fl.Field().Interface().([]string))
	})

	bindErr := context.ShouldBind(&dormFilterDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(dormFilterDTO)

	if validateError != nil {
		panic(validateError)
	}

	dormDTOs := dormController.dormService.GetDorms(dormFilterDTO)

	context.IndentedJSON(http.StatusOK, dormDTOs)
}

// @Summary get dorm details from dorm id
// @Tags dorm
// @Produce json
// @Success 200 {object} model.Dorm "OK"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Param id path int true "Dorm ID"
// @Router /dorm/{id} [get]
func (dormController *dormController) GetDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormID := context.Param("id")

	if _, err := strconv.Atoi(dormID); err != nil {
		context.IndentedJSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "id is not an integer"})

		return
	}

	dorm := dormController.dormService.GetDorm(dormID)

	context.IndentedJSON(http.StatusOK, dorm)
}

// @Summary get dorm names from first letter
// @Tags dorm
// @Produce json
// @Success 200 {array} dto.DormSuggestionDTO "OK"
// @Param letter path string true "First Letter"
// @Router /dorm/suggest/{letter} [get]
func (dormController *dormController) GetDormSuggestions(context *gin.Context) {
	letter := context.Param("letter")

	dormSuggestionDTOs := dormController.dormService.GetDormSuggestions(letter)

	context.IndentedJSON(http.StatusOK, dormSuggestionDTOs)
}

// @Summary get dorm facilities
// @Tags dorm
// @Produce json
// @Success 200 {array} string "OK"
// @Router /dorm/facility [get]
func (dormController *dormController) GetAllDormFacilities(context *gin.Context) {
	dormFacilities := dormController.dormService.GetAllDormFacilities()

	context.IndentedJSON(http.StatusOK, dormFacilities)
}

// @Summary get dorm zones
// @Tags dorm
// @Produce json
// @Success 200 {array} string "OK"
// @Router /dorm/zone [get]
func (dormController *dormController) GetDormZones(context *gin.Context) {
	dormZones := dormController.dormService.GetDormZones()

	context.IndentedJSON(http.StatusOK, dormZones)
}

// @Summary create dorm
// @Tags dorm
// @Produce json
// @Param data body dto.RegisterDormDTO true "register dorm"
// @Success 201 {object} model.Dorm "OK"
// @Router /dorm [post]
func (dormController *dormController) CreateDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	var registerDormDTO dto.RegisterDormDTO
	validate := validator.New()
	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return dormController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("dormfacilities", func(fl validator.FieldLevel) bool {
		return dormController.fieldValidator.ValidDormFacility(fl.Field().Interface().([]string))
	})
	bindErr := context.ShouldBind(&registerDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(registerDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdDorm := dormController.dormService.CreateDorm(registerDormDTO)

	context.IndentedJSON(http.StatusCreated, createdDorm)
}

// @Summary update dorm
// @Tags dorm
// @Produce json
// @Param id path int true "Dorm ID"
// @Param data body dto.UpdateDormDTO true "register dorm"
// @Success 201 {object} model.Dorm "OK"
// @Router /dorm/{id} [put]
func (dormController *dormController) UpdateDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormID := context.Param("id")
	if _, err := strconv.Atoi(dormID); err != nil {
		context.IndentedJSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "id is not an integer"})

		return
	}
	var updateDormDTO dto.UpdateDormDTO
	validate := validator.New()
	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return dormController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("dormfacilities", func(fl validator.FieldLevel) bool {
		return dormController.fieldValidator.ValidDormFacility(fl.Field().Interface().([]string))
	})
	bindErr := context.ShouldBind(&updateDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(updateDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	updatedDorm := dormController.dormService.UpdateDorm(dormID, updateDormDTO)

	context.IndentedJSON(http.StatusOK, updatedDorm)
}

// @Summary update dorm pictures
// @Tags dorm
// @Produce json
// @Accept  multipart/form-data
// @Param id path int true "Dorm ID"
// @Param data formData dto.DormRoomPicturesDTO true "data"
// @Param pictures formData file false "upload multiple room pictures,test this out in postman"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /dorm/{id}/picture [put]
func (dormController *dormController) UpdateDormPictures(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormID := context.Param("id")
	validate := validator.New()
	var dormPicturesDTO dto.DormRoomPicturesDTO
	bindErr := context.ShouldBind(&dormPicturesDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(dormPicturesDTO)

	if validateError != nil {
		panic(validateError)
	}

	if dormPicturesDTO.Pictures == nil {
		context.IndentedJSON(http.StatusOK, "")

		return
	}

	if !dormController.dormService.CanUpdateDorm(dormPicturesDTO.DormOwnerID, dormID) {
		panic(errortype.ErrInvalidDormOwner)
	}

	files := context.Request.MultipartForm.File["pictures"]
	var dormPicturesUrl []string

	for _, dormPicture := range files {
		picture, err := dormPicture.Open()

		if err != nil {
			panic(err)
		}

		dormPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.DormPictureFolder, dormID), dormPicture.Filename, context.Request)
		dormPicturesUrl = append(dormPicturesUrl, dormPictureUrl)
	}

	updatedDorm := dormController.dormService.UpdateDormPictures(dormID, dormPicturesUrl)

	context.IndentedJSON(http.StatusOK, updatedDorm)
}

// @Summary delete dorm
// @Tags dorm
// @Produce json
// @Param id path int true "Dorm ID"
// @Param dorm-owner-id query int true "Dorm Owner ID"
// @Success 200 "OK"
// @Router /dorm/{id} [delete]
func (dormController *dormController) DeleteDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	dormID := context.Param("id")
	dormOwnerID := context.Query("dorm-owner-id")

	dormController.dormService.DeleteDorm(dormID, dormOwnerID)

	context.IndentedJSON(http.StatusOK, "")
}
