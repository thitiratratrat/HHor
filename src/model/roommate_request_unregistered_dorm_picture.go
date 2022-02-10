package model

type RoommateRequestUnregisteredDormPicture struct {
	ID                                           uint   `gorm:"primaryKey" json:"-"`
	RoommateRequestWithUnregisteredDormStudentID string `json:"-"`
	PictureUrl                                   string `gorm:"type:text" json:"picture_url"`
}
