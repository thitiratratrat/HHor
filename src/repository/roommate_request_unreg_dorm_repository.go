package repository

import (
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoommateReqUnregDormRepository interface {
	FindRoommateReqUnregDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithUnregisteredDorm
	FindRoommateReqUnregDorm(id string) (model.RoommateRequestWithUnregisteredDorm, error)
	CreateRoommateReqUnregDorm(roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error)
	UpdateRoommateReqUnregDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithUnregisteredDorm, error)
	UpdateRoommateReqUnregDorm(id string, roommateRequest model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error)
	DeleteRoommateReqUnregDorm(id string) error
}

func RoommateReqUnregDormRepositoryHandler(db *gorm.DB) RoommateReqUnregDormRepository {
	return &roommateReqUnregDormRepository{
		db: db,
	}
}

type roommateReqUnregDormRepository struct {
	db *gorm.DB
}

func (repository *roommateReqUnregDormRepository) FindRoommateReqUnregDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithUnregisteredDorm {
	var roommateRequests []model.RoommateRequestWithUnregisteredDorm
	condition := getRoomCondition(roommateRequestRoomFilterDTO, constant.RoommateRequestUnregDorm)

	repository.db.Preload("RoomPictures").Preload("RoomFacilities").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateReqUnregDormRepository) FindRoommateReqUnregDorm(id string) (model.RoommateRequestWithUnregisteredDorm, error) {
	var roommateRequest model.RoommateRequestWithUnregisteredDorm

	err := repository.db.Preload("RoomPictures").Preload("RoomFacilities").Joins("Student").First(&roommateRequest, id).Error

	return roommateRequest, err
}

func (repository *roommateReqUnregDormRepository) CreateRoommateReqUnregDorm(roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error) {
	err := repository.db.Create(&roommateRequestWithUnregisteredDorm).Error

	return roommateRequestWithUnregisteredDorm, err
}

func (repository *roommateReqUnregDormRepository) UpdateRoommateReqUnregDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithUnregisteredDorm, error) {
	var roomPictures []model.RoommateRequestUnregisteredDormPicture
	var roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm

	for _, pictureUrl := range pictureUrls {
		roomPictures = append(roomPictures, model.RoommateRequestUnregisteredDormPicture{
			PictureUrl: pictureUrl,
			RoommateRequestWithUnregisteredDormStudentID: id,
		})
	}

	repository.db.Table("roommate_request_unregistered_dorm_pictures").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.RoommateRequestUnregisteredDormPicture{})
	repository.db.Create(&roomPictures)

	err := repository.db.Preload("RoomPictures").Preload("RoomFacilities").Where("student_id = ?", id).First(&roommateRequestWithUnregisteredDorm).Error

	return roommateRequestWithUnregisteredDorm, err
}

func (repository *roommateReqUnregDormRepository) UpdateRoommateReqUnregDorm(id string, roommateRequest model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error) {
	err := repository.db.Model(&model.RoommateRequestWithUnregisteredDorm{StudentID: roommateRequest.StudentID}).Omit("RoomFacilities").Updates(roommateRequest).Error

	if err != nil {
		return model.RoommateRequestWithUnregisteredDorm{}, err
	}

	repository.db.Table("roommate_request_unregistered_dorm_room_facility").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.AllRoomFacility{})
	repository.db.Model(&roommateRequest).Association("RoomFacilities").Append(roommateRequest.RoomFacilities)

	return repository.FindRoommateReqUnregDorm(id)
}

func (repository *roommateReqUnregDormRepository) DeleteRoommateReqUnregDorm(id string) error {
	repository.db.Table("roommate_request_unregistered_dorm_pictures").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.RoommateRequestUnregisteredDormPicture{})
	repository.db.Table("roommate_request_unregistered_dorm_room_facility").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.AllRoomFacility{})

	err := repository.db.Delete(&model.RoommateRequestWithUnregisteredDorm{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
