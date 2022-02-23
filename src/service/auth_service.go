package service

import (
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterStudent(dto.RegisterStudentDTO) model.Student
	LoginStudent(dto.LoginCredentialsDTO)
}

func AuthServiceHandler(studentRepository repository.StudentRepository, dormOwnerRepository repository.DormOwnerRepository) AuthService {
	return &authService{
		studentRepository:   studentRepository,
		dormOwnerRepository: dormOwnerRepository,
	}
}

type authService struct {
	studentRepository   repository.StudentRepository
	dormOwnerRepository repository.DormOwnerRepository
}

func (authService *authService) RegisterStudent(registerStudentDTO dto.RegisterStudentDTO) model.Student {
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(registerStudentDTO.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		panic(hashErr)
	}

	student := model.Student{
		Firstname:      registerStudentDTO.Firstname,
		Lastname:       registerStudentDTO.Lastname,
		ID:             registerStudentDTO.StudentID,
		Email:          registerStudentDTO.Email,
		Password:       string(hashedPassword),
		EnrollmentYear: registerStudentDTO.EnrollmentYear,
		GenderName:     registerStudentDTO.Gender,
		FacultyName:    registerStudentDTO.Faculty,
	}

	student, err := authService.studentRepository.CreateStudent(student)

	if err != nil {
		panic(err)
	}

	return student
}

func (authService *authService) LoginStudent(loginCredentialsDTO dto.LoginCredentialsDTO) {
	student, getStudentError := authService.studentRepository.FindStudentByEmail(loginCredentialsDTO.Email)

	if getStudentError == nil {
		if comparePassword(loginCredentialsDTO.Password, student.Password) {
			return
		}

		panic(errortype.ErrUnauthorized)
	}

	panic(errortype.ErrUserNotFound)
}

func comparePassword(password string, hashedPassword string) bool {
	compareError := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return compareError == nil
}
