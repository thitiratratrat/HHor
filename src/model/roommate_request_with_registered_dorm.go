package model

type RoommateRequestWithRegisteredDorm struct {
	StudentID         string                                 `gorm:"primaryKey" json:"-"`
	Student           Student                                `json:"-"`
	SharedRoomPrice   int                                    `json:"shared_room_price"`
	NumberOfRoommates int                                    `json:"number_of_roommates"`
	RoomID            uint                                   `json:"room_id"`
	Room              Room                                   `json:"-"`
	DormID            uint                                   `json:"dorm_id"`
	Dorm              Dorm                                   `json:"-"`
	RoomPictures      []RoommateRequestRegisteredDormPicture `json:"room_pictures"`
}
