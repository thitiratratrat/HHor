package dto

import "mime/multipart"

type StudentUpdateDTO struct {
	Email          string                `structs:",omitempty" form:"email" validate:"required,email"`
	ProfilePicture *multipart.FileHeader `structs:",omitempty" form:"profile_picture,omitempty" swaggerignore:"true"`
	FacebookUrl    *string               `structs:",omitempty" form:"facebook_url,omitempty" validate:"omitempty,url,contains=facebook"`
	TwitterUrl     *string               `structs:",omitempty" form:"twitter_url,omitempty" validate:"omitempty,url,contains=twitter"`
	LinkedinUrl    *string               `structs:",omitempty" form:"linkedin_url,omitempty" validate:"omitempty,url,contains=linkedin"`
	InstagramUrl   *string               `structs:",omitempty" form:"instagram_url,omitempty" validate:"omitempty,url,contains=instagram"`
	Biography      *string               `structs:",omitempty" form:"biography,omitempty"`
}
