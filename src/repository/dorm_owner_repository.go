package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type DormOwnerRepository interface {
	FindDormOwnerByEmail(email string) (model.DormOwner, error)
	FindDormOwnerByID(id string) (model.DormOwner, error)
	CreateDormOwner(model.DormOwner) (model.DormOwner, error)
	UpdateDormOwner(string, model.DormOwner) (model.DormOwner, error)
	DeleteBankAccount(string) (model.DormOwner, error)
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

func (repository *dormOwnerRepository) FindDormOwnerByID(id string) (model.DormOwner, error) {
	var dormOwner model.DormOwner

	err := repository.db.Preload("Dorms.Rooms.Pictures").Preload("Dorms.Rooms.Facilities").Preload("Dorms.Facilities").Preload("Dorms.Rooms", func(db *gorm.DB) *gorm.DB {
		return db.Order("rooms.price ASC")
	}).Preload("Dorms.Pictures").Preload("Dorms.DormZone").Preload("Dorms.NearbyLocations").Preload("Dorms").Where("id = ?", id).First(&dormOwner).Error

	return dormOwner, err
}

func (repository *dormOwnerRepository) CreateDormOwner(dormOwner model.DormOwner) (model.DormOwner, error) {
	err := repository.db.Create(&dormOwner).Error

	if err != nil {
		return model.DormOwner{}, err
	}

	return repository.FindDormOwnerByEmail(dormOwner.Email)
}

func (repository *dormOwnerRepository) UpdateDormOwner(id string, dormOwner model.DormOwner) (model.DormOwner, error) {
	err := repository.db.Model(&model.DormOwner{}).Where("id = ?", id).Updates(dormOwner).Error

	if err != nil {
		return model.DormOwner{}, err
	}

	return repository.FindDormOwnerByID(id)
}

func (repository *dormOwnerRepository) DeleteBankAccount(id string) (model.DormOwner, error) {
	err := repository.db.Model(&model.DormOwner{}).Where("id = ?", id).Updates(map[string]interface{}{"bank_qr_url": nil, "account_name": nil, "account_number": nil, "bank": nil}).Error

	if err != nil {
		return model.DormOwner{}, err
	}

	return repository.FindDormOwnerByID(id)
}
