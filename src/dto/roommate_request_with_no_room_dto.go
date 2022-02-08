package dto

type RoommateRequestWithNoRoomDTO struct {
	StudentID string   `json:"student_id" validate:"required"`
	Budget    int      `json:"budget" validate:"required,gt=0"`
	Zone      []string `json:"zones" validate:"required,dormzone"`
}
