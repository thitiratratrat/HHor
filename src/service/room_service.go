package service

import (
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type RoomService interface {
	GetAllRoomFacilities() []string
	GetRoom(id string) model.Room
}

func RoomServiceHandler(roomRepository repository.RoomRepository) RoomService {
	return &roomService{
		roomRepository: roomRepository,
	}
}

type roomService struct {
	roomRepository repository.RoomRepository
}

func (roomService *roomService) GetAllRoomFacilities() []string {
	roomFacilities := roomService.roomRepository.FindAllRoomFacilities()

	return roomFacilities
}

func (roomService *roomService) GetRoom(id string) model.Room {
	room, err := roomService.roomRepository.FindRoom(id)

	if err != nil {
		panic(errortype.ErrResourceNotFound.Error())
	}

	return room
}
