package model

type PetHabit struct {
	Name string `gorm:"primaryKey" json:"name"`
}
