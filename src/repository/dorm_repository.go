package repository

import (
	"fmt"
	"strings"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/utils"
	"gorm.io/gorm"
)

type DormRepository interface {
	FindDorms(dormFilterDTO dto.DormFilterDTO) []model.Dorm
	FindDorm(dormID string) (model.Dorm, error)
	FindDormNames(firstLetter string) []dto.DormSuggestionDTO
}

func DormRepositoryHandler(db *gorm.DB) DormRepository {
	return &dormRepository{
		db: db,
	}
}

type dormRepository struct {
	db *gorm.DB
}

const distance = 5 //km

func (repository *dormRepository) FindDorms(dormFilterDTO dto.DormFilterDTO) []model.Dorm {
	var dorms []model.Dorm

	nameCondition := "name LIKE '%" + dormFilterDTO.Name + "%'"
	typeCondition := getTypeCondition(dormFilterDTO.Type)
	zoneCondition := getZoneCondition(dormFilterDTO.Zone)
	capacityCondition := getCapacityCondition(dormFilterDTO.Capacity)
	latLongCondition := getLatLongCondition(dormFilterDTO.Lat, dormFilterDTO.Long)
	priceCondition := getPriceCondition(dormFilterDTO.LowerPrice, dormFilterDTO.UpperPrice)
	roomFacilitiesCondition := getRoomFacilitiesCondition(dormFilterDTO.RoomFacilities)
	dormFacilitiesCondition := getDormFacilitiesCondition(dormFilterDTO.DormFacilities)

	roomWhereCondition := fmt.Sprintf("AND 0 != (select count(*) from rooms where dorms.id = rooms.dorm_id %s %s %s)", capacityCondition, priceCondition, roomFacilitiesCondition)
	dormWhereCondition := fmt.Sprintf("%s %s %s %s %s %s", nameCondition, typeCondition, zoneCondition, latLongCondition, dormFacilitiesCondition, roomWhereCondition)

	repository.db.Preload("Rooms").Preload("Pictures").Preload("DormZone").Where(dormWhereCondition).Find(&dorms)

	return dorms
}

func (repository *dormRepository) FindDorm(dormID string) (model.Dorm, error) {
	var dorm model.Dorm
	var nearbyLocations []model.NearbyLocation

	err := repository.db.Preload("Pictures").Preload("Rooms", "available = ?", "TRUE").Preload("Rooms.Pictures").Preload("Rooms.Facilities").Preload("Account").Preload("Facilities").First(&dorm, dormID).Error
	repository.db.Where("dorm_id = ?", dormID).Preload("Location").Find(&nearbyLocations)

	dorm.NearbyLocations = nearbyLocations

	return dorm, err
}

func (repository *dormRepository) FindDormNames(firstLetter string) []dto.DormSuggestionDTO {
	var dormNames []dto.DormSuggestionDTO

	repository.db.Table("dorms").Where("name LIKE ?", firstLetter+"%").Select("name", "id").Find(&dormNames)

	return dormNames
}

func getTypeCondition(typeFilter []string) string {
	if len(typeFilter) == 0 {
		return ""
	}

	formattedTypeFilter := "'" + strings.Join(typeFilter, "', '") + "'"

	return fmt.Sprintf("AND type in (%s)", formattedTypeFilter)
}

func getZoneCondition(zone string) string {
	if len(zone) == 0 {
		return ""
	}

	return fmt.Sprintf("AND dorm_zone_name = '%s'", zone)
}

func getCapacityCondition(capacity int) string {
	if capacity == 0 {
		return ""
	}

	return fmt.Sprintf("AND capacity >= %d", capacity)
}

func getLatLongCondition(lat float64, long float64) string {
	if lat == 0 || long == 0 {
		return ""
	}

	minLat, _ := utils.GetLatLongFromDistance(lat, long, distance, 180)
	maxLat, _ := utils.GetLatLongFromDistance(lat, long, distance, 0)
	_, minLong := utils.GetLatLongFromDistance(lat, long, distance, 270)
	_, maxLong := utils.GetLatLongFromDistance(lat, long, distance, 90)

	return fmt.Sprintf("AND latitude BETWEEN %f AND %f AND longitude BETWEEN %f AND %f", minLat, maxLat, minLong, maxLong)
}

func getPriceCondition(lowerPrice int, upperPrice int) string {
	if lowerPrice == 0 || upperPrice == 0 || upperPrice < lowerPrice {
		return ""
	}

	return fmt.Sprintf("AND price BETWEEN %d AND %d", lowerPrice, upperPrice)
}

func getRoomFacilitiesCondition(roomFacilities []string) string {
	if len(roomFacilities) == 0 {
		return ""
	}

	formattedRoomFacilities := "'" + strings.Join(roomFacilities, "', '") + "'"

	return fmt.Sprintf("AND %d = (select count(*) from room_facility where rooms.id = room_facility.room_id and all_room_facility_name IN (%s))", len(roomFacilities), formattedRoomFacilities)
}

func getDormFacilitiesCondition(dormFacilities []string) string {
	if len(dormFacilities) == 0 {
		return ""
	}

	formattedDormFacilities := "'" + strings.Join(dormFacilities, "', '") + "'"

	return fmt.Sprintf("AND %d = (select count(*) from dorm_facility where dorms.id = dorm_facility.dorm_id and all_dorm_facility_name IN (%s))", len(dormFacilities), formattedDormFacilities)
}
