package dto

type StudentUpdateDTO struct {
	Biography          *string            `structs:",omitempty" json:"biography,omitempty"`
	EnrollmentYear     int                `structs:",omitempty" json:"enrollment_year" validate:"omitempty,gte=2014"`
	SocialMedia        SocialMedia        `structs:",omitempty" json:"social_media"`
	PersonalHabit      PersonalHabit      `structs:",omitempty" json:"personal_habit"`
	RoommatePreference RoommatePreference `structs:",omitempty" json:"roommate_preference"`
}

type SocialMedia struct {
	FacebookUrl  *string `structs:",omitempty" json:"facebook_url,omitempty" validate:"omitempty,url,contains=facebook"`
	TwitterUrl   *string `structs:",omitempty" json:"twitter_url,omitempty" validate:"omitempty,url,contains=twitter"`
	LinkedinUrl  *string `structs:",omitempty" json:"linkedin_url,omitempty" validate:"omitempty,url,contains=linkedin"`
	InstagramUrl *string `structs:",omitempty" json:"instagram_url,omitempty" validate:"omitempty,url,contains=instagram"`
}

type PersonalHabit struct {
	SmokeHabitID    *string `structs:",omitempty" json:"smoke_habit_id,omitempty"  validate:"omitempty,number"`
	RoomCareHabitID *string `structs:",omitempty" json:"room_care_habit_id,omitempty"  validate:"omitempty,number"`
	SleepHabitID    *string `structs:",omitempty" json:"sleep_habit_id,omitempty" validate:"omitempty,number"`
	StudyHabitID    *string `structs:",omitempty" json:"study_habit_id,omitempty"  validate:"omitempty,number"`
	PetHabitID      *string `structs:",omitempty" json:"pet_habit_id,omitempty" validate:"omitempty,number"`
}

type RoommatePreference struct {
	PreferredSmokeHabitID    *string `structs:",omitempty" json:"preferred_smoke_habit_id,omitempty" validate:"omitempty,number"`
	PreferredRoomCareHabitID *string `structs:",omitempty" json:"preferred_room_care_habit_id,omitempty" validate:"omitempty,number"`
	PreferredSleepHabitID    *string `structs:",omitempty" json:"preferred_sleep_habit_id,omitempty" validate:"omitempty,number"`
	PreferredStudyHabitID    *string `structs:",omitempty" json:"preferred_study_habit_id,omitempty" validate:"omitempty,number"`
	PreferredPetHabitID      *string `structs:",omitempty" json:"preferred_pet_habit_id,omitempty"  validate:"omitempty,number"`
	PreferredGenderName      *string `structs:",omitempty" json:"preferred_gender,omitempty" validate:"omitempty"`
	OtherPreference          *string `structs:",omitempty" json:"other_preference,omitempty"`
}
