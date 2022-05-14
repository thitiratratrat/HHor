package repository

import (
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoommateReqRegDormRepository interface {
	FindRoommateReqRegDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithRegisteredDorm
	FindRoommateReqRegDorm(id string) (model.RoommateRequestWithRegisteredDorm, error)
	CreateRoommateReqRegDorm(roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error)
	UpdateRoommateReqRegDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithRegisteredDorm, error)
	UpdateRoommateReqRegDorm(id string, req model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error)
	DeleteRoommateReqRegDorm(id string) error
}

func RoommateReqRegDormRepositoryHandler(db *gorm.DB) RoommateReqRegDormRepository {
	return &roommateReqRegDormRepository{
		db: db,
	}
}

type roommateReqRegDormRepository struct {
	db *gorm.DB
}

func (repository *roommateReqRegDormRepository) FindRoommateReqRegDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithRegisteredDorm {
	var roommateRequests []model.RoommateRequestWithRegisteredDorm
	condition := getRoomCondition(roommateRequestRoomFilterDTO, constant.RoommateRequestRegDorm)

	repository.db.Preload("RoomPictures").Joins("Dorm").Joins("Room").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateReqRegDormRepository) FindRoommateReqRegDorm(id string) (model.RoommateRequestWithRegisteredDorm, error) {
	var roommateRequest model.RoommateRequestWithRegisteredDorm

	err := repository.db.Preload("Room.Pictures").Preload("Dorm.Pictures").Preload("Room.Facilities").Preload("RoomPictures").Joins("Dorm").Joins("Room").Joins("Student").First(&roommateRequest, id).Error

	return roommateRequest, err
}

func (repository *roommateReqRegDormRepository) CreateRoommateReqRegDorm(roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error) {
	err := repository.db.Create(&roommateRequestWithRegisteredDorm).Error

	return roommateRequestWithRegisteredDorm, err
}

func (repository *roommateReqRegDormRepository) UpdateRoommateReqRegDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithRegisteredDorm, error) {
	var roomPictures []model.RoommateRequestRegisteredDormPicture
	var roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm

	for _, pictureUrl := range pictureUrls {
		roomPictures = append(roomPictures, model.RoommateRequestRegisteredDormPicture{
			PictureUrl: pictureUrl,
			RoommateRequestWithRegisteredDormStudentID: id,
		})
	}

	repository.db.Table("roommate_request_registered_dorm_pictures").Where("roommate_request_with_registered_dorm_student_id = ?", id).Delete(model.RoommateRequestRegisteredDormPicture{})
	repository.db.Create(&roomPictures)

	err := repository.db.Preload("RoomPictures").Where("student_id = ?", id).First(&roommateRequestWithRegisteredDorm).Error

	return roommateRequestWithRegisteredDorm, err
}

func (repository *roommateReqRegDormRepository) UpdateRoommateReqRegDorm(id string, roommateRequest model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error) {
	err := repository.db.Save(&roommateRequest).Error

	if err != nil {
		return model.RoommateRequestWithRegisteredDorm{}, err
	}

	return repository.FindRoommateReqRegDorm(id)
}

func (repository *roommateReqRegDormRepository) DeleteRoommateReqRegDorm(id string) error {
	repository.db.Table("roommate_request_registered_dorm_pictures").Where("roommate_request_with_registered_dorm_student_id = ?", id).Delete(model.RoommateRequestRegisteredDormPicture{})

	err := repository.db.Delete(&model.RoommateRequestWithRegisteredDorm{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
