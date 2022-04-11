package dto

type UpdatePreferenceDTO struct {
	PreferredSmokeHabitID    *string `json:"preferred_smoke_habit_id" validate:"omitempty,number"`
	PreferredRoomCareHabitID *string `json:"preferred_room_care_habit_id" validate:"omitempty,number"`
	PreferredSleepHabitID    *string `json:"preferred_sleep_habit_id" validate:"omitempty,number"`
	PreferredStudyHabitID    *string `json:"preferred_study_habit_id" validate:"omitempty,number"`
	PreferredPetHabitID      *string `json:"preferred_pet_habit_id"  validate:"omitempty,number"`
	PreferredGenderName      *string `json:"preferred_gender" validate:"omitempty,oneof=male female lgbtq+"`
	OtherPreference          *string `json:"other_preference"`
}
