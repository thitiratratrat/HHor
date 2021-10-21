package model

type DormPicture struct {
	ID         uint   `gorm:"primaryKey" json:"-"`
	PictureUrl string `gorm:"type:text" json:"picture_url"`
	DormID     uint   `json:"-"`
}
