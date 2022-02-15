package model

//TODO: availability: available boolean, available from ___
type Room struct {
	ID          uint              `gorm:"primaryKey" json:"id"`
	Name        string            `gorm:"not null" json:"name"`
	Price       int               `gorm:"not null" json:"price"`
	Size        float32           `json:"size"`
	Description string            `gorm:"type:text;not null" json:"description"`
	Capacity    int               `gorm:"not null" json:"capacity"`
	Available   bool              `gorm:"not null" json:"available"` //TODO: change to available from? keep as date
	DormID      uint              `json:"-"`
	Pictures    []RoomPicture     `json:"pictures"`
	Facilities  []AllRoomFacility `gorm:"many2many:room_facility;" json:"facilities"`
}
