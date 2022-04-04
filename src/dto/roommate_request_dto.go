package dto

import (
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/model"
)

type RoommateRequestDTO struct {
	Type                constant.RoommateRequestType `json:"type"`
	Student             model.Student                `json:"student"`
	*Room               `json:"room"`
	*Dorm               `json:"dorm"`
	Budget              *int              `json:"budget"`
	InterestedDormZones *[]model.DormZone `json:"interested_dorm_zones"`
}

type Room struct {
	ID                *uint                   `json:"id"`
	Name              *string                 `json:"name"`
	Price             int                     `json:"price"`
	Size              float32                 `json:"size"`
	Description       string                  `json:"description"`
	Pictures          []model.Picture         `json:"pictures"`
	Facilities        []model.AllRoomFacility `json:"facilities"`
	SharedRoomPrice   int                     `json:"shared_room_price"`
	NumberOfRoommates int                     `json:"number_of_roommates"`
}

type Dorm struct {
	ID      *uint   `json:"id"`
	Name    string  `json:"name"`
	Zone    string  `json:"zone"`
	Picture *string `json:"picture"`
}
