package model

type RoommateRequestWithUnregisteredDorm struct {
	StudentID         string                                   `gorm:"primaryKey" json:"student_id"`
	Student           Student                                  `json:"-"`
	DormName          string                                   `gorm:"type:citext" json:"dorm_name"`
	DormZoneName      string                                   `json:"zone"`
	DormZone          DormZone                                 `json:"-"`
	RoomDescription   string                                   `gorm:"type:text" json:"room_description"`
	RoomPrice         int                                      `gorm:"not null" json:"room_price"`
	RoomSize          float32                                  `json:"size"`
	RoomFacilities    []AllRoomFacility                        `gorm:"many2many:roommate_request_unregistered_dorm_room_facility;" json:"facilities"`
	NumberOfRoommates int                                      `json:"number_of_roommates"`
	SharedRoomPrice   int                                      `json:"shared_room_price"`
	RoomPictures      []RoommateRequestUnregisteredDormPicture `json:"room_pictures"`
	Longitude         float64                                  `gorm:"type:decimal(9,6);" json:"long"`
	Latitude          float64                                  `gorm:"type:decimal(8,6);" json:"lat"`
}
