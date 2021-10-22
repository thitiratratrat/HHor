package model

type RoomPicture struct {
	ID         uint   `gorm:"primaryKey" json:"-"`
	PictureUrl string `gorm:"type:text" json:"picture_url"`
	RoomID     uint   `json:"-"`
}
