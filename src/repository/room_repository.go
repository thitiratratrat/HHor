package repository

import (
	"fmt"
	"strconv"

	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoomRepository interface {
	FindAllRoomFacilities() []string
	FindRoom(id string) (model.Room, error)
	CreateRoom(model.Room) (model.Room, error)
	UpdateRoom(model.Room) (model.Room, error)
	UpdateRoomPictures(id string, pictureUrls []string) (model.Room, error)
	DeleteRoom(id string) error
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

	return repository.FindRoom(fmt.Sprintf("%v", room.ID))
}

func (repository *roomRepository) UpdateRoomPictures(id string, pictureUrls []string) (model.Room, error) {
	var roomPictures []model.RoomPicture
	roomID, _ := strconv.Atoi(id)

	for _, pictureUrl := range pictureUrls {
		roomPictures = append(roomPictures, model.RoomPicture{
			PictureUrl: pictureUrl,
			RoomID:     uint(roomID),
		})
	}

	repository.db.Table("room_pictures").Where("room_id = ?", id).Delete(model.RoomPicture{})
	err := repository.db.Create(&roomPictures).Error

	if err != nil {
		return model.Room{}, err
	}

	return repository.FindRoom(id)
}

func (repository *roomRepository) DeleteRoom(id string) error {
	repository.db.Table("room_pictures").Where("room_id = ?", id).Delete(model.RoomPicture{})
	repository.db.Table("room_facility").Where("room_id = ?", id).Delete(model.AllRoomFacility{})

	err := repository.db.Delete(&model.Room{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
