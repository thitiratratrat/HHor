package dto

import "mime/multipart"

type RegisterStudentDTO struct {
	ProfilePicture *multipart.FileHeader `structs:",omitempty" form:"profile_picture,omitempty" swaggerignore:"true"`
	Firstname      string                `json:"firstname" form:"firstname" validate:"required,min=2"`
	Lastname       string                `json:"lastname" form:"lastname" validate:"required,min=2"`
	StudentID      string                `json:"student_id" form:"student_id" validate:"required,len=8,numeric"`
	Email          string                `json:"email" form:"email" validate:"required,email"`
	Password       string                `json:"password" form:"password" validate:"required,min=4"`
	EnrollmentYear int                   `json:"enrollment_year" form:"enrollment_year" validate:"required,gte=2014"`
	Gender         string                `json:"gender" form:"gender" validate:"required,oneof=male female lgbtq+"`
	Faculty        string                `json:"faculty" form:"faculty" validate:"required,faculty"`
}
