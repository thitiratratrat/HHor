package service

import (
	"fmt"
	"strconv"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type DormService interface {
	GetDorms(dormFilterDTO dto.DormFilterDTO) []dto.DormDTO
	GetDorm(dormID string) model.Dorm
	GetDormSuggestions(firstLetter string) []dto.DormSuggestionDTO
	GetAllDormFacilities() []string
	GetDormZones() []string
	CreateDorm(dto.RegisterDormDTO) model.Dorm
	UpdateDorm(id string, dorm dto.UpdateDormDTO) model.Dorm
	UpdateDormPictures(id string, pictures []string) model.Dorm
	DeleteDorm(id string, dormOwnerID string)
	CanUpdateDorm(dormOwnerID string, dormID string) bool
}

func DormServiceHandler(dormRepository repository.DormRepository, roomRepository repository.RoomRepository) DormService {
	return &dormService{
		dormRepository: dormRepository,
		roomRepository: roomRepository,
	}
}

type dormService struct {
	dormRepository repository.DormRepository
	roomRepository repository.RoomRepository
}

func (dormService *dormService) GetDorms(dormFilterDTO dto.DormFilterDTO) []dto.DormDTO {
	dorms := dormService.dormRepository.FindDorms(dormFilterDTO)
	dormDTOs := []dto.DormDTO{}

	for _, dorm := range dorms {
		dormDTO := dto.DormDTO{}

		dormDTO.ID = fmt.Sprint(dorm.ID)

		if len(dorm.Pictures) != 0 {
			dormDTO.Picture = dorm.Pictures[0].PictureUrl
		}

		dormDTO.Name = dorm.Name
		dormDTO.StartingPrice = getCheapestRoomPrice(dorm.Rooms)
		dormDTO.Zone = dorm.DormZone.Name
		dormDTO.Latitude = dorm.Latitude
		dormDTO.Longitude = dorm.Longitude

		dormDTOs = append(dormDTOs, dormDTO)
	}

	return dormDTOs
}

func (dormService *dormService) GetDorm(dormID string) model.Dorm {
	dorm, err := dormService.dormRepository.FindDorm(dormID)

	if err != nil {
		panic(errortype.ErrResourceNotFound)
	}

	return dorm
}

func (dormService *dormService) GetDormSuggestions(firstLetter string) []dto.DormSuggestionDTO {
	dormNames := dormService.dormRepository.FindDormNames(firstLetter)

	return dormNames
}

func (dormService *dormService) GetAllDormFacilities() []string {
	dormFacilities := dormService.dormRepository.FindAllDormFacilities()

	return dormFacilities
}

func (dormService *dormService) GetDormZones() []string {
	dormZones := dormService.dormRepository.FindDormZones()

	return dormZones
}

func (dormService *dormService) CreateDorm(registerDormDTO dto.RegisterDormDTO) model.Dorm {
	dorm := mapCreateDorm(registerDormDTO)
	createdDorm, err := dormService.dormRepository.CreateDorm(dorm)

	if err != nil {
		panic(err)
	}

	return createdDorm
}

func (dormService *dormService) UpdateDorm(dormId string, updateDormDTO dto.UpdateDormDTO) model.Dorm {
	if !dormService.CanUpdateDorm(updateDormDTO.DormOwnerID, dormId) {
		panic(errortype.ErrInvalidDormOwner)
	}

	dorm := mapUpdateDorm(dormId, updateDormDTO)
	updatedDorm, err := dormService.dormRepository.UpdateDorm(dorm)

	if err != nil {
		panic(err)
	}

	return updatedDorm
}

func (dormService *dormService) UpdateDormPictures(id string, pictures []string) model.Dorm {
	updatedDorm, err := dormService.dormRepository.UpdateDormPictures(id, pictures)

	if err != nil {
		panic(err)
	}

	return updatedDorm
}

func (dormService *dormService) DeleteDorm(id string, dormOwnerID string) {
	if !dormService.CanUpdateDorm(dormOwnerID, id) {
		panic(errortype.ErrInvalidDormOwner)
	}

	dorm, err := dormService.dormRepository.FindDorm(id)

	if err != nil {
		panic(err)
	}

	for _, room := range dorm.Rooms {
		err := dormService.roomRepository.DeleteRoom(strconv.FormatUint(uint64(room.ID), 10))

		if err != nil {
			panic(err)
		}
	}

	err = dormService.dormRepository.DeleteDorm(id)

	if err != nil {
		panic(err)
	}
}

func getCheapestRoomPrice(rooms []model.Room) int {
	min := rooms[0].Price

	for _, room := range rooms {
		if room.Price < min {
			min = room.Price
		}
	}

	return min
}

func (dormService *dormService) CanUpdateDorm(dormOwnerID string, dormID string) bool {
	dorm, err := dormService.dormRepository.FindDorm(dormID)

	if err != nil {
		panic(err)
	}

	return strconv.Itoa(dorm.DormOwnerID) == dormOwnerID
}

func mapCreateDorm(registerDormDTO dto.RegisterDormDTO) model.Dorm {
	dormOwnerID, _ := strconv.Atoi(registerDormDTO.DormOwnerID)
	facilities := make([]model.AllDormFacility, len(registerDormDTO.Facilities))

	for index, facility := range registerDormDTO.Facilities {
		facilities[index] = model.AllDormFacility{
			Name: facility,
		}
	}

	return model.Dorm{
		Name:         registerDormDTO.Name,
		Type:         registerDormDTO.Type,
		Rules:        registerDormDTO.Rules,
		Longitude:    registerDormDTO.Long,
		Latitude:     registerDormDTO.Lat,
		Address:      registerDormDTO.Address,
		Description:  registerDormDTO.Description,
		DormZoneName: registerDormDTO.Zone,
		DormOwnerID:  dormOwnerID,
		Facilities:   facilities,
	}

}

func mapUpdateDorm(id string, updateDormDTO dto.UpdateDormDTO) model.Dorm {
	dormID, _ := strconv.Atoi(id)
	facilities := make([]model.AllDormFacility, len(updateDormDTO.Facilities))

	for index, facility := range updateDormDTO.Facilities {
		facilities[index] = model.AllDormFacility{
			Name: facility,
		}
	}

	return model.Dorm{
		ID:           uint(dormID),
		Name:         updateDormDTO.Name,
		Type:         updateDormDTO.Type,
		Rules:        updateDormDTO.Rules,
		Longitude:    updateDormDTO.Long,
		Latitude:     updateDormDTO.Lat,
		Address:      updateDormDTO.Address,
		Description:  updateDormDTO.Description,
		DormZoneName: updateDormDTO.Zone,
		Facilities:   facilities,
	}
}
