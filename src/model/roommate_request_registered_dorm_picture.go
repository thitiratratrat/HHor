package model

type RoommateRequestRegisteredDormPicture struct {
	ID                                         uint   `gorm:"primaryKey" json:"-"`
	RoommateRequestWithRegisteredDormStudentID string `json:"-"`
	PictureUrl                                 string `gorm:"type:text" json:"picture_url"`
}
