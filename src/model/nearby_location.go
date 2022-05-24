package model

type NearbyLocation struct {
	DormID    int     `gorm:"primaryKey" json:"-"`
	Name      string  `gorm:"primaryKey;not null" json:"name"`
	Longitude float64 `gorm:"type:decimal(9,6);not null"  json:"long"`
	Latitude  float64 `gorm:"type:decimal(8,6);not null" json:"lat"`
	Distance  float64 `gorm:"not null" json:"distance"`
}
