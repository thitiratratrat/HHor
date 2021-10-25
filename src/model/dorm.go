package model

type dormType string

const (
	Mixed  dormType = "mixed"
	Female dormType = "female"
	Male   dormType = "male"
)

type Dorm struct {
	ID              uint              `gorm:"primaryKey" json:"id"`
	Name            string            `gorm:"not null;type:citext" json:"name"`
	Type            string            `gorm:"not null" json:"type" sql:"dorm_type"`
	Rules           string            `gorm:"type:text;" json:"rules"`
	Longitude       float64           `gorm:"type:decimal(9,6);not null" json:"longitude"`
	Latitude        float64           `gorm:"type:decimal(8,6);not null" json:"latitude"`
	Address         string            `gorm:"not null" json:"address"`
	DormZoneName    string            `json:"zone"`
	DormZone        DormZone          `json:"-"`
	AccountID       int               `gorm:"column:owner" json:"-"`
	Account         Account           `json:"account"`
	Facilities      []AllDormFacility `gorm:"many2many:dorm_facility;" json:"facilities"`
	Pictures        []DormPicture     `gorm:"foreignKey:DormID" json:"pictures"`
	NearbyLocations []NearbyLocation  `gorm:"many2many:nearby_locations;" json:"nearby_locations"`
	Rooms           []Room            `json:"rooms"`
}
