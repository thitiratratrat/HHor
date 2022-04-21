package dto

import "github.com/thitiratratrat/hhor/src/model"

type DormDetailDTO struct {
	model.Dorm
	DormOwner DormOwnerDTO `json:"dorm_owner"`
}

type DormOwnerDTO struct {
	ID         uint    `json:"id"`
	Firstname  string  `json:"firstname"`
	Lastname   string  `json:"lastname"`
	Email      string  `json:"email"`
	PictureUrl *string `json:"picture_url"`
}
