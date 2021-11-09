package model

type RoomCareHabit struct {
	Name string `gorm:"primaryKey" json:"name"`
}
