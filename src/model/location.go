package model

type Location struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"not null" json:"name"`
	Longitude float64 `gorm:"type:decimal(9,6);not null"  json:"long"`
	Latitude  float64 `gorm:"type:decimal(8,6);not null" json:"lat"`
}
