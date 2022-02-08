package model

type PetHabit struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
