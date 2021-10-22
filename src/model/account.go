package model

type Account struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Firstname string `gorm:"not null" json:"firstname"`
	Lastname  string `gorm:"not null" json:"lastname"`
}
