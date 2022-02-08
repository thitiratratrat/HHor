package model

type RoommateRequestWithNoRoom struct {
	StudentID string     `gorm:"primaryKey"`
	Student   Student    `json:"student"`
	Budget    int        `json:"budget"`
	Zones     []DormZone `gorm:"many2many:roommate_request_no_room_zone;" json:"zones"`
}
