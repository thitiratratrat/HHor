package model

type AllRoomFacility struct {
	Name string `gorm:"primaryKey;not null" json:"name"`
}