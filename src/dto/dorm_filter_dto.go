package dto

type DormFilterDTO struct {
	Type           []string `form:"type" json:"type,omitempty"`
	Capacity       int      `form:"capacity" json:"capacity,omitempty"`
	DormFacilities []string `form:"dorm_facilities" json:"dorm_facilities,omitempty"`
	RoomFacilities []string `form:"room_facilities" json:"room_facilities,omitempty"`
	Zone           string   `form:"zone" json:"zone,omitempty"`
	Name           string   `form:"name" json:"name,omitempty"`
	Lat            float64  `form:"lat" json:"lat,omitempty"`
	Long           float64  `form:"long" json:"long,omitempty"`
	LowerPrice     int      `form:"lower_price" json:"lower_price,omitempty"`
	UpperPrice     int      `form:"upper_price" json:"upper_price,omitempty"`
}
