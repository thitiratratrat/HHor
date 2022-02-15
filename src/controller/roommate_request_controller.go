package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

const (
	RoommateRequestWithNoRoom         string = "NO_ROOM"
	RoommateRequestWithRegisteredDorm string = "REGISTERED_DORM"
)

type RoommateRequestController interface {
	CreateRoomRequestWithNoRoom(context *gin.Context)
	CreateRoommateRequestWithRegisteredDorm(context *gin.Context)
	CreateRoommateRequestWithUnregisteredDorm(context *gin.Context)
	UpdateRoommateRequestWithRegisteredDormPictures(context *gin.Context)
	UpdateRoommateRequestWithUnregisteredDormPictures(context *gin.Context)
}

func RoommateRequestControllerHandler(roommateRequestService service.RoommateRequestService, dormService service.DormService, roomService service.RoomService) RoommateRequestController {
	return &roommateRequestController{
		roommateRequestService: roommateRequestService,
		dormService:            dormService,
		roomService:            roomService,
	}
}

type roommateRequestController struct {
	roommateRequestService service.RoommateRequestService
	dormService            service.DormService
	roomService            service.RoomService
}

// @Summary create roommate request with no room
// @Description create roommate request with no room
// @Tags roommate-request
// @Produce json
// @Param data body dto.RoommateRequestWithNoRoomDTO true "no room request"
// @Success 200 {object} model.RoommateRequestWithNoRoom "OK"
// @Router /roommate-request/no-room [post]
func (roommateRequestController *roommateRequestController) CreateRoomRequestWithNoRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	_ = validate.RegisterValidation("dormzone", func(fl validator.FieldLevel) bool {
		return roommateRequestController.validDormZone(fl.Field().Interface().([]string))
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
// @Description create roommate request with registered dorm
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
		panic(errors.New("room does not belong to dorm"))
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDormDTO)

	//TODO: delete old files from bucket

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary create roommate request with unregistered dorm
// @Description create roommate request with unregistered dorm
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
		return roommateRequestController.validDormZone([]string{fl.Field().String()})
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

	if !roommateRequestController.validRoomFacility(roommateRequestWithUnregisteredDormDTO.RoomFacilities) {
		panic(errors.New("invalid room facility"))
	}

	createdRoommateRequest := roommateRequestController.roommateRequestService.CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDormDTO)

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

// @Summary update roommate request with reg dorm pictures
// @Description update roommate request with reg dorm pictures
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

	if !roommateRequestController.roommateRequestService.CanUpdateRoommateRequestPicture(roommateRequestPictureDTO.StudentID, service.RoommateRequestWithRegisteredDorm) {
		panic(errors.New("student does not have this type of roommate request opened"))
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
// @Description update roommate request with reg dorm pictures
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

	if !roommateRequestController.roommateRequestService.CanUpdateRoommateRequestPicture(roommateRequestPictureDTO.StudentID, service.RoommateRequestWithUnregisteredDorm) {
		panic(errors.New("student does not have this type of roommate request opened"))
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

func (roommateRequestController *roommateRequestController) validDormZone(inputDormZones []string) bool {
	dormZones := roommateRequestController.dormService.GetDormZones()

	for _, inputDormZone := range inputDormZones {
		for index, dormZone := range dormZones {
			if dormZone == inputDormZone {
				break
			} else if index == len(dormZones)-1 {
				return false
			}
		}
	}

	return true
}

func (roommateRequestController *roommateRequestController) validRoomFacility(inputRoomFacilities []string) bool {
	roomFacilities := roommateRequestController.roomService.GetAllRoomFacilities()

	for _, inputRoomFacility := range inputRoomFacilities {
		for index, roomFacility := range roomFacilities {
			if roomFacility == inputRoomFacility {
				break
			} else if index == len(roomFacilities)-1 {
				return false
			}
		}
	}

	return true
}
