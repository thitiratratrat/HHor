package dto

import "mime/multipart"

type RegisterStudentDTO struct {
	Firstname      string                `json:"firstname" form:"firstname" validate:"required,min=2"`
	Lastname       string                `json:"lastname" form:"lastname" validate:"required,min=2"`
	StudentID      string                `json:"student_id" form:"student_id" validate:"required,len=8,numeric"`
	Email          string                `json:"email"  form:"email" validate:"required,email"`
	Password       string                `json:"password" form:"password" validate:"required,min=4"`
	YearOfStudy    int                   `json:"year_of_study" form:"year_of_study" validate:"required,gte=1,lte=8"`
	Gender         string                `json:"gender" form:"gender" validate:"required,oneof=male female lgbtq"`
	Faculty        string                `json:"faculty" form:"faculty" validate:"required,faculty"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture" swaggerignore:"true"`
}
