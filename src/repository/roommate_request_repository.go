package repository

import (
	"fmt"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoommateRequestRepository interface {
	FindRoommateRequestWithRegisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithRegisteredDorm
	FindRoommateRequestWithUnregisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithUnregisteredDorm
	CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoom model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error)
	CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error)
	CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error)
	UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithRegisteredDorm, error)
	UpdateRoommateRequestWithUnregisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithUnregisteredDorm, error)
}

func RoommateRequestRepositoryHandler(db *gorm.DB) RoommateRequestRepository {
	return &roommateRequestRepository{
		db: db,
	}
}

type roommateRequestRepository struct {
	db *gorm.DB
}

func (repository *roommateRequestRepository) FindRoommateRequestWithRegisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithRegisteredDorm {
	var roommateRequests []model.RoommateRequestWithRegisteredDorm
	dormNameCondition := `"Dorm".name` + " LIKE '%" + roommateRequestFilterDTO.DormName + "%'"

	condition := dormNameCondition + repository.getCondition(roommateRequestFilterDTO)

	repository.db.Preload("RoomPictures").Joins("Dorm").Joins("Room").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateRequestRepository) FindRoommateRequestWithUnregisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithUnregisteredDorm {
	var roommateRequests []model.RoommateRequestWithUnregisteredDorm
	dormNameCondition := "dorm_name LIKE '%" + roommateRequestFilterDTO.DormName + "%'"

	condition := dormNameCondition + repository.getCondition(roommateRequestFilterDTO)

	repository.db.Preload("RoomPictures").Preload("RoomFacilities").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateRequestRepository) CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoom model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error) {
	err := repository.db.Create(&roommateRequestWithNoRoom).Error

	return roommateRequestWithNoRoom, err
}

func (repository *roommateRequestRepository) CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error) {
	err := repository.db.Create(&roommateRequestWithRegisteredDorm).Error

	return roommateRequestWithRegisteredDorm, err
}

func (repository *roommateRequestRepository) CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error) {
	err := repository.db.Create(&roommateRequestWithUnregisteredDorm).Error

	return roommateRequestWithUnregisteredDorm, err
}

func (repository *roommateRequestRepository) UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithRegisteredDorm, error) {
	var roomPictures []model.RoommateRequestRegisteredDormPicture
	var roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm

	for _, pictureUrl := range pictureUrls {
		roomPictures = append(roomPictures, model.RoommateRequestRegisteredDormPicture{PictureUrl: pictureUrl,
			RoommateRequestWithRegisteredDormStudentID: id,
		})
	}

	repository.db.Table("roommate_request_registered_dorm_pictures").Where("roommate_request_with_registered_dorm_student_id = ?", id).Delete(model.RoommateRequestRegisteredDormPicture{})

	repository.db.Create(&roomPictures)

	err := repository.db.Preload("RoomPictures").Where("student_id = ?", id).First(&roommateRequestWithRegisteredDorm).Error

	return roommateRequestWithRegisteredDorm, err
}

func (repository *roommateRequestRepository) UpdateRoommateRequestWithUnregisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithUnregisteredDorm, error) {
	var roomPictures []model.RoommateRequestUnregisteredDormPicture
	var roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm

	for _, pictureUrl := range pictureUrls {
		roomPictures = append(roomPictures, model.RoommateRequestUnregisteredDormPicture{PictureUrl: pictureUrl,
			RoommateRequestWithUnregisteredDormStudentID: id,
		})
	}

	repository.db.Table("roommate_request_unregistered_dorm_pictures").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.RoommateRequestUnregisteredDormPicture{})

	repository.db.Create(&roomPictures)

	err := repository.db.Preload("RoomPictures").Preload("RoomFacilities").Where("student_id = ?", id).First(&roommateRequestWithUnregisteredDorm).Error

	return roommateRequestWithUnregisteredDorm, err
}

func (repository *roommateRequestRepository) getCondition(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) string {
	zoneCondition := repository.getZoneCondition(roommateRequestFilterDTO.Zone)
	condition := fmt.Sprintf("%s", zoneCondition)

	return condition
}

func (repository *roommateRequestRepository) getZoneCondition(zone string) string {
	if len(zone) == 0 {
		return ""
	}

	return fmt.Sprintf("AND dorm_zone_name = '%s'", zone)
}
