package dto

type UpdateBankAccountDTO struct {
	Bank          string `json:"bank" validate:"required,min=3"`
	AccountName   string `json:"account_name" validate:"required,min=3"`
	AccountNumber string `json:"account_number" validate:"required,numeric,min=4"`
}
