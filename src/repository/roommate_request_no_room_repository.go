package repository

import (
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoommateReqNoRoomRepository interface {
	FindRoommateReqNoRooms(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithNoRoom
	FindRoommateReqNoRoom(id string) (model.RoommateRequestWithNoRoom, error)
	CreateRoommateReqNoRoom(roommateRequestWithNoRoom model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error)
	UpdateRoommateReqNoRoom(id string, roommateRequest model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error)
	DeleteRoommateReqNoRoom(id string) error
}

func RoommateReqNoRoomRepositoryHandler(db *gorm.DB) RoommateReqNoRoomRepository {
	return &roommateReqNoRoomRepository{
		db: db,
	}
}

type roommateReqNoRoomRepository struct {
	db *gorm.DB
}

func (repository *roommateReqNoRoomRepository) FindRoommateReqNoRooms(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithNoRoom {
	var roommateRequests []model.RoommateRequestWithNoRoom
	condition := getCondition(roommateRequestFilterDTO, constant.RoommateRequestNoRoom)

	repository.db.Preload("Zones").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateReqNoRoomRepository) FindRoommateReqNoRoom(id string) (model.RoommateRequestWithNoRoom, error) {
	var roommateRequest model.RoommateRequestWithNoRoom

	err := repository.db.Preload("Zones").Joins("Student").First(&roommateRequest, id).Error

	return roommateRequest, err
}

func (repository *roommateReqNoRoomRepository) CreateRoommateReqNoRoom(roommateRequestWithNoRoom model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error) {
	err := repository.db.Create(&roommateRequestWithNoRoom).Error

	return roommateRequestWithNoRoom, err
}

func (repository *roommateReqNoRoomRepository) UpdateRoommateReqNoRoom(id string, roommateRequest model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error) {
	err := repository.db.Model(&model.RoommateRequestWithNoRoom{StudentID: roommateRequest.StudentID}).Omit("Zones").Updates(roommateRequest).Error

	if err != nil {
		return model.RoommateRequestWithNoRoom{}, err
	}

	repository.db.Table("roommate_request_no_room_zone").Where("roommate_request_with_no_room_student_id = ?", id).Delete(model.DormZone{})
	repository.db.Model(&roommateRequest).Association("Zones").Append(roommateRequest.Zones)

	return repository.FindRoommateReqNoRoom(id)
}

func (repository *roommateReqNoRoomRepository) DeleteRoommateReqNoRoom(id string) error {
	repository.db.Table("roommate_request_no_room_zone").Where("roommate_request_with_no_room_student_id = ?", id).Delete(model.DormZone{})
	err := repository.db.Delete(&model.RoommateRequestWithNoRoom{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
