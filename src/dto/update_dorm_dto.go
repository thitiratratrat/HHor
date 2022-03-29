package dto

type UpdateDormDTO struct {
	Name        string   `json:"name" validate:"omitempty,min=1"`
	Type        string   `json:"type" validate:"omitempty,oneof=male female mixed"`
	Rules       string   `json:"rules"`
	Long        float64  `json:"long" validate:"omitempty,longitude"`
	Lat         float64  `json:"lat"  validate:"omitempty,latitude"`
	Address     string   `json:"address" validate:"omitempty,min=3"`
	Description string   `json:"description"`
	Zone        string   `json:"zone"  validate:"omitempty,dormzone"`
	Facilities  []string `json:"facilities" validate:"omitempty,min=1,dormfacilities"`
	DormOwnerID string   `json:"owner_id" validate:"required,numeric"`
}
