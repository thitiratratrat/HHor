package model

type Student struct {
	ID                 string             `gorm:"primaryKey;check:id ~ '^\d*$'" json:"id"`
	Email              string             `gorm:"not null;unique" json:"email"`
	Firstname          string             `gorm:"not null" json:"firstname"`
	Lastname           string             `gorm:"not null" json:"lastname"`
	Password           string             `gorm:"not null" json:"-"`
	EnrollmentYear     int                `gorm:"not null;check:enrollment_year > 2014" json:"enrollment_year"`
	GenderName         string             `gorm:"not null" json:"gender"`
	Gender             Gender             `json:"-"`
	FacultyName        string             `gorm:"not null" json:"faculty"`
	Faculty            Faculty            `json:"-"`
	PictureUrl         string             `gorm:"type:text" json:"picture_url"`
	SocialMedia        SocialMedia        `gorm:"embedded" json:"social_media"`
	Biography          string             `gorm:"type:text" json:"biography"`
	PetPictures        []PetPicture       `gorm:"foreignKey:StudentID" json:"pet_pictures"`
	PersonalHabit      PersonalHabit      `gorm:"embedded" json:"personal_habit"`
	RoommatePreference RoommatePreference `gorm:"embedded" json:"roommate_preference"`
	RoommateRequest    *string            `json:"-"`
}

type SocialMedia struct {
	FacebookUrl  string `gorm:"type:text" json:"facebook_url"`
	TwitterUrl   string `gorm:"type:text" json:"twitter_url"`
	LinkedinUrl  string `gorm:"type:text" json:"linkedin_url"`
	InstagramUrl string `gorm:"type:text" json:"instagram_url"`
}

type PersonalHabit struct {
	StudyHabitID    *string       `gorm:"default:null" json:"-"`
	StudyHabit      StudyHabit    `json:"study_habit"`
	SleepHabitID    *string       `gorm:"default:null" json:"-"`
	SleepHabit      SleepHabit    `json:"sleep_habit"`
	SmokeHabitID    *string       `gorm:"default:null" json:"-"`
	SmokeHabit      SmokeHabit    `json:"smoke_habit"`
	RoomCareHabitID *string       `gorm:"default:null" json:"-"`
	RoomCareHabit   RoomCareHabit `json:"room_care_habit"`
	PetHabitID      *string       `gorm:"default:null" json:"-"`
	PetHabit        PetHabit      `json:"pet_habit"`
}

type RoommatePreference struct {
	PreferredGenderName      *string       `gorm:"default:null" json:"preferred_gender"`
	PreferredGender          Gender        `json:"-"`
	PreferredStudyHabitID    *string       `json:"-"`
	PreferredStudyHabit      StudyHabit    `json:"preferred_study_habit"`
	PreferredSleepHabitID    *string       `json:"-"`
	PreferredSleepHabit      SleepHabit    `json:"preferred_sleep_habit"`
	PreferredSmokeHabitID    *string       `json:"-"`
	PreferredSmokeHabit      SmokeHabit    `json:"preferred_smoke_habit"`
	PreferredRoomCareHabitID *string       `json:"-"`
	PreferredRoomCareHabit   RoomCareHabit `json:"preferred_room_care_habit"`
	PreferredPetHabitID      *string       `json:"-"`
	PreferredPetHabit        PetHabit      `json:"preferred_pet_habit"`
	OtherPreference          *string       `json:"other_preference"`
}
