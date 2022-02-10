package model

type RoommateRequestWithNoRoom struct {
	StudentID string     `gorm:"primaryKey" json:"student_id"`
	Student   Student    `json:"-"`
	Budget    int        `json:"budget"`
	Zones     []DormZone `gorm:"many2many:roommate_request_no_room_zone;" json:"zones"`
}
