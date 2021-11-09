package model

type PetPicture struct {
	ID         uint   `gorm:"primaryKey" json:"-"`
	PictureUrl string `gorm:"type:text" json:"picture_url"`
	StudentID  uint   `json:"-"`
}
