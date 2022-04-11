package dto

type UpdateDormOwnerDTO struct {
	Firstname   string `json:"firstname" validate:"required,min=2"`
	Lastname    string `json:"lastname" validate:"required,min=2"`
	LineID      string `json:"line_id" validate:"required,min=1"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric,phone"`
}
