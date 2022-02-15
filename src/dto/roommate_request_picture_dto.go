package dto

import "mime/multipart"

type RoommateRequestPictureDTO struct {
	StudentID    string                  `json:"student_id" form:"student_id" validate:"required"`
	RoomPictures *[]multipart.FileHeader `structs:",omitempty" form:"room_pictures,omitempty" swaggerignore:"true"`
}
