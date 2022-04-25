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
	CreateRoom(dormOwnerID string, registerRoomDTO dto.RegisterRoomDTO) model.Room
	UpdateRoom(id string, dormOwnerID string, updateRoomDTO dto.UpdateRoomDTO) model.Room
	UpdateRoomPictures(id string, pictures []string) model.Room
	DeleteRoom(id string, dormOwnerID string)
	CanUpdateRoom(roomID string, dormOwnerID string) bool
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

func (roomService *roomService) CreateRoom(dormOwnerID string, registerRoomDTO dto.RegisterRoomDTO) model.Room {
	if !roomService.canCreateRoom(registerRoomDTO.DormID, dormOwnerID) {
		panic(errortype.ErrInvalidDormOwner)
	}

	room := mapRoom(registerRoomDTO)
	createdRoom, err := roomService.roomRepository.CreateRoom(room)

	if err != nil {
		panic(err)
	}

	return createdRoom
}

func (roomService *roomService) UpdateRoom(id string, dormOwnerID string, updateRoomDTO dto.UpdateRoomDTO) model.Room {
	if !roomService.CanUpdateRoom(id, dormOwnerID) {
		panic(errortype.ErrInvalidDormOwner)
	}

	room := mapUpdateRoom(id, updateRoomDTO)
	updatedRoom, err := roomService.roomRepository.UpdateRoom(room)

	if err != nil {
		panic(err)
	}

	return updatedRoom
}

func (roomService *roomService) UpdateRoomPictures(id string, pictures []string) model.Room {
	updatedRoom, err := roomService.roomRepository.UpdateRoomPictures(id, pictures)

	if err != nil {
		panic(err)
	}

	return updatedRoom
}

func (roomService *roomService) DeleteRoom(id string, dormOwnerID string) {
	if !roomService.CanUpdateRoom(id, dormOwnerID) {
		panic(errortype.ErrInvalidDormOwner)
	}

	err := roomService.roomRepository.DeleteRoom(id)

	if err != nil {
		panic(err)
	}
}

func (roomService *roomService) CanUpdateRoom(roomID string, dormOwnerID string) bool {
	room, err := roomService.roomRepository.FindRoom(roomID)

	if err != nil {
		panic(err)
	}

	dorm, err := roomService.dormRepository.FindDorm(strconv.FormatUint(uint64(room.DormID), 10))

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(dorm.DormOwnerID) == dormOwnerID
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
	facilities := make([]model.AllRoomFacility, len(registerRoomDTO.Facilities))

	for index, facility := range registerRoomDTO.Facilities {
		facilities[index] = model.AllRoomFacility{
			Name: facility,
		}
	}

	room := model.Room{
		Name:        registerRoomDTO.Name,
		Price:       registerRoomDTO.Price,
		Size:        registerRoomDTO.Size,
		Description: registerRoomDTO.Description,
		Capacity:    registerRoomDTO.Capacity,
		DormID:      uint(dormID),
		Facilities:  facilities,
	}

	if registerRoomDTO.AvailableFrom != nil {
		availableFrom, err := time.Parse("2006-01-02", *registerRoomDTO.AvailableFrom)

		if err != nil {
			panic(err)
		}

		room.AvailableFrom = &availableFrom

	}

	return room
}

func mapUpdateRoom(roomID string, updateRoomDTO dto.UpdateRoomDTO) model.Room {
	id, _ := strconv.Atoi(roomID)
	facilities := make([]model.AllRoomFacility, len(updateRoomDTO.Facilities))

	for index, facility := range updateRoomDTO.Facilities {
		facilities[index] = model.AllRoomFacility{
			Name: facility,
		}
	}

	room := model.Room{
		ID:          uint(id),
		Name:        updateRoomDTO.Name,
		Price:       updateRoomDTO.Price,
		Size:        updateRoomDTO.Size,
		Description: updateRoomDTO.Description,
		Capacity:    updateRoomDTO.Capacity,
		Facilities:  facilities,
	}

	if updateRoomDTO.AvailableFrom != nil {
		availableFrom, err := time.Parse("2006-01-02", *updateRoomDTO.AvailableFrom)

		if err != nil {
			panic(err)
		}

		room.AvailableFrom = &availableFrom
	}

	return room

}
