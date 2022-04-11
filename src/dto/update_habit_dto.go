package dto

type UpdateHabitDTO struct {
	SmokeHabitID    string `json:"smoke_habit_id" validate:"required,number"`
	RoomCareHabitID string `json:"room_care_habit_id" validate:"required,number"`
	SleepHabitID    string `json:"sleep_habit_id" validate:"required,number"`
	StudyHabitID    string `json:"study_habit_id" validate:"required,number"`
	PetHabitID      string `json:"pet_habit_id" validate:"required,number"`
}
