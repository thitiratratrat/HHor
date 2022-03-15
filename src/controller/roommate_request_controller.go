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
	GetRoommateRequestsWithNoRoom(context *gin.Context)
	GetRoommateRequest(context *gin.Context)
	CreateRoommateRequestWithNoRoom(context *gin.Context)
	CreateRoommateRequestWithRegisteredDorm(context *gin.Context)
	CreateRoommateRequestWithUnregisteredDorm(context *gin.Context)
	UpdateRoommateRequestWithRegisteredDormPictures(context *gin.Context)
	UpdateRoommateRequestWithUnregisteredDormPictures(context *gin.Context)
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
func (roommateRequestController *roommateRequestController) GetRoommateRequestsWithNoRoom(context *gin.Context) {
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

	roommateRequests := roommateRequestController.roommateRequestService.GetRoommateRequestsWithNoRoom(roommateRequestFilterDTO)

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
// @Param data body dto.RoommateRequestWithNoRoomDTO true "no room request"
// @Success 200 {object} model.RoommateRequestWithNoRoom "OK"
// @Router /roommate-request/no-room [post]
func (roommateRequestController *roommateRequestController) CreateRoommateRequestWithNoRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone(fl.Field().Interface().([]string))
	})

	var createRoommateRequestWithNoRoomDTO dto.RoommateRequestWithNoRoomDTO
	bindErr := context.ShouldBind(&createRoommateRequestWithNoRoomDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(createRoommateRequestWithNoRoomDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestWithNoRoom(createRoommateRequestWithNoRoomDTO)

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary create roommate request with registered dorm
// @Tags roommate-request
// @Produce json
// @Accept  json
// @Param data body dto.RoommateRequestWithRegisteredDormDTO true "registered dorm request"
// @Success 200 {object} model.RoommateRequestWithRegisteredDorm "OK"
// @Router /roommate-request/registered-dorm [post]
func (roommateRequestController *roommateRequestController) CreateRoommateRequestWithRegisteredDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	var roommateRequestWithRegisteredDormDTO dto.RoommateRequestWithRegisteredDormDTO

	bindErr := context.ShouldBind(&roommateRequestWithRegisteredDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestWithRegisteredDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	if !roommateRequestController.roomBelongsToDorm(roommateRequestWithRegisteredDormDTO.RoomID, roommateRequestWithRegisteredDormDTO.DormID) {
		panic(errortype.ErrRoomDoesNotBelongToDorm)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDormDTO)

	//TODO: delete old files from bucket

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary create roommate request with unregistered dorm
// @Tags roommate-request
// @Produce json
// @Accept json
// @Param data body dto.RoommateRequestWithUnregisteredDormDTO true "unregistered dorm request"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /roommate-request/unregistered-dorm [post]
func (roommateRequestController *roommateRequestController) CreateRoommateRequestWithUnregisteredDorm(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidDormZone([]string{fl.Field().String()})
	})
	_ = validate.RegisterValidation("roomfacilities", func(fl validator.FieldLevel) bool {
		return roommateRequestController.fieldValidator.ValidRoomFacility(fl.Field().Interface().([]string))
	})

	var roommateRequestWithUnregisteredDormDTO dto.RoommateRequestWithUnregisteredDormDTO

	bindErr := context.ShouldBind(&roommateRequestWithUnregisteredDormDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestWithUnregisteredDormDTO)

	if validateError != nil {
		panic(validateError)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDormDTO)

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary update roommate request with reg dorm pictures
// @Tags roommate-request
// @Produce json
// @Accept  multipart/form-data
// @Param data formData dto.RoommateRequestPictureDTO true "data"
// @Param room_pictures formData file false "upload multiple room pictures,test this out in postman"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /roommate-request/registered-dorm/picture [put]
func (roommateRequestController *roommateRequestController) UpdateRoommateRequestWithRegisteredDormPictures(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	var roommateRequestPictureDTO dto.RoommateRequestPictureDTO

	bindErr := context.ShouldBind(&roommateRequestPictureDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestPictureDTO)

	if validateError != nil {
		panic(validateError)
	}

	if roommateRequestPictureDTO.RoomPictures == nil {
		context.IndentedJSON(http.StatusOK, "")

		return
	}

	if !roommateRequestController.roommateRequestService.CanUpdateRoommateRequestPicture(roommateRequestPictureDTO.StudentID, constant.RoommateRequestWithRegisteredDorm) {
		panic(errortype.ErrMismatchRoommateRequestType)
	}

	files := context.Request.MultipartForm.File["room_pictures"]
	var roomPictureUrls []string

	for _, roomPicture := range files {
		picture, err := roomPicture.Open()

		if err != nil {
			panic(err)
		}

		roomPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.RoommateRequestRoomPictureFolder, roommateRequestPictureDTO.StudentID), roomPicture.Filename, context.Request)
		roomPictureUrls = append(roomPictureUrls, roomPictureUrl)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.UpdateRoommateRequestWithRegisteredDormPictures(roommateRequestPictureDTO.StudentID, roomPictureUrls)

	context.IndentedJSON(http.StatusOK, createdRoommateRequest)
}

// @Summary update roommate request with unreg dorm pictures
// @Tags roommate-request
// @Produce json
// @Accept  multipart/form-data
// @Param data formData dto.RoommateRequestPictureDTO true "data"
// @Param room_pictures formData file false "upload multiple room pictures,test this out in postman"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /roommate-request/unregistered-dorm/picture [put]
func (roommateRequestController *roommateRequestController) UpdateRoommateRequestWithUnregisteredDormPictures(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	var roommateRequestPictureDTO dto.RoommateRequestPictureDTO

	bindErr := context.ShouldBind(&roommateRequestPictureDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roommateRequestPictureDTO)

	if validateError != nil {
		panic(validateError)
	}

	if roommateRequestPictureDTO.RoomPictures == nil {
		context.IndentedJSON(http.StatusOK, "")

		return
	}

	if !roommateRequestController.roommateRequestService.CanUpdateRoommateRequestPicture(roommateRequestPictureDTO.StudentID, constant.RoommateRequestWithUnregisteredDorm) {
		panic(errortype.ErrMismatchRoommateRequestType)
	}

	//TODO: delete old files from bucket
	files := context.Request.MultipartForm.File["room_pictures"]
	var roomPictureUrls []string

	for _, roomPicture := range files {
		picture, err := roomPicture.Open()

		if err != nil {
			panic(err)
		}

		roomPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.RoommateRequestRoomPictureFolder, roommateRequestPictureDTO.StudentID), roomPicture.Filename, context.Request)
		roomPictureUrls = append(roomPictureUrls, roomPictureUrl)
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.UpdateRoommateRequestWithUnregisteredDormPictures(roommateRequestPictureDTO.StudentID, roomPictureUrls)

	context.IndentedJSON(http.StatusOK, createdRoommateRequest)
}

func (roommateRequestController *roommateRequestController) roomBelongsToDorm(roomID string, dormID string) bool {
	room := roommateRequestController.roomService.GetRoom(roomID)
	convertedDormID, _ := strconv.Atoi(dormID)

	return room.DormID == uint(convertedDormID)
}
