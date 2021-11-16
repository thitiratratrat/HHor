package model

type Student struct {
	Email               string        `gorm:"primaryKey;unique" json:"email"`
	StudentID           string        `gorm:"not null;unique;check:student_id ~ '^\d*$'" json:"student_id"`
	Firstname           string        `gorm:"not null" json:"firstname"`
	Lastname            string        `gorm:"not null" json:"lastname"`
	Password            string        `gorm:"not null" json:"-"`
	YearOfStudy         int           `gorm:"not null;check:year_of_study > 0 and year_of_study < 9" json:"year_of_study"`
	GenderName          string        `gorm:"not null" json:"gender"`
	Gender              Gender        `json:"-"`
	FacultyName         string        `gorm:"not null" json:"faculty"`
	Faculty             Faculty       `json:"-"`
	PictureUrl          string        `gorm:"type:text" json:"picture_url"`
	FacebookUrl         string        `gorm:"type:text" json:"facebook_url"`
	TwitterUrl          string        `gorm:"type:text" json:"twitter_url"`
	LinkedinUrl         string        `gorm:"type:text" json:"linkedin_url"`
	InstagramUrl        string        `gorm:"type:text" json:"instagram_url"`
	Biography           string        `gorm:"type:text"`
	PetPictures         []PetPicture  `gorm:"foreignKey:StudentID" json:"pet_pictures"`
	PreferredGenderName *string       `gorm:"default:null" json:"preferred_gender"`
	PreferredGender     Gender        `json:"-"`
	StudyHabitName      *string       `gorm:"default:null" json:"study_habit"`
	StudyHabit          StudyHabit    `json:"-"`
	SleepHabitName      *string       `gorm:"default:null" json:"sleeping_habit"`
	SleepHabit          SleepHabit    `json:"-"`
	SmokeHabitName      *string       `gorm:"default:null" json:"smoke_habit"`
	SmokeHabit          SmokeHabit    `json:"-"`
	RoomCareHabitName   *string       `gorm:"default:null" json:"room_care_habit"`
	RoomCareHabit       RoomCareHabit `json:"-"`
	PetHabitName        *string       `gorm:"default:null" json:"pet_habit"`
	PetHabit            PetHabit      `json:"-"`
}
