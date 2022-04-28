package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/fieldvalidator"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type RoomController interface {
	GetAllRoomFacilities(context *gin.Context)
	GetRoom(context *gin.Context)
	CreateRoom(context *gin.Context)
	UpdateRoom(context *gin.Context)
	UpdateRoomPictures(context *gin.Context)
	DeleteRoom(context *gin.Context)
}

func RoomControllerHandler(roomService service.RoomService, jwtService service.JWTService, fieldValidator fieldvalidator.FieldValidator) RoomController {
	return &roomController{
		roomService:    roomService,
		jwtService:     jwtService,
		fieldValidator: fieldValidator,
	}
}

type roomController struct {
	roomService    service.RoomService
	jwtService     service.JWTService
	fieldValidator fieldvalidator.FieldValidator
}

// @Summary get room facilities
// @Tags room
// @Produce json
// @Success 200 {array} string "OK"
// @Router /room/facility [get]
func (roomController *roomController) GetAllRoomFacilities(context *gin.Context) {
	facilities := roomController.roomService.GetAllRoomFacilities()

	context.IndentedJSON(http.StatusOK, facilities)
}

// @Summary get room by id
// @Tags room
// @Produce json
// @Success 200 {object} model.Room "OK"
// @Param id path int true "Room ID"
// @Router /room/{id} [get]
func (roomController *roomController) GetRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	id := context.Param("id")
	room := roomController.roomService.GetRoom(id)

	context.IndentedJSON(http.StatusOK, room)
}

// @Security BearerAuth
// @Summary create room
// @Tags room
// @Produce json
// @Param data body dto.RegisterRoomDTO true "register room"
// @Success 201 {object} model.Room "OK"
// @Router /room [post]
func (roomController *roomController) CreateRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	validate := validator.New()
	_ = validate.RegisterValidation("roomfacility", func(fl validator.FieldLevel) bool {
		return roomController.fieldValidator.ValidRoomFacility(fl.Field().Interface().([]string))
	})
	var registerRoomDTO dto.RegisterRoomDTO
	bindErr := context.ShouldBind(&registerRoomDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(registerRoomDTO)

	if validateError != nil {
		panic(validateError)
	}

	claims := roomController.jwtService.GetClaims(context.GetHeader("Authorization"))
	createdRoom := roomController.roomService.CreateRoom(claims["id"].(string), registerRoomDTO)

	context.IndentedJSON(http.StatusCreated, createdRoom)
}

// @Security BearerAuth
// @Summary update room
// @Tags room
// @Produce json
// @Param id path int true "Room ID"
// @Param data body dto.UpdateRoomDTO true "update room"
// @Success 201 {object} model.Room "OK"
// @Router /room/{id} [put]
func (roomController *roomController) UpdateRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	roomID := context.Param("id")
	validate := validator.New()
	_ = validate.RegisterValidation("roomfacility", func(fl validator.FieldLevel) bool {
		return roomController.fieldValidator.ValidRoomFacility(fl.Field().Interface().([]string))
	})
	var updateRoomDTO dto.UpdateRoomDTO
	bindErr := context.ShouldBind(&updateRoomDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(updateRoomDTO)

	if validateError != nil {
		panic(validateError)
	}

	claims := roomController.jwtService.GetClaims(context.GetHeader("Authorization"))
	createdRoom := roomController.roomService.UpdateRoom(roomID, claims["id"].(string), updateRoomDTO)

	context.IndentedJSON(http.StatusCreated, createdRoom)
}

// @Security BearerAuth
// @Summary update room pictures
// @Tags room
// @Produce json
// @Accept  multipart/form-data
// @Param id path int true "Room ID"
// @Param data formData dto.DormRoomPicturesDTO true "data"
// @Param pictures formData file false "upload multiple room pictures,test this out in postman"
// @Success 200 {object} model.RoommateRequestWithUnregisteredDorm "OK"
// @Router /room/{id}/picture [put]
func (roomController *roomController) UpdateRoomPictures(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	roomID := context.Param("id")
	validate := validator.New()
	var roomPicturesDTO dto.DormRoomPicturesDTO
	bindErr := context.ShouldBind(&roomPicturesDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(roomPicturesDTO)

	if validateError != nil {
		panic(validateError)
	}

	if roomPicturesDTO.Pictures == nil {
		context.IndentedJSON(http.StatusOK, "")

		return
	}

	claims := roomController.jwtService.GetClaims(context.GetHeader("Authorization"))
	if !roomController.roomService.CanUpdateRoom(roomID, claims["id"].(string)) {
		panic(errortype.ErrInvalidDormOwner)
	}

	files := context.Request.MultipartForm.File["pictures"]
	var roomPicturesUrl []string

	for _, roomPicture := range files {
		picture, err := roomPicture.Open()

		if err != nil {
			panic(err)
		}

		roomPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.RoomPictureFolder, roomID), roomPicture.Filename)
		roomPicturesUrl = append(roomPicturesUrl, roomPictureUrl)
	}

	updatedRoom := roomController.roomService.UpdateRoomPictures(roomID, roomPicturesUrl)

	context.IndentedJSON(http.StatusOK, updatedRoom)
}

// @Security BearerAuth
// @Summary delete room
// @Tags room
// @Produce json
// @Param id path int true "Room ID"
// @Success 201 {object} model.Room "OK"
// @Router /room/{id} [delete]
func (roomController *roomController) DeleteRoom(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	roomID := context.Param("id")

	claims := roomController.jwtService.GetClaims(context.GetHeader("Authorization"))
	roomController.roomService.DeleteRoom(roomID, claims["id"].(string))

	context.IndentedJSON(http.StatusOK, "")
}
