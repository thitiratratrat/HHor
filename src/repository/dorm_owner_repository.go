package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type DormOwnerRepository interface {
	FindDormOwnerByEmail(email string) (model.DormOwner, error)
	CreateDormOwner(model.DormOwner) (model.DormOwner, error)
}

func DormOwnerRepositoryHandler(db *gorm.DB) DormOwnerRepository {
	return &dormOwnerRepository{
		db: db,
	}
}

type dormOwnerRepository struct {
	db *gorm.DB
}

func (repository *dormOwnerRepository) FindDormOwnerByEmail(email string) (model.DormOwner, error) {
	var dormOwner model.DormOwner

	err := repository.db.Where("email = ?", email).First(&dormOwner).Error

	return dormOwner, err
}

func (repository *dormOwnerRepository) CreateDormOwner(dormOwner model.DormOwner) (model.DormOwner, error) {
	err := repository.db.Create(&dormOwner).Error

	if err != nil {
		return model.DormOwner{}, err
	}

	return repository.FindDormOwnerByEmail(dormOwner.Email)
}
