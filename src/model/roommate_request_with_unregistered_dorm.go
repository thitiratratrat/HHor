package model

type RoommateRequestWithUnregisteredDorm struct {
	StudentID         string                                   `gorm:"primaryKey" json:"student_id"`
	Student           Student                                  `json:"-"`
	Dorm              string                                   `json:"dorm"`
	Zone              DormZone                                 `json:"zone"`
	RoomDescription   string                                   `gorm:"type:text" json:"room_description"`
	RoomSize          float32                                  `json:"size"`
	RoomFacilities    []AllRoomFacility                        `gorm:"many2many:room_facility;" json:"facilities"`
	NumberOfRoommates int                                      `json:"number_of_roommates"`
	SharedRoomPrice   int                                      `json:"shared_room_price"`
	RoomPictures      []RoommateRequestUnregisteredDormPicture `json:"room_pictures"`
}
