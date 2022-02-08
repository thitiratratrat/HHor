package repository

import (
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(model.Student) (model.Student, error)
	GetStudent(id string) (model.Student, error)
	GetStudentByEmail(email string) (model.Student, error)
	UpdateStudent(id string, studentUpdate map[string]interface{}) (model.Student, error)
	UpdateStudentPetPictures(id string, pictureUrls []string) (model.Student, error)
	GetFaculties() []string
}

func StudentRepositoryHandler(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

type studentRepository struct {
	db *gorm.DB
}

func (repository *studentRepository) CreateStudent(student model.Student) (model.Student, error) {
	err := repository.db.Create(&student).Error

	if err != nil {
		return model.Student{}, err
	}

	return repository.GetStudent(student.Email)
}

func (repository *studentRepository) GetStudent(id string) (model.Student, error) {
	var student model.Student

	err := repository.db.Preload("SmokeHabit").Preload("StudyHabit").Preload("RoomCareHabit").Preload("PetHabit").Preload("SleepHabit").Preload("PreferredSmokeHabit").Preload("PreferredStudyHabit").Preload("PreferredRoomCareHabit").Preload("PreferredPetHabit").Preload("PreferredSleepHabit").Preload("PetPictures").Where("id = ?", id).First(&student).Error

	return student, err
}

func (repository *studentRepository) GetStudentByEmail(email string) (model.Student, error) {
	var student model.Student

	err := repository.db.Preload("SmokeHabit").Preload("StudyHabit").Preload("RoomCareHabit").Preload("PetHabit").Preload("SleepHabit").Preload("PreferredSmokeHabit").Preload("PreferredStudyHabit").Preload("PreferredRoomCareHabit").Preload("PreferredPetHabit").Preload("PreferredSleepHabit").Preload("PetPictures").Where("email = ?", email).First(&student).Error

	return student, err
}

func (repository *studentRepository) UpdateStudent(id string, studentUpdate map[string]interface{}) (model.Student, error) {
	err := repository.db.Model(&model.Student{}).Where("id = ?", id).Updates(studentUpdate).Error

	if err != nil {
		return model.Student{}, err
	}

	return repository.GetStudent(id)
}

func (repository *studentRepository) UpdateStudentPetPictures(id string, pictureUrls []string) (model.Student, error) {
	var petPictures []model.PetPicture

	for _, pictureUrl := range pictureUrls {
		petPictures = append(petPictures, model.PetPicture{PictureUrl: pictureUrl, StudentID: id})
	}

	repository.db.Table("pet_pictures").Where("student_email = ?", id).Delete(model.PetPicture{})
	repository.db.Create(&petPictures)

	student, err := repository.GetStudent(id)

	return student, err
}

func (repository *studentRepository) GetFaculties() []string {
	var faculties []string

	repository.db.Model(&model.Faculty{}).Pluck("name", &faculties)

	return faculties
}
