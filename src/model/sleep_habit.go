package model

type SleepHabit struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
