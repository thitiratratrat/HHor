package dto

import "mime/multipart"

type DormRoomPicturesDTO struct {
	Pictures *[]multipart.FileHeader `structs:",omitempty" form:"pictures,omitempty" swaggerignore:"true"`
}
