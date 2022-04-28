package dto

import "mime/multipart"

type RegisterDormOwnerDTO struct {
	ProfilePicture *multipart.FileHeader `structs:",omitempty" form:"profile_picture,omitempty" swaggerignore:"true"`
	Firstname      string                `json:"firstname" form:"firstname" validate:"required,min=2"`
	Lastname       string                `json:"lastname" form:"lastname" validate:"required,min=2"`
	Email          string                `json:"email" form:"email" validate:"required,email"`
	Password       string                `json:"password" form:"password" validate:"required,min=4"`
	LineID         string                `json:"line_id" form:"line_id" validate:"required,min=1"`
	PhoneNumber    string                `json:"phone_number" form:"phone_number" validate:"required,numeric,phone"`
}
