package service

import (
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type StudentService interface {
	GetStudent(id string) (model.Student, error)
	UpdateStudent(id string, studentUpdate map[string]interface{}) (model.Student, error)
	UpdateStudentPetPictures(id string, pictureUrls []string) (model.Student, error)
}

func StudentServiceHandler(studentRepository repository.StudentRepository) StudentService {
	return &studentService{
		studentRepository: studentRepository,
	}
}

type studentService struct {
	studentRepository repository.StudentRepository
}

func (studentService *studentService) GetStudent(id string) (model.Student, error) {
	return studentService.studentRepository.GetStudent(id)
}

func (studentService *studentService) UpdateStudent(id string, studentUpdate map[string]interface{}) (model.Student, error) {
	return studentService.studentRepository.UpdateStudent(id, studentUpdate)
}

func (studentService *studentService) UpdateStudentPetPictures(id string, pictureUrls []string) (model.Student, error) {
	return studentService.studentRepository.UpdateStudentPetPictures(id, pictureUrls)
}
