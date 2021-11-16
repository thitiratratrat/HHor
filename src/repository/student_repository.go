package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(model.Student) error
	GetStudent(email string) (model.Student, error)
	UpdateStudent(email string, studentUpdate map[string]interface{}) error
}

func StudentRepositoryHandler(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

type studentRepository struct {
	db *gorm.DB
}

func (repository *studentRepository) CreateStudent(student model.Student) error {
	return repository.db.Create(&student).Error
}

func (repository *studentRepository) GetStudent(email string) (model.Student, error) {
	var student model.Student

	err := repository.db.Where("email = ?", email).First(&student).Error

	return student, err
}

func (repository *studentRepository) UpdateStudent(email string, studentUpdate map[string]interface{}) error {
	err := repository.db.Model(&model.Student{}).Where("email = ?", email).Updates(studentUpdate).Error

	return err
}
