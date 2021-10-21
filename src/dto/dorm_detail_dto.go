package dto

type DormDetailDTO struct {
	ID              uint
	Name            string
	Type            string
	Rules           string
	Longitude       float64
	Latitude        float64
	Address         string
	Zone            string `json:"zone"`
	Owner           Owner
	Facilities      []string
	Pictures        []string
	NearbyLocations []NearbyLocation
	Rooms           []Room
}

type Owner struct {
	Firstname string
	Lastname  string
}

type NearbyLocation struct {
	Name     string
	Distance string
}

type Room struct {
	Name        string
	Price       int
	Size        float32
	Description string
	Capacity    int
	Pictures    []string
	Facilities  []string
}
