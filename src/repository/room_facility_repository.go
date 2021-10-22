package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoomFacilityRepository interface {
	FindAllRoomFacilities() []string
}

func RoomFacilityRepositoryHandler(db *gorm.DB) RoomFacilityRepository {
	return &roomFacility{
		db: db,
	}
}

type roomFacility struct {
	db *gorm.DB
}

func (repository *roomFacility) FindAllRoomFacilities() []string {
	var facilities []string

	repository.db.Model(&model.AllRoomFacility{}).Pluck("name", &facilities)

	return facilities
}
