package dto

type RegisterStudentDTO struct {
	Firstname   string `json:"firstname" validate:"required,min=2"`
	Lastname    string `json:"lastname" validate:"required,min=2"`
	StudentID   string `json:"student_id" validate:"required,len=8,numeric"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=4"`
	YearOfStudy int    `json:"year_of_study" validate:"required,gte=1,lte=8"`
	Gender      string `json:"gender" validate:"required,oneof=male female lgbtq"`
	Faculty     string `json:"faculty" validate:"required,faculty"`
}
