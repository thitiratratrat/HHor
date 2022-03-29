package dto

import "mime/multipart"

type PicturesDTO struct {
	Pictures *[]multipart.FileHeader `structs:",omitempty" form:"pictures,omitempty" swaggerignore:"true"`
}
