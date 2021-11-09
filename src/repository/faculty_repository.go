package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type FacultyRepository interface {
	GetFaculties() []string
}

func FacultyRepositoryHandler(db *gorm.DB) FacultyRepository {
	return &facultyRepository{
		db: db,
	}
}

type facultyRepository struct {
	db *gorm.DB
}

func (repository *facultyRepository) GetFaculties() []string {
	var faculties []string

	repository.db.Model(&model.Faculty{}).Pluck("name", &faculties)

	return faculties
}
