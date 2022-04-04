package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/utils"
	"gorm.io/gorm"
)

type DormRepository interface {
	FindDorms(dormFilterDTO dto.DormFilterDTO) []model.Dorm
	FindDorm(dormID string) (model.Dorm, error)
	FindDormOwnerDorms(dormOwnerID string) []model.Dorm
	FindDormNames(firstLetter string) []dto.DormSuggestionDTO
	FindAllDormFacilities() []string
	FindDormZones() []string
	CreateDorm(model.Dorm) (model.Dorm, error)
	UpdateDorm(dorm model.Dorm) (model.Dorm, error)
	UpdateDormPictures(id string, pictureUrls []string) (model.Dorm, error)
	DeleteDorm(id string) error
}

func DormRepositoryHandler(db *gorm.DB) DormRepository {
	return &dormRepository{
		db: db,
	}
}

type dormRepository struct {
	db *gorm.DB
}

func (repository *dormRepository) FindDorms(dormFilterDTO dto.DormFilterDTO) []model.Dorm {
	var dorms []model.Dorm

	nameCondition := getNameCondition(dormFilterDTO.Name)
	typeCondition := getTypeCondition(dormFilterDTO.Type)
	zoneCondition := getZoneCondition(dormFilterDTO.Zone)
	capacityCondition := getCapacityCondition(dormFilterDTO.Capacity)
	latLongCondition := getLatLongCondition(dormFilterDTO.Lat, dormFilterDTO.Long)
	priceCondition := getPriceCondition(dormFilterDTO.LowerPrice, dormFilterDTO.UpperPrice)
	roomFacilitiesCondition := getRoomFacilitiesCondition(dormFilterDTO.RoomFacilities)
	dormFacilitiesCondition := getDormFacilitiesCondition(dormFilterDTO.DormFacilities)

	roomWhereCondition := fmt.Sprintf("AND 0 != (select count(*) from rooms where dorms.id = rooms.dorm_id and available_from IS NOT NULL %s %s %s)", capacityCondition, priceCondition, roomFacilitiesCondition)
	dormWhereCondition := fmt.Sprintf("%s %s %s %s %s %s", nameCondition, typeCondition, zoneCondition, latLongCondition, dormFacilitiesCondition, roomWhereCondition)

	repository.db.Preload("Rooms").Preload("Pictures").Preload("DormZone").Where(dormWhereCondition).Find(&dorms)

	return dorms
}

func (repository *dormRepository) FindDormOwnerDorms(dormOwnerID string) []model.Dorm {
	var dorms []model.Dorm

	repository.db.Preload("Rooms").Preload("Pictures").Preload("DormZone").Where("owner = ?", dormOwnerID).Find(&dorms)

	return dorms
}

func (repository *dormRepository) FindDorm(dormID string) (model.Dorm, error) {
	var dorm model.Dorm
	var nearbyLocations []model.NearbyLocation

	err := repository.db.Preload("Pictures").Preload("Rooms", "available_from IS NOT NULL").Preload("Rooms.Pictures").Preload("Rooms.Facilities").Preload("DormOwner").Preload("Facilities").First(&dorm, dormID).Error
	repository.db.Where("dorm_id = ?", dormID).Preload("Location").Find(&nearbyLocations)

	dorm.NearbyLocations = nearbyLocations

	return dorm, err
}

func (repository *dormRepository) FindDormNames(firstLetter string) []dto.DormSuggestionDTO {
	var dormNames []dto.DormSuggestionDTO

	repository.db.Table("dorms").Where("name LIKE ?", firstLetter+"%").Select("name", "id").Find(&dormNames)

	return dormNames
}

func (repository *dormRepository) FindAllDormFacilities() []string {
	var facilities []string

	repository.db.Model(&model.AllDormFacility{}).Pluck("name", &facilities)

	return facilities
}

