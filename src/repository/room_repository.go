package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoomRepository interface {
	FindAllRoomFacilities() []string
	FindRoom(id string) (model.Room, error)
	CreateRoom(model.Room) (model.Room, error)
}

func RoomRepositoryHandler(db *gorm.DB) RoomRepository {
	return &roomRepository{
		db: db,
	}
}

type roomRepository struct {
	db *gorm.DB
}

func (repository *roomRepository) FindAllRoomFacilities() []string {
	var facilities []string

	repository.db.Model(&model.AllRoomFacility{}).Pluck("name", &facilities)

	return facilities
}

func (repository *roomRepository) FindRoom(id string) (model.Room, error) {
	var room model.Room

	err := repository.db.Preload("Pictures").Preload("Facilities").Model(&model.Room{}).Where("id = ?", id).First(&room).Error

	return room, err
}

func (repository *roomRepository) CreateRoom(room model.Room) (model.Room, error) {
	err := repository.db.Create(&room).Error

	if err != nil {
		return model.Room{}, err
	}

	return room, err
}
