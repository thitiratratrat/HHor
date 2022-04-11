package dto

import "github.com/thitiratratrat/hhor/src/customtype"

//TODO: add update last name, first name
type StudentUpdateDTO struct {
	Biography                customtype.JSONString `json:"biography"`
	EnrollmentYear           customtype.JSONInt    `json:"enrollment_year" validate:"required,gte=2014"`
	FacebookUrl              customtype.JSONString `json:"facebook_url" validate:"omitempty,url,contains=facebook"`
	TwitterUrl               customtype.JSONString `json:"twitter_url" validate:"omitempty,url,contains=twitter"`
	LinkedinUrl              customtype.JSONString `json:"linkedin_url" validate:"omitempty,url,contains=linkedin"`
	InstagramUrl             customtype.JSONString `json:"instagram_url" validate:"omitempty,url,contains=instagram"`
	SmokeHabitID             customtype.JSONString `json:"smoke_habit_id"`
	RoomCareHabitID          customtype.JSONString `json:"room_care_habit_id"`
	SleepHabitID             customtype.JSONString `json:"sleep_habit_id"`
	StudyHabitID             customtype.JSONString `json:"study_habit_id"`
	PetHabitID               customtype.JSONString `json:"pet_habit_id" validate:"omitempty,number"`
	PreferredSmokeHabitID    customtype.JSONString `json:"preferred_smoke_habit_id" validate:"omitempty,number"`
	PreferredRoomCareHabitID customtype.JSONString `json:"preferred_room_care_habit_id" validate:"omitempty,number"`
	PreferredSleepHabitID    customtype.JSONString `json:"preferred_sleep_habit_id" validate:"omitempty,number"`
	PreferredStudyHabitID    customtype.JSONString `json:"preferred_study_habit_id" validate:"omitempty,number"`
	PreferredPetHabitID      customtype.JSONString `json:"preferred_pet_habit_id"  validate:"omitempty,number"`
	PreferredGenderName      customtype.JSONString `json:"preferred_gender" validate:"omitempty"`
	OtherPreference          customtype.JSONString `json:"other_preference"`
}

type StudentUpdateSwagDTO struct {
	Biography                *string `json:"biography"`
	EnrollmentYear           *int    `json:"enrollment_year" validate:"required,gte=2014"`
	FacebookUrl              *string `json:"facebook_url" validate:"omitempty,url,contains=facebook"`
	TwitterUrl               *string `json:"twitter_url" validate:"omitempty,url,contains=twitter"`
	LinkedinUrl              *string `json:"linkedin_url" validate:"omitempty,url,contains=linkedin"`
	InstagramUrl             *string `json:"instagram_url" validate:"omitempty,url,contains=instagram"`
	SmokeHabitID             *string `json:"smoke_habit_id"`
	RoomCareHabitID          *string `json:"room_care_habit_id"`
	SleepHabitID             *string `json:"sleep_habit_id"`
	StudyHabitID             *string `json:"study_habit_id"`
	PetHabitID               *string `json:"pet_habit_id" validate:"omitempty,number"`
	PreferredSmokeHabitID    *string `json:"preferred_smoke_habit_id" validate:"omitempty,number"`
	PreferredRoomCareHabitID *string `json:"preferred_room_care_habit_id" validate:"omitempty,number"`
	PreferredSleepHabitID    *string `json:"preferred_sleep_habit_id" validate:"omitempty,number"`
	PreferredStudyHabitID    *string `json:"preferred_study_habit_id" validate:"omitempty,number"`
	PreferredPetHabitID      *string `json:"preferred_pet_habit_id"  validate:"omitempty,number"`
	PreferredGenderName      *string `json:"preferred_gender" validate:"omitempty"`
	OtherPreference          *string `json:"other_preference"`
}
