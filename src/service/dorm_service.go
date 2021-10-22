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
	GetDormNames(firstLetter string) []string
	GetAllDormFacilities() []string
}

func DormServiceHandler(dormRepository repository.DormRepository, dormFacilityRepository repository.DormFacilityRepository) DormService {
	return &dormService{
		dormRepository:         dormRepository,
		dormFacilityRepository: dormFacilityRepository,
	}
}

type dormService struct {
	dormRepository         repository.DormRepository
	dormFacilityRepository repository.DormFacilityRepository
}

func (dormService *dormService) GetDorms(dormFilterDTO dto.DormFilterDTO) []dto.DormDTO {
	dorms := dormService.dormRepository.FindDorms(dormFilterDTO)
	dormDTOs := []dto.DormDTO{}

	for _, dorm := range dorms {
		dormDTO := dto.DormDTO{}

		dormDTO.ID = fmt.Sprint(dorm.ID)
		dormDTO.Picture = dorm.Pictures[0].PictureUrl
		dormDTO.Name = dorm.Name
		dormDTO.StartingPrice = getCheapestRoomPrice(dorm.Rooms)
		dormDTO.Zone = dorm.DormZone.Name

		dormDTOs = append(dormDTOs, dormDTO)
	}

	return dormDTOs
}

func (dormService *dormService) GetDorm(dormID string) (model.Dorm, error) {
	dorm, err := dormService.dormRepository.FindDorm(dormID)

	return dorm, err
}

func (dormService *dormService) GetDormNames(firstLetter string) []string {
	dormNames := dormService.dormRepository.FindDormNames(firstLetter)

	return dormNames
}

func (dormService *dormService) GetAllDormFacilities() []string {
	dormFacilities := dormService.dormFacilityRepository.FindAllDormFacilities()

	return dormFacilities
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
