package dto

type DormFilterDTO struct {
	Type           []string `form:"type" json:"type,omitempty"`
	Capacity       *int     `form:"capacity" json:"capacity,omitempty" validate:"omitempty,gte=1"`
	DormFacilities []string `form:"dorm_facilities" json:"dorm_facilities,omitempty" validate:"omitempty,dormfacilities"`
	RoomFacilities []string `form:"room_facilities" json:"room_facilities,omitempty" validate:"omitempty,roomfacilities"`
	Zone           *string  `form:"zone" json:"zone,omitempty" validate:"omitempty,dormzone"`
	Name           *string  `form:"name" json:"name,omitempty" validate:"omitempty,min=1"`
	Lat            *float64 `form:"lat" json:"lat,omitempty" validate:"required_with=Long,omitempty,latitude"`
	Long           *float64 `form:"long" json:"long,omitempty"  validate:"required_with=Lat,omitempty,longitude"`
	LowerPrice     *int     `form:"lower_price" json:"lower_price,omitempty" validate:"required_with=UpperPrice,omitempty,gt=0"`
	UpperPrice     *int     `form:"upper_price" json:"upper_price,omitempty" validate:"required_with=LowerPrice,gtefield=LowerPrice"`
}
