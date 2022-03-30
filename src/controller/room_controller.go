package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/fieldvalidator"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type RoomController interface {
	GetAllRoomFacilities(context *gin.Context)
	GetRoom(context *gin.Context)
	CreateRoom(context *gin.Context)
}

func RoomControllerHandler(roomService service.RoomService, fieldValidator fieldvalidator.FieldValidator) RoomController {
	return &roomController{
		roomService:    roomService,
		fieldValidator: fieldValidator,
	}
}

type roomController struct {
	roomService    service.RoomService
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

	createdRoom := roomController.roomService.CreateRoom(registerRoomDTO)

	context.IndentedJSON(http.StatusCreated, createdRoom)
}
