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
	FindDormNames(firstLetter string) []dto.DormSuggestionDTO
	FindAllDormFacilities() []string
	FindDormZones() []string
	CreateDorm(model.Dorm) (model.Dorm, error)
	UpdateDorm(dorm model.Dorm) (model.Dorm, error)
	UpdateNearbyLocations(id string, nearbyLocations []model.NearbyLocation) (model.Dorm, error)
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

	nameCondition := repository.getNameCondition(dormFilterDTO.Name)
	typeCondition := repository.getTypeCondition(dormFilterDTO.Type)
	zoneCondition := repository.getZoneCondition(dormFilterDTO.Zone)
	capacityCondition := repository.getCapacityCondition(dormFilterDTO.Capacity)
	latLongCondition := repository.getLatLongCondition(dormFilterDTO.Lat, dormFilterDTO.Long)
	priceCondition := repository.getPriceCondition(dormFilterDTO.LowerPrice, dormFilterDTO.UpperPrice)
	roomFacilitiesCondition := repository.getRoomFacilitiesCondition(dormFilterDTO.RoomFacilities)
	dormFacilitiesCondition := repository.getDormFacilitiesCondition(dormFilterDTO.DormFacilities)

	roomWhereCondition := fmt.Sprintf("name LIKE '%%' %s %s %s", capacityCondition, priceCondition, roomFacilitiesCondition)
	dormWhereCondition := fmt.Sprintf("%s %s %s %s %s", nameCondition, typeCondition, zoneCondition, latLongCondition, dormFacilitiesCondition)

	repository.db.Preload("Rooms", roomWhereCondition, func(db *gorm.DB) *gorm.DB {
		return db.Order("rooms.price ASC")
	}).Preload("Pictures").Preload("DormZone").Where(dormWhereCondition).Find(&dorms)

	return dorms
}

func (repository *dormRepository) FindDorm(dormID string) (model.Dorm, error) {
	var dorm model.Dorm

	err := repository.db.Preload("Pictures").Preload("Rooms", func(db *gorm.DB) *gorm.DB {
		return db.Order("rooms.price ASC")
	}).Preload("Rooms.Pictures").Preload("Rooms.Facilities").Preload("Facilities").Preload("NearbyLocations").First(&dorm, dormID).Error

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
	err := repository.db.Model(&model.Dorm{}).Where("id = ?", dorm.ID).Select("Name", "Type", "Rules", "Longitude", "Latitude", "Address", "Description", "DormZoneName").Updates(dorm).Error

	if err != nil {
		return model.Dorm{}, err
	}

	repository.db.Table("dorm_facility").Where("dorm_id = ?", dorm.ID).Delete(model.AllDormFacility{})
	repository.db.Model(&dorm).Association("Facilities").Append(dorm.Facilities)

	return repository.FindDorm(fmt.Sprint(dorm.ID))
}

func (repository *dormRepository) UpdateNearbyLocations(id string, nearbyLocations []model.NearbyLocation) (model.Dorm, error) {
	repository.db.Table("nearby_locations").Where("dorm_id = ?", id).Delete(model.NearbyLocation{})
	err := repository.db.Create(&nearbyLocations).Error

	if err != nil {
		return model.Dorm{}, err
	}

	return repository.FindDorm(id)
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
	err := repository.db.Create(&dormPictures).Error

	if err != nil {
		return model.Dorm{}, err
	}

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

func (repository *dormRepository) getNameCondition(name *string) string {
	if name == nil {
		return "name LIKE '%%'"
	}

	if len(*name) == 1 {
		return "name" + ` LIKE '` + *name + `%'`
	}

	return fmt.Sprintf("'%s'", *name) + "% ANY(STRING_TO_ARRAY(name,' '))"
}

func (repository *dormRepository) getTypeCondition(typeFilter []string) string {
	if len(typeFilter) == 0 {
		return ""
	}

	formattedTypeFilter := "'" + strings.Join(typeFilter, "', '") + "'"

	return fmt.Sprintf("AND type in (%s)", formattedTypeFilter)
}

func (repository *dormRepository) getZoneCondition(zone *string) string {
	if zone == nil {
		return ""
	}

	return fmt.Sprintf("AND dorm_zone_name = '%s'", *zone)
}

func (repository *dormRepository) getCapacityCondition(capacity *int) string {
	if capacity == nil {
		return ""
	}

	return fmt.Sprintf("AND capacity >= %d", *capacity)
}

func (repository *dormRepository) getLatLongCondition(lat *float64, long *float64) string {
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

func (repository *dormRepository) getPriceCondition(lowerPrice *int, upperPrice *int) string {
	if lowerPrice == nil || upperPrice == nil || *upperPrice < *lowerPrice {
		return ""
	}

	return fmt.Sprintf("AND price BETWEEN %d AND %d", *lowerPrice, *upperPrice)
}

func (repository *dormRepository) getRoomFacilitiesCondition(roomFacilities []string) string {
	if len(roomFacilities) == 0 {
		return ""
	}

	formattedRoomFacilities := "'" + strings.Join(roomFacilities, "', '") + "'"

	return fmt.Sprintf("AND %d = (select count(*) from room_facility where rooms.id = room_facility.room_id and all_room_facility_name IN (%s))", len(roomFacilities), formattedRoomFacilities)
}

func (repository *dormRepository) getDormFacilitiesCondition(dormFacilities []string) string {
	if len(dormFacilities) == 0 {
		return ""
	}

	formattedDormFacilities := "'" + strings.Join(dormFacilities, "', '") + "'"

	return fmt.Sprintf("AND %d = (select count(*) from dorm_facility where dorms.id = dorm_facility.dorm_id and all_dorm_facility_name IN (%s))", len(dormFacilities), formattedDormFacilities)
}
