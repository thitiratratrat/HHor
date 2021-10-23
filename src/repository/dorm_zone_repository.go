package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type DormZoneRepository interface {
	FindDormZones() []string
}

func DormZoneRepositoryHandler(db *gorm.DB) DormZoneRepository {
	return &dormZoneRepository{
		db: db,
	}
}

type dormZoneRepository struct {
	db *gorm.DB
}

func (repository *dormZoneRepository) FindDormZones() []string {
	var zones []string

	repository.db.Model(&model.DormZone{}).Pluck("name", &zones)

	return zones
}
