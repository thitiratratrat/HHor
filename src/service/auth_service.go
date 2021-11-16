package service

import (
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GetFaculties() []string
	RegisterStudent(dto.RegisterStudentDTO) error
	LoginStudent(dto.LoginCredentialsDTO) error
}

func AuthServiceHandler(facultyRepository repository.FacultyRepository, studentRepository repository.StudentRepository, dormOwnerRepository repository.DormOwnerRepository) AuthService {
	return &authService{
		facultyRepository:   facultyRepository,
		studentRepository:   studentRepository,
		dormOwnerRepository: dormOwnerRepository,
	}
}

type authService struct {
	facultyRepository   repository.FacultyRepository
	studentRepository   repository.StudentRepository
	dormOwnerRepository repository.DormOwnerRepository
}

func (authService *authService) GetFaculties() []string {
	faculties := authService.facultyRepository.GetFaculties()

	return faculties
}

func (authService *authService) RegisterStudent(registerStudentDTO dto.RegisterStudentDTO) error {
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(registerStudentDTO.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		panic(hashErr)
	}

	student := model.Student{
		Firstname:   registerStudentDTO.Firstname,
		Lastname:    registerStudentDTO.Lastname,
		StudentID:   registerStudentDTO.StudentID,
		Email:       registerStudentDTO.Email,
		Password:    string(hashedPassword),
		YearOfStudy: registerStudentDTO.YearOfStudy,
		GenderName:  registerStudentDTO.Gender,
		FacultyName: registerStudentDTO.Faculty,
	}

	return authService.studentRepository.CreateStudent(student)
}

func (authService *authService) LoginStudent(loginCredentialsDTO dto.LoginCredentialsDTO) error {
	student, getStudentError := authService.studentRepository.GetStudent(loginCredentialsDTO.Email)

	if getStudentError == nil {
		if comparePassword(loginCredentialsDTO.Password, student.Password) {
			return nil
		}

		return errortype.ErrUnauthorized
	}

	return errortype.ErrUserNotFound
}

func comparePassword(password string, hashedPassword string) bool {
	compareError := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return compareError == nil
}
