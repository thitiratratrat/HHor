package model

type DormOwner struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Firstname string `gorm:"not null" json:"firstname"`
	Lastname  string `gorm:"not null" json:"lastname"`
	Email     string `gorm:"not null;unique" json:"email"`
	Password  string `gorm:"not null" json:"-"`
}
