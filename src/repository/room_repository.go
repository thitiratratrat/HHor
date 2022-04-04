package repository

import (
	"strconv"

	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoomRepository interface {
	FindAllRoomFacilities() []string
	FindRoom(id string) (model.Room, error)
	CreateRoom(model.Room) (model.Room, error)
	UpdateRoom(model.Room) (model.Room, error)
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

func (repository *roomRepository) UpdateRoom(room model.Room) (model.Room, error) {
	err := repository.db.Model(&model.Room{ID: room.ID}).Select("Name", "Price", "Size", "Description", "Capacity", "AvailableFrom").Updates(room).Error

	if err != nil {
		return model.Room{}, err
	}

	repository.db.Table("room_facility").Where("room_id = ?", room.ID).Delete(model.AllRoomFacility{})
	repository.db.Model(&room).Association("Facilities").Append(room.Facilities)

	return repository.FindRoom(strconv.FormatUint(uint64(room.ID), 10))
}
