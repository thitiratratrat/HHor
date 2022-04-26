package dto

type ResendCodeDTO struct {
	Email string `json:"email" validate:"required,email"`
}
