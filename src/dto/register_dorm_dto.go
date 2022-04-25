package dto

type RegisterDormDTO struct {
	Name        string   `json:"name" validate:"required,min=1"`
	Type        string   `json:"type" validate:"required,oneof=male female mixed"`
	Rules       string   `json:"rules"`
	Long        float64  `json:"long" validate:"required,longitude"`
	Lat         float64  `json:"lat"  validate:"required,latitude"`
	Address     string   `json:"address" validate:"required,min=3"`
	Description string   `json:"description"`
	Zone        string   `json:"zone"  validate:"required,dormzone"`
	Facilities  []string `json:"facilities" validate:"required,min=1,dormfacilities"`
}
