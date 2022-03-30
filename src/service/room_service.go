package service

import (
	"strconv"
	"time"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type RoomService interface {
	GetAllRoomFacilities() []string
	GetRoom(id string) model.Room
	CreateRoom(registerRoomDTO dto.RegisterRoomDTO) model.Room
}

func RoomServiceHandler(dormRepository repository.DormRepository, roomRepository repository.RoomRepository) RoomService {
	return &roomService{
		roomRepository: roomRepository,
		dormRepository: dormRepository,
	}
}

type roomService struct {
	dormRepository repository.DormRepository
	roomRepository repository.RoomRepository
}

func (roomService *roomService) GetAllRoomFacilities() []string {
	roomFacilities := roomService.roomRepository.FindAllRoomFacilities()

	return roomFacilities
}

func (roomService *roomService) GetRoom(id string) model.Room {
	room, err := roomService.roomRepository.FindRoom(id)

	if err != nil {
		panic(errortype.ErrResourceNotFound)
	}

	return room
}

func (roomService *roomService) CreateRoom(registerRoomDTO dto.RegisterRoomDTO) model.Room {
	if !roomService.canCreateRoom(registerRoomDTO.DormID, registerRoomDTO.DormOwnerID) {
		panic(errortype.ErrInvalidDormOwner)
	}

	room := mapRoom(registerRoomDTO)
	createdRoom, err := roomService.roomRepository.CreateRoom(room)

	if err != nil {
		panic(err)
	}

	return createdRoom
}

func (roomService *roomService) canCreateRoom(dormID string, dormOwnerID string) bool {
	dorm, err := roomService.dormRepository.FindDorm(dormID)

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(dorm.DormOwnerID) == dormOwnerID
}

func mapRoom(registerRoomDTO dto.RegisterRoomDTO) model.Room {
	dormID, _ := strconv.Atoi(registerRoomDTO.DormID)
	availableFrom, err := time.Parse("2006-01-02", *registerRoomDTO.AvailableFrom)
	room := model.Room{
		Name:        registerRoomDTO.Name,
		Price:       registerRoomDTO.Price,
		Size:        registerRoomDTO.Size,
		Description: registerRoomDTO.Description,
		Capacity:    registerRoomDTO.Capacity,
		DormID:      uint(dormID),
	}

	if err == nil {
		room.AvailableFrom = &availableFrom
	}

	return room
}
