package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type DormFacilityRepository interface {
	FindAllDormFacilities() []string
}

func DormFacilityRepositoryHandler(db *gorm.DB) DormFacilityRepository {
	return &dormFacility{
		db: db,
	}
}

type dormFacility struct {
	db *gorm.DB
}

func (repository *dormFacility) FindAllDormFacilities() []string {
	var facilities []string

	repository.db.Model(&model.AllDormFacility{}).Pluck("name", &facilities)

	return facilities
}
