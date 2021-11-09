package model

type SmokeHabit struct {
	Name string `gorm:"primaryKey" json:"name"`
}
