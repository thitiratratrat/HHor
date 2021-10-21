package model

type AllDormFacility struct {
	Name string `gorm:"primaryKey;not null" json:"name"`
}
