package dto

import "mime/multipart"

type RoommateRequestPictureDTO struct {
	RoomPictures *[]multipart.FileHeader `structs:",omitempty" form:"room_pictures,omitempty" swaggerignore:"true"`
}