func (repository *dormRepository) FindDormZones() []string {
	var zones []string

	repository.db.Model(&model.DormZone{}).Pluck("name", &zones)

	return zones
}

func (repository *dormRepository) CreateDorm(dorm model.Dorm) (model.Dorm, error) {
	err := repository.db.Create(&dorm).Error

	if err != nil {
		return model.Dorm{}, err
	}

	return dorm, err
}

func (repository *dormRepository) UpdateDorm(dorm model.Dorm) (model.Dorm, error) {
	err := repository.db.Model(&model.Dorm{}).Where("id = ?", dorm.ID).Select("Name", "Type", "Rules", "Longitude", "Latitude", "Address", "Description", "Zone").Updates(dorm).Error

	if err != nil {
		return model.Dorm{}, err
	}

	repository.db.Table("dorm_facility").Where("dorm_id = ?", dorm.ID).Delete(model.AllDormFacility{})
	repository.db.Model(&dorm).Association("Facilities").Append(dorm.Facilities)

	return repository.FindDorm(strconv.FormatUint(uint64(dorm.ID), 10))
}

func (repository *dormRepository) UpdateDormPictures(id string, pictureUrls []string) (model.Dorm, error) {
	var dormPictures []model.DormPicture
	dormId, _ := strconv.Atoi(id)

	for _, pictureUrl := range pictureUrls {
		dormPictures = append(dormPictures, model.DormPicture{
			PictureUrl: pictureUrl,
			DormID:     uint(dormId),
		})
	}

	repository.db.Table("dorm_pictures").Where("dorm_id = ?", id).Delete(model.DormPicture{})
	repository.db.Create(&dormPictures)

	return repository.FindDorm(id)
}

func (repository *dormRepository) DeleteDorm(id string) error {
	repository.db.Table("dorm_pictures").Where("dorm_id = ?", id).Delete(model.DormPicture{})
	repository.db.Table("dorm_facility").Where("dorm_id = ?", id).Delete(model.AllDormFacility{})
	repository.db.Table("nearby_locations").Where("dorm_id = ?", id).Delete(model.NearbyLocation{})

	err := repository.db.Delete(&model.Dorm{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func getNameCondition(name *string) string {
	if name == nil {
		return "name LIKE '%%'"
	}

	return "name LIKE '%" + *name + "%'"
}

func getTypeCondition(typeFilter []string) string {
	if len(typeFilter) == 0 {
		return ""
	}

	formattedTypeFilter := "'" + strings.Join(typeFilter, "', '") + "'"

	return fmt.Sprintf("AND type in (%s)", formattedTypeFilter)
}

func getZoneCondition(zone *string) string {
	if zone == nil {
		return ""
	}

	return fmt.Sprintf("AND dorm_zone_name = '%s'", *zone)
}

func getCapacityCondition(capacity *int) string {
	if capacity == nil {
		return ""
	}

	return fmt.Sprintf("AND capacity >= %d", *capacity)
}

func getLatLongCondition(lat *float64, long *float64) string {
	if lat == nil || long == nil {
		return ""
	}

	const distance = 5 //km
	minLat, _ := utils.GetLatLongFromDistance(*lat, *long, distance, 180)
	maxLat, _ := utils.GetLatLongFromDistance(*lat, *long, distance, 0)
	_, minLong := utils.GetLatLongFromDistance(*lat, *long, distance, 270)
	_, maxLong := utils.GetLatLongFromDistance(*lat, *long, distance, 90)

	return fmt.Sprintf("AND latitude BETWEEN %f AND %f AND longitude BETWEEN %f AND %f", minLat, maxLat, minLong, maxLong)
}

func getPriceCondition(lowerPrice *int, upperPrice *int) string {
	if lowerPrice == nil || upperPrice == nil || *upperPrice < *lowerPrice {
		return ""
	}

	return fmt.Sprintf("AND price BETWEEN %d AND %d", *lowerPrice, *upperPrice)
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
