package dto

import "mime/multipart"

type RoommateRequestWithRegisteredDormDTO struct {
	StudentID         string                  `json:"student_id" form:"student_id" validate:"required"`
	DormID            string                  `json:"dorm_id" form:"dorm_id" validate:"required,number"`
	RoomID            string                  `json:"room_id" form:"room_id" validate:"required,number"`
	NumberOfRoommates int                     `json:"number_of_roommates" form:"number_of_roommates" validate:"required,gt=0"`
	SharedRoomPrice   int                     `json:"shared_room_price" form:"shared_room_price" validate:"required,gt=0"`
	RoomPictures      *[]multipart.FileHeader `structs:",omitempty" form:"room_pictures,omitempty" swaggerignore:"true"`
}
