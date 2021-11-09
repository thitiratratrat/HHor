package model

type SleepHabit struct {
	Name string `gorm:"primaryKey" json:"name"`
}
