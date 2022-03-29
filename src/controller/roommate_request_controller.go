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

type RoommateRequestController interface {
	GetRoommateRequestsWithRoom(context *gin.Context)
	GetRoommateRequestsNoRoom(context *gin.Context)
	GetRoommateRequest(context *gin.Context)
	CreateRoommateRequestNoRoom(context *gin.Context)
	CreateRoommateRequestRegDorm(context *gin.Context)
	CreateRoommateRequestUnregDorm(context *gin.Context)
	UpdateRoommateRequestRegDormPictures(context *gin.Context)
	UpdateRoommateRequestUnregDormPictures(context *gin.Context)
	UpdateRoommateRequestRegDorm(context *gin.Context)
	UpdateRoommateRequestUnregDorm(context *gin.Context)
	UpdateRoommateRequestNoRoom(context *gin.Context)
	DeleteRoommateRequest(context *gin.Context)
}

func RoommateRequestControllerHandler(roommateRequestService service.RoommateRequestService, dormService service.DormService, roomService service.RoomService, fieldValidator fieldvalidator.FieldValidator) RoommateRequestController {
	return &roommateRequestController{
		roommateRequestService: roommateRequestService,
		dormService:            dormService,
		roomService:            roomService,
		fieldValidator:         fieldValidator,
	}
}

type roommateRequestController struct {
	roommateRequestService service.RoommateRequestService
	dormService            service.DormService
	roomService            service.RoomService
	fieldValidator         fieldvalidator.FieldValidator
}

// @Summary get roommate requests with room
// @Tags roommate-request
// @Produce json
// @Param data query dto.RoommateRequestRoomFilterDTO true "room request filter"
// @Param gender query []string false "gender" collectionFormat(multi)
// @Param faculties query []string false "faculty" collectionFormat(multi)
// @Param year_of_study query []string false "year of study" collectionFormat(multi)
// @Param number_of_roommates query []string false "number of roommates" collectionFormat(multi)
// @Param room_facilities query []string false "room facilities" collectionFormat(multi)
// @Success 200 {array} dto.RoommateRequestWithRoomDTO "OK"
// @Router /roommate-request/room [get]
func (roommateRequestController *roommateRequestController) GetRoommateRequestsWithRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("faculty", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidFaculty(fl.Field().Interface().([]string))
	})
	_ = validate.RegisterValidation("roomfacility", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidRoomFacility(fl.Field().Interface().([]string))
	})

	var roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO
	bindErr := context.ShouldBind(&roommateRequestRoomFilterDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestRoomFilterDTO)

	if validateError != nil {
		panic(validateError)
	}

	roommateRequests := roommateRequestController.roommateRequestService.GetRoommateRequestsWithRoom(roommateRequestRoomFilterDTO)

	context.IndentedJSON(http.StatusOK, roommateRequests)
}

// @Summary get roommate requests with no room
// @Tags roommate-request
// @Produce json
// @Param data query dto.RoommateRequestFilterDTO true "room request filter"
// @Param gender query []string false "gender" collectionFormat(multi)
// @Param faculties query []string false "faculty" collectionFormat(multi)
// @Param year_of_study query []string false "year of study" collectionFormat(multi)
// @Success 200 {array} model.RoommateRequestWithNoRoom "OK"
// @Router /roommate-request/no-room [get]
func (roommateRequestController *roommateRequestController) GetRoommateRequestsNoRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("faculty", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidFaculty(fl.Field().Interface().([]string))
	})

	var roommateRequestFilterDTO dto.RoommateRequestFilterDTO
	bindErr := context.ShouldBind(&roommateRequestFilterDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestFilterDTO)

	if validateError != nil {
		panic(validateError)
	}

	roommateRequests := roommateRequestController.roommateRequestService.GetRoommateRequestsNoRoom(roommateRequestFilterDTO)

	context.IndentedJSON(http.StatusOK, roommateRequests)
}

// @Summary get roommate request
// @Tags roommate-request
// @Produce json
// @Success 200 {array} dto.RoommateRequestDTO "OK"
// @Param id path int true "Roommate request ID"
// @Router /roommate-request/{id} [get]
func (roommateRequestController *roommateRequestController) GetRoommateRequest(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	roommateRequestID := context.Param("id")

	if _, err := strconv.Atoi(roommateRequestID); err != nil {
		context.IndentedJSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "id is not an integer"})

		return
	}

	roommateRequests := roommateRequestController.roommateRequestService.GetRoommateRequest(roommateRequestID)

	context.IndentedJSON(http.StatusOK, roommateRequests)
}

