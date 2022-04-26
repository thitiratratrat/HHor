package model

type DormOwner struct {
	ID               uint         `gorm:"primaryKey" json:"id"`
	Firstname        string       `gorm:"not null" json:"firstname"`
	Lastname         string       `gorm:"not null" json:"lastname"`
	Email            string       `gorm:"not null;unique" json:"email"`
	Password         string       `gorm:"not null" json:"-"`
	LineID           string       `json:"line_id"`
	PhoneNumber      string       `json:"phone_number"`
	Dorms            []Dorm       `json:"dorms"`
	PictureUrl       *string      `gorm:"default:null;type:text" json:"picture_url"`
	BankAccount      *BankAccount `gorm:"embedded" json:"bank_account"`
	VerificationCode *string      `json:"-"`
	HasVerified      bool         `gorm:"default:false;" json:"-"`
}

type BankAccount struct {
	Bank          *string `gorm:"default:null" json:"bank"`
	AccountName   *string `gorm:"default:null" json:"account_name"`
	AccountNumber *string `gorm:"default:null" json:"account_number"`
	BankQrUrl     *string `gorm:"default:null;type:text" json:"bank_qr_url"`
}
