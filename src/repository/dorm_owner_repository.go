package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type DormOwnerRepository interface {
	GetDormOwner(email string) (model.DormOwner, error)
}

func DormOwnerRepositoryHandler(db *gorm.DB) DormOwnerRepository {
	return &dormOwnerRepository{
		db: db,
	}
}

type dormOwnerRepository struct {
	db *gorm.DB
}

func (repository *dormOwnerRepository) GetDormOwner(email string) (model.DormOwner, error) {
	var dormOwner model.DormOwner

	err := repository.db.Where("email = ?", email).First(&dormOwner).Error

	return dormOwner, err
}
