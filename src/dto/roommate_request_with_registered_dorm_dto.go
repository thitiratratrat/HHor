package dto

type RoommateRequestWithRegisteredDormDTO struct {
	StudentID         string `json:"student_id" validate:"required"`
	DormID            string `json:"dorm_id"  validate:"required,number"`
	RoomID            string `json:"room_id"  validate:"required,number"`
	NumberOfRoommates int    `json:"number_of_roommates"  validate:"required,gt=0"`
	SharedRoomPrice   int    `json:"shared_room_price"  validate:"required,gt=0"`
}
