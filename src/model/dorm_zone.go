package model

type DormZone struct {
	Name string `gorm:"primaryKey;not null;type:citext" json:"name"`
}
