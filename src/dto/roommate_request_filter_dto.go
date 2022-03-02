package dto

type RoommateRequestFilterDTO struct {
	Zone              *string  `form:"zone" json:"zone,omitempty" validate:"omitempty,dormzone"`
	DormName          *string  `form:"dorm_name" json:"dorm_name,omitempty" validate:"omitempty,min=1"`
	Gender            []string `form:"gender" json:"gender,omitempty"  validate:"omitempty,dive,oneof=male female lgbtq+" swaggerignore:"true"`
	Faculty           []string `form:"faculties" json:"faculties,omitempty"  validate:"omitempty,faculty" swaggerignore:"true"`
	YearOfStudy       []int    `form:"year_of_study" json:"year_of_study,omitempty" validate:"omitempty,dive,min=1,max=4" swaggerignore:"true"`
	LowerPrice        *int     `form:"lower_price" json:"lower_price,omitempty" validate:"required_with=UpperPrice,omitempty,gt=0"`
	UpperPrice        *int     `form:"upper_price" json:"upper_price,omitempty" validate:"required_with=LowerPrice,gtefield=LowerPrice"`
	NumberOfRoommates []int    `form:"number_of_roommates" json:"number_of_roommates,omitempty" validate:"omitempty,dive,min=0" swaggerignore:"true"`
	RoomFacilities    []string `form:"room_facilities" json:"room_facilities,omitempty" validate:"omitempty,roomfacility" swaggerignore:"true"`
	Preferences
}

type Preferences struct {
	SmokeHabitID    *string `form:"smoke_habit_id" json:"smoke_habit_id,omitempty"  validate:"omitempty,number"`
	RoomCareHabitID *string `form:"room_care_habit_id" json:"room_care_habit_id,omitempty"  validate:"omitempty,number"`
	SleepHabitID    *string `form:"sleep_habit_id" json:"sleep_habit_id,omitempty" validate:"omitempty,number"`
	StudyHabitID    *string `form:"study_habit_id" json:"study_habit_id,omitempty"  validate:"omitempty,number"`
	PetHabitID      *string `form:"pet_habit_id" json:"pet_habit_id,omitempty" validate:"omitempty,number"`
}
