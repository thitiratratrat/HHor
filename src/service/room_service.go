package service

import (
	"github.com/thitiratratrat/hhor/src/repository"
)

type RoomService interface {
	GetAllRoomFacilities() []string
}

func RoomServiceHandler(roomRepository repository.RoomFacilityRepository) RoomService {
	return &roomService{
		roomRepository: roomRepository,
	}
}

type roomService struct {
	roomRepository repository.RoomFacilityRepository
}

func (roomService *roomService) GetAllRoomFacilities() []string {
	roomFacilities := roomService.roomRepository.FindAllRoomFacilities()

	return roomFacilities
}
