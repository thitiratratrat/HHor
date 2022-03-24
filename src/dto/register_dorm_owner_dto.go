package dto

type RegisterDormOwnerDTO struct {
	Firstname string `json:"firstname"  validate:"required,min=2"`
	Lastname  string `json:"lastname" validate:"required,min=2"`
	Email     string `json:"email"  validate:"required,email"`
	Password  string `json:"password" validate:"required,min=4"`
}
