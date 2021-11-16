package service

import (
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type StudentService interface {
	GetStudent(email string) (model.Student, error)
	UpdateStudent(email string, studentUpdate map[string]interface{}) error
}

func StudentServiceHandler(studentRepository repository.StudentRepository) StudentService {
	return &studentService{
		studentRepository: studentRepository,
	}
}

type studentService struct {
	studentRepository repository.StudentRepository
}

func (studentService *studentService) GetStudent(email string) (model.Student, error) {
	return studentService.studentRepository.GetStudent(email)
}

func (studentService *studentService) UpdateStudent(email string, studentUpdate map[string]interface{}) error {
	return studentService.studentRepository.UpdateStudent(email, studentUpdate)
}
