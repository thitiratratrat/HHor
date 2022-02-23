package service

import (
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type StudentService interface {
	GetFaculties() []string
	GetStudent(id string) model.Student
	UpdateStudent(id string, studentUpdate map[string]interface{}) model.Student
	UpdateStudentPetPictures(id string, pictureUrls []string) model.Student
}

func StudentServiceHandler(studentRepository repository.StudentRepository) StudentService {
	return &studentService{
		studentRepository: studentRepository,
	}
}

type studentService struct {
	studentRepository repository.StudentRepository
}

func (studentService *studentService) GetFaculties() []string {
	return studentService.studentRepository.FindFaculties()
}

func (studentService *studentService) GetStudent(id string) model.Student {
	student, err := studentService.studentRepository.FindStudent(id)

	if err != nil {
		panic(errortype.ErrResourceNotFound)
	}

	return student
}

func (studentService *studentService) UpdateStudent(id string, studentUpdate map[string]interface{}) model.Student {
	student, err := studentService.studentRepository.UpdateStudent(id, studentUpdate)

	if err != nil {
		panic(err)
	}

	return student
}

func (studentService *studentService) UpdateStudentPetPictures(id string, pictureUrls []string) model.Student {
	student, err := studentService.studentRepository.UpdateStudentPetPictures(id, pictureUrls)

	if err != nil {
		panic(err)
	}

	return student
}
