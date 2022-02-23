package dto

import "github.com/thitiratratrat/hhor/src/model"

type RoommateRequestWithRoomDTO struct {
	ID              string        `json:"id"`
	RoomPicture     string        `json:"picture"`
	DormName        string        `json:"dorm_name"`
	Student         model.Student `json:"student"`
	SharedRoomPrice int           `json:"shared_room_price"`
}
