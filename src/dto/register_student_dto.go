package dto

type RegisterStudentDTO struct {
	Firstname      string `json:"firstname"  validate:"required,min=2"`
	Lastname       string `json:"lastname" validate:"required,min=2"`
	StudentID      string `json:"student_id"  validate:"required,len=8,numeric"`
	Email          string `json:"email"  validate:"required,email"`
	Password       string `json:"password" validate:"required,min=4"`
	EnrollmentYear int    `json:"enrollment_year" validate:"required,gte=2014"`
	Gender         string `json:"gender"  validate:"required,oneof=male female lgbtq+"`
	Faculty        string `json:"faculty"  validate:"required,faculty"`
}
