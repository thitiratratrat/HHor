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
		dormZones := roommateRequestController.dormService.GetDormZones()

		inputDorms := fl.Field().Interface().([]string)

		for _, inputDorm := range inputDorms {
			for indexZone, dormZone := range dormZones {
				if dormZone == inputDorm {
					break
				} else if indexZone == len(dormZones)-1 {
					return false
				}
			}
		}

		return true
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
// @Accept  multipart/form-data
// @Param data formData dto.RoommateRequestWithRegisteredDormDTO true "registered dorm request"
// @Param room_pictures formData file false "upload multiple room pictures,test this out in postman"
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
	if roommateRequestWithRegisteredDormDTO.RoomPictures != nil {

		files := context.Request.MultipartForm.File["room_pictures"]
		var roomPictureUrls []string

		for _, roomPicture := range files {
			picture, err := roomPicture.Open()

			if err != nil {
				panic(err)
			}

			roomPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.RoommateRequestRoomPictureFolder, roommateRequestWithRegisteredDormDTO.StudentID), roomPicture.Filename, context.Request)
			roomPictureUrls = append(roomPictureUrls, roomPictureUrl)
		}

		createdRoommateRequest = roommateRequestController.roommateRequestService.UpdateRoommateRequestWithRegisteredDormPictures(roommateRequestWithRegisteredDormDTO.StudentID, roomPictureUrls)
	}

	context.IndentedJSON(http.StatusCreated, createdRoommateRequest)
}

func (roommateRequestController *roommateRequestController) roomBelongsToDorm(roomID string, dormID string) bool {
	room := roommateRequestController.roomService.GetRoom(roomID)
	convertedDormID, _ := strconv.Atoi(dormID)

	return room.DormID == uint(convertedDormID)
}