// @Summary create roommate request with no room
// @Tags roommate-request
// @Produce json
// @Param id path int true "Student ID"
// @Param data body dto.RoommateRequestNoRoomDTO true "no room request"
// @Success 200 {object} model.RoommateRequestWithNoRoom "OK"
// @Router /roommate-request/no-room/{id} [post]
func (roommateRequestController *roommateRequestController) CreateRoommateRequestNoRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()
	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone(fl.Field().Interface().([]string))
	})

	var createRoommateRequestNoRoomDTO dto.RoommateRequestNoRoomDTO
	bindErr := context.ShouldBind(&createRoommateRequestNoRoomDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(createRoommateRequestNoRoomDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestNoRoom(studentId, createRoommateRequestNoRoomDTO)

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary create roommate request with reg dorm
// @Tags roommate-request
// @Produce json
// @Accept  json
// @Param id path int true "Student ID"
// @Param data body dto.RoommateRequestRegDormDTO true "reg dorm request"
// @Success 201 {object} model.RoommateRequestWithRegisteredDorm "OK"
// @Router /roommate-request/registered-dorm/{id} [post]
func (roommateRequestController *roommateRequestController) CreateRoommateRequestRegDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()
	var roommateRequestRegDormDTO dto.RoommateRequestRegDormDTO
	bindErr := context.ShouldBind(&roommateRequestRegDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestRegDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	if !roommateRequestController.roomBelongsToDorm(roommateRequestRegDormDTO.RoomID, roommateRequestRegDormDTO.DormID) {
		panic(errortype.ErrRoomDoesNotBelongToDorm)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestRegDorm(studentId, roommateRequestRegDormDTO)

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary create roommate request with unreg dorm
// @Tags roommate-request
// @Produce json
// @Accept json
// @Param id path int true "Student ID"
// @Param data body dto.RoommateRequestUnregDormDTO true "unreg dorm request"
// @Success 201 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /roommate-request/unregistered-dorm/{id} [post]
func (roommateRequestController *roommateRequestController) CreateRoommateRequestUnregDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()
	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("roomfacilities", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidRoomFacility(fl.Field().Interface().([]string))
	})
	var roommateRequestUnregDormDTO dto.RoommateRequestUnregDormDTO
	bindErr := context.ShouldBind(&roommateRequestUnregDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestUnregDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestUnregDorm(studentId, roommateRequestUnregDormDTO)

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary update roommate request with reg dorm pictures
// @Tags roommate-request
// @Produce json
// @Accept  multipart/form-data
// @Param id path int true "Student ID"
// @Param data formData dto.PicturesDTO true "data"
// @Param pictures formData file false "upload multiple room pictures,test this out in postman"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /roommate-request/registered-dorm/{id}/picture [put]
func (roommateRequestController *roommateRequestController) UpdateRoommateRequestRegDormPictures(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()

	var roommateRequestPictureDTO dto.PicturesDTO

	bindErr := context.ShouldBind(&roommateRequestPictureDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestPictureDTO)

	if validateError != nil {
		panic(validateError)
	}

	if roommateRequestPictureDTO.Pictures == nil {
		context.IndentedJSON(http.StatusOK, "")

		return
	}

	if !roommateRequestController.roommateRequestService.CanUpdateRoommateRequest(studentId, constant.RoommateRequestRegDorm) {
		panic(errortype.ErrMismatchRoommateRequestType)
	}

	files := context.Request.MultipartForm.File["pictures"]
	var roomPictureUrls []string

	for _, roomPicture := range files {
		picture, err := roomPicture.Open()

		if err != nil {
			panic(err)
		}

		roomPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.RoommateRequestRoomPictureFolder, studentId), roomPicture.Filename, context.Request)
		roomPictureUrls = append(roomPictureUrls, roomPictureUrl)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.UpdateRoommateRequestRegDormPictures(studentId, roomPictureUrls)

	context.IndentedJSON(http.StatusOK, createdRoommateRequest)
}

// @Summary update roommate request with unreg dorm pictures
// @Tags roommate-request
// @Produce json
// @Accept  multipart/form-data
// @Param id path int true "Student ID"
// @Param data formData dto.PicturesDTO true "data"
// @Param pictures formData file false "upload multiple room pictures,test this out in postman"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /roommate-request/unregistered-dorm/{id}/picture [put]
func (roommateRequestController *roommateRequestController) UpdateRoommateRequestUnregDormPictures(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()
	var roommateRequestPictureDTO dto.PicturesDTO
	bindErr := context.ShouldBind(&roommateRequestPictureDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestPictureDTO)

	if validateError != nil {
		panic(validateError)
	}

	if roommateRequestPictureDTO.Pictures == nil {
		context.IndentedJSON(http.StatusOK, "")

		return
	}

	if !roommateRequestController.roommateRequestService.CanUpdateRoommateRequest(studentId, constant.RoommateRequestUnregDorm) {
		panic(errortype.ErrMismatchRoommateRequestType)
	}

	//TODO: delete old files from bucket
	files := context.Request.MultipartForm.File["pictures"]
	var roomPictureUrls []string

	for _, roomPicture := range files {
		picture, err := roomPicture.Open()

		if err != nil {
			panic(err)
		}

		roomPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.RoommateRequestRoomPictureFolder, studentId), roomPicture.Filename, context.Request)
		roomPictureUrls = append(roomPictureUrls, roomPictureUrl)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.UpdateRoommateRequestUnregDormPictures(studentId, roomPictureUrls)

	context.IndentedJSON(http.StatusOK, createdRoommateRequest)
}

// @Summary update roommate request reg dorm
// @Tags roommate-request
// @Produce json
// @Accept  json
// @Param id path int true "Student ID"
// @Param data body dto.RoommateRequestRegDormDTO true "reg dorm update"
// @Success 200 {object} model.RoommateRequestWithRegisteredDorm "OK"
// @Router /roommate-request/registered-dorm/{id} [put]
func (roommateRequestController *roommateRequestController) UpdateRoommateRequestRegDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()
	var roommateRequestRegDormDTO dto.RoommateRequestRegDormDTO
	bindErr := context.ShouldBind(&roommateRequestRegDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestRegDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	if !roommateRequestController.roomBelongsToDorm(roommateRequestRegDormDTO.RoomID, roommateRequestRegDormDTO.DormID) {
		panic(errortype.ErrRoomDoesNotBelongToDorm)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.UpdateRoommateRequestRegDorm(studentId, roommateRequestRegDormDTO)

	context.IndentedJSON(http.StatusOK, createdRoommateRequest)
}

// @Summary update roommate request with unreg dorm
// @Tags roommate-request
// @Produce json
// @Accept json
// @Param id path string true "Student ID"
// @Param data body dto.RoommateRequestUnregDormDTO true "unreg dorm request"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /roommate-request/unregistered-dorm/{id} [put]
func (roommateRequestController *roommateRequestController) UpdateRoommateRequestUnregDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()
	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("roomfacilities", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidRoomFacility(fl.Field().Interface().([]string))
	})
	var roommateRequestUnregDormDTO dto.RoommateRequestUnregDormDTO
	bindErr := context.ShouldBind(&roommateRequestUnregDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestUnregDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.UpdateRoommateRequestUnregDorm(studentId, roommateRequestUnregDormDTO)

	context.IndentedJSON(http.StatusOK, createdRoommateRequest)
}

// @Summary update roommate request with no room
// @Tags roommate-request
// @Produce json
// @Param id path int true "Student ID"
// @Param data body dto.RoommateRequestNoRoomDTO true "no room request"
// @Success 200 {object} model.RoommateRequestWithNoRoom "OK"
// @Router /roommate-request/no-room/{id} [put]
func (roommateRequestController *roommateRequestController) UpdateRoommateRequestNoRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	studentId := context.Param("id")
	validate := validator.New()
	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone(fl.Field().Interface().([]string))
	})
	var updatedRoommateRequestNoRoomDTO dto.RoommateRequestNoRoomDTO
	bindErr := context.ShouldBind(&updatedRoommateRequestNoRoomDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(updatedRoommateRequestNoRoomDTO)

	if validateError != nil {
		panic(validateError)
	}

	updatedRoommateRequest := roommateRequestController.roommateRequestService.UpdateRoommateRequestNoRoom(studentId, updatedRoommateRequestNoRoomDTO)

	context.IndentedJSON(http.StatusOK, updatedRoommateRequest)
}

// @Summary delete roommate request
// @Tags roommate-request
// @Produce json
// @Success 200 "OK"
// @Param id path int true "Roommate request ID"
// @Router /roommate-request/{id} [delete]
func (roommateRequestController *roommateRequestController) DeleteRoommateRequest(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	roommateRequestID := context.Param("id")

	if _, err := strconv.Atoi(roommateRequestID); err != nil {
		context.IndentedJSON(http.StatusBadRequest, &dto.ErrorResponse{Message: "id is not an integer"})

		return
	}

	roommateRequestController.roommateRequestService.DeleteRoommateRequest(roommateRequestID)

	context.IndentedJSON(http.StatusOK, "")
}

func (roommateRequestController *roommateRequestController) roomBelongsToDorm(roomID string, dormID string) bool {
	room := roommateRequestController.roomService.GetRoom(roomID)
	convertedDormID, _ := strconv.Atoi(dormID)

	return room.DormID == uint(convertedDormID)
}
