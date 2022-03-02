package dto

import "github.com/thitiratratrat/hhor/src/model"

type HabitDTO struct {
	PetHabit      []model.PetHabit      `json:"pet_habit"`
	SleepHabit    []model.SleepHabit    `json:"sleep_habit"`
	SmokeHabit    []model.SmokeHabit    `json:"smoke_habit"`
	StudyHabit    []model.StudyHabit    `json:"study_habit"`
	RoomCareHabit []model.RoomCareHabit `json:"room_care_habit"`
}
