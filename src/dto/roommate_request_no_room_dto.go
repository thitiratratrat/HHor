package dto

type RoommateRequestNoRoomDTO struct {
	Budget int      `json:"budget" validate:"required,gt=0"`
	Zone   []string `json:"zones" validate:"required,dormzone"`
}
