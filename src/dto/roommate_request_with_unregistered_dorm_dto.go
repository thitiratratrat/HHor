package dto

type RoommateRequestWithUnregisteredDormDTO struct {
	StudentID         string   `json:"student_id" validate:"required"`
	DormName          string   `json:"dorm_name" validate:"required,min=2"`
	Zone              string   `json:"zone"  validate:"required,dormzone"`
	RoomDescription   string   `json:"room_description" validate:"required"`
	RoomPrice         int      `json:"room_price" validate:"required,gt=0"`
	RoomSize          float32  `json:"room_size" validate:"required,gt=0"`
	RoomFacilities    []string `json:"room_facilities" validate:"required"`
	NumberOfRoommates int      `json:"number_of_roommates" validate:"required,gt=0"`
	SharedRoomPrice   int      `json:"shared_room_price" validate:"required,gt=0"`
}
