package dto

import "mime/multipart"

type DormOwnerPictureDTO struct {
	ProfilePicture *multipart.FileHeader `structs:",omitempty" form:"profile_picture,omitempty" swaggerignore:"true"`
	BankQR         *multipart.FileHeader `structs:",omitempty" form:"bank_qr,omitempty" swaggerignore:"true"`
}
