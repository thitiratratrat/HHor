package model

type StudyHabit struct {
	Name string `gorm:"primaryKey" json:"name"`
}
