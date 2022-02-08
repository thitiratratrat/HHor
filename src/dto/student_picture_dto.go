package dto

import "mime/multipart"

type StudentPictureDTO struct {
	ProfilePicture *multipart.FileHeader   `structs:",omitempty" form:"profile_picture,omitempty" swaggerignore:"true"`
	PetPictures    *[]multipart.FileHeader `structs:",omitempty" form:"pet_pictures,omitempty" swaggerignore:"true"`
}
