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
	PictureUrl         *string            `gorm:"default:null;type:text" json:"picture_url"`
	SocialMedia        SocialMedia        `gorm:"embedded" json:"social_media"`
	Biography          *string            `gorm:"default:null;type:text" json:"biography"`
	PetPictures        []PetPicture       `gorm:"foreignKey:StudentID" json:"pet_pictures"`
	PersonalHabit      PersonalHabit      `gorm:"embedded" json:"personal_habit"`
	RoommatePreference RoommatePreference `gorm:"embedded" json:"roommate_preference"`
	RoommateRequest    *string            `json:"-"`
	VerificationCode   *string            `json:"-"`
	HasVerified        bool               `gorm:"default:false;" json:"-"`
}

type SocialMedia struct {
	FacebookUrl  *string `gorm:"default:null;type:text" json:"facebook_url"`
	TwitterUrl   *string `gorm:"default:null;type:text" json:"twitter_url"`
	LinkedinUrl  *string `gorm:"default:null;type:text" json:"linkedin_url"`
	InstagramUrl *string `gorm:"default:null;type:text" json:"instagram_url"`
}

type PersonalHabit struct {
	StudyHabitID    *uint          `gorm:"default:null" json:"-"`
	StudyHabit      *StudyHabit    `json:"study_habit"`
	SleepHabitID    *uint          `gorm:"default:null" json:"-"`
	SleepHabit      *SleepHabit    `json:"sleep_habit"`
	SmokeHabitID    *uint          `gorm:"default:null" json:"-"`
	SmokeHabit      *SmokeHabit    `json:"smoke_habit"`
	RoomCareHabitID *uint          `gorm:"default:null" json:"-"`
	RoomCareHabit   *RoomCareHabit `json:"room_care_habit"`
	PetHabitID      *uint          `gorm:"default:null" json:"-"`
	PetHabit        *PetHabit      `json:"pet_habit"`
}

type RoommatePreference struct {
	PreferredGenderName      *string        `gorm:"default:null" json:"preferred_gender"`
	PreferredGender          Gender         `json:"-"`
	PreferredStudyHabitID    *uint          `gorm:"default:null" json:"-"`
	PreferredStudyHabit      *StudyHabit    `json:"preferred_study_habit"`
	PreferredSleepHabitID    *uint          `gorm:"default:null" json:"-"`
	PreferredSleepHabit      *SleepHabit    `json:"preferred_sleep_habit"`
	PreferredSmokeHabitID    *uint          `gorm:"default:null" json:"-"`
	PreferredSmokeHabit      *SmokeHabit    `json:"preferred_smoke_habit"`
	PreferredRoomCareHabitID *uint          `gorm:"default:null" json:"-"`
	PreferredRoomCareHabit   *RoomCareHabit `json:"preferred_room_care_habit"`
	PreferredPetHabitID      *uint          `gorm:"default:null" json:"-"`
	PreferredPetHabit        *PetHabit      `json:"preferred_pet_habit"`
	OtherPreference          *string        `gorm:"default:null" json:"other_preference"`
}
