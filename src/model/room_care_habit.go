package model

type RoomCareHabit struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
