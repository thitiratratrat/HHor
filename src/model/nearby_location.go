package model

type NearbyLocation struct {
	LocationID int      `gorm:"primaryKey" json:"-"`
	DormID     int      `gorm:"primaryKey" json:"-"`
	Location   Location `json:"location"`
	Dorm       Dorm     `json:"-"`
	Distance   float64  `gorm:"not null" json:"distance"`
}
