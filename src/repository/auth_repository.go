package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateStudent(model.Student) error
	GetStudent(email string) (model.Student, error)
	GetDormOwner(email string) (model.DormOwner, error)
}

func AuthRepositoryHandler(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

type authRepository struct {
	db *gorm.DB
}

func (repository *authRepository) CreateStudent(student model.Student) error {
	return repository.db.Create(&student).Error
}

func (repository *authRepository) GetStudent(email string) (model.Student, error) {
	var student model.Student

	err := repository.db.Where("email = ?", email).First(&student).Error

	return student, err
}

func (repository *authRepository) GetDormOwner(email string) (model.DormOwner, error) {
	var dormOwner model.DormOwner

	err := repository.db.Where("email = ?", email).First(&dormOwner).Error

	return dormOwner, err
}
