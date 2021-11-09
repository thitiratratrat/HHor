package model

type Faculty struct {
	Name string `gorm:"primaryKey" json:"name"`
}
