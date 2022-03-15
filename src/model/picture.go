package model

type Picture struct {
	PictureUrl string `gorm:"type:text" json:"picture_url"`
}
