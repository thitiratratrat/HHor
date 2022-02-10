package dto

import "mime/multipart"

type RoommateRequestWithUnregisteredDormDTO struct {
	StudentID         string                  `json:"student_id" form:"student_id" validate:"required"`
	Dorm              string                  `json:"dorm" form:"dorm" validate:"required"`
	Zone              string                  `json:"zone"`
	RoomDescription   string                  `gorm:"type:text" json:"room_description"`
	RoomSize          float32                 `json:"size"`
	RoomFacilities    []string                `gorm:"many2many:room_facility;" json:"facilities"`
	NumberOfRoommates int                     `json:"number_of_roommates"`
	SharedRoomPrice   int                     `json:"shared_room_price"`
	RoomPictures      *[]multipart.FileHeader `structs:",omitempty" form:"room_pictures,omitempty" swaggerignore:"true"`
}
