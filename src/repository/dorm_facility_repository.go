package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type DormFacilityRepository interface {
	FindAllDormFacilities() []string
}

func DormFacilityRepositoryHandler(db *gorm.DB) DormFacilityRepository {
	return &dormFacilityRepository{
		db: db,
	}
}

type dormFacilityRepository struct {
	db *gorm.DB
}

func (repository *dormFacilityRepository) FindAllDormFacilities() []string {
	var facilities []string

	repository.db.Model(&model.AllDormFacility{}).Pluck("name", &facilities)

	return facilities
}
