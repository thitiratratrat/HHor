package model

type SmokeHabit struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
