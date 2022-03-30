package dto

import "mime/multipart"

type DormPicturesDTO struct {
	Pictures    *[]multipart.FileHeader `structs:",omitempty" form:"pictures,omitempty" swaggerignore:"true"`
	DormOwnerID string                  `json:"owner_id" form:"owner_id" validate:"required,numeric"`
}
