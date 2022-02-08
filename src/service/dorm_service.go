package service

import (
	"fmt"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type DormService interface {
	GetDorms(dormFilterDTO dto.DormFilterDTO) []dto.DormDTO
	GetDorm(dormID string) (model.Dorm, error)
	GetDormSuggestions(firstLetter string) []dto.DormSuggestionDTO
	GetAllDormFacilities() []string
	GetDormZones() []string
}

func DormServiceHandler(dormRepository repository.DormRepository) DormService {
	return &dormService{
		dormRepository: dormRepository,
	}
}

type dormService struct {
	dormRepository repository.DormRepository
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

func (dormService *dormService) GetDorm(dormID string) (model.Dorm, error) {
	dorm, err := dormService.dormRepository.FindDorm(dormID)

	return dorm, err
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

func getCheapestRoomPrice(rooms []model.Room) int {
	min := rooms[0].Price

	for _, room := range rooms {
		if room.Price < min {
			min = room.Price
		}
	}

	return min
}
