package dto

type RoommateRequestFilterDTO struct {
	Zone     string `form:"zone" json:"zone,omitempty"`
	DormName string `form:"dorm_name" json:"dorm_name,omitempty"`
}
