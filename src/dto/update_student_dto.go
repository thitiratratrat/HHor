package dto

type UpdateStudentDTO struct {
	Firstname    string  `json:"firstname"  validate:"required,min=2"`
	Lastname     string  `json:"lastname" validate:"required,min=2"`
	Faculty      string  `json:"faculty"  validate:"required,faculty"`
	Biography    *string `json:"biography"`
	FacebookUrl  *string `json:"facebook_url" validate:"omitempty,url,contains=facebook"`
	TwitterUrl   *string `json:"twitter_url" validate:"omitempty,url,contains=twitter"`
	LinkedinUrl  *string `json:"linkedin_url" validate:"omitempty,url,contains=linkedin"`
	InstagramUrl *string `json:"instagram_url" validate:"omitempty,url,contains=instagram"`
}
