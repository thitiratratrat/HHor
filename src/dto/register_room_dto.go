package dto

type RegisterRoomDTO struct {
	Name          string   `json:"name" validate:"required,min=2"`
	Price         int      `json:"price" validate:"required,min=100"`
	Size          float32  `json:"size" validate:"required,min=10"`
	Description   string   `json:"description"`
	Capacity      int      `json:"capacity" validate:"required,min=1"`
	AvailableFrom *string  `json:"available_from" validate:"required,datetime=2006-01-02"`
	DormID        string   `json:"dorm_id" validate:"required,numeric"`
	Facilities    []string `json:"facilities" validate:"required,roomfacility"`
	DormOwnerID   string   `json:"dorm_owner_id" validate:"required,numeric"`
}
