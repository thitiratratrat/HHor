package service

import (
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterStudent(dto.RegisterStudentDTO)
	RegisterDormOwner(dto.RegisterDormOwnerDTO)
	LoginStudent(dto.LoginCredentialsDTO) model.Student
	LoginDormOwner(dto.LoginCredentialsDTO) model.DormOwner
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

func (authService *authService) RegisterStudent(registerStudentDTO dto.RegisterStudentDTO) {
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

	_, err := authService.studentRepository.CreateStudent(student)

	if err != nil {
		panic(err)
	}

}

func (authService *authService) RegisterDormOwner(registerDormOwnerDTO dto.RegisterDormOwnerDTO) {
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(registerDormOwnerDTO.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		panic(hashErr)
	}

	dormOwner := model.DormOwner{
		Firstname:   registerDormOwnerDTO.Firstname,
		Lastname:    registerDormOwnerDTO.Lastname,
		Email:       registerDormOwnerDTO.Email,
		Password:    string(hashedPassword),
		LineID:      registerDormOwnerDTO.LineID,
		PhoneNumber: registerDormOwnerDTO.PhoneNumber,
	}

	_, err := authService.dormOwnerRepository.CreateDormOwner(dormOwner)

	if err != nil {
		panic(err)
	}

}

func (authService *authService) LoginStudent(loginCredentialsDTO dto.LoginCredentialsDTO) model.Student {
	student, err := authService.studentRepository.FindStudentByEmail(loginCredentialsDTO.Email)

	if err == nil {
		if comparePassword(loginCredentialsDTO.Password, student.Password) {
			return student
		}

		panic(errortype.ErrUnauthorized)
	}

	panic(errortype.ErrUserNotFound)
}

func (authService *authService) LoginDormOwner(loginCredentialsDTO dto.LoginCredentialsDTO) model.DormOwner {
	dormOwner, err := authService.dormOwnerRepository.FindDormOwnerByEmail(loginCredentialsDTO.Email)

	if err == nil {
		if comparePassword(loginCredentialsDTO.Password, dormOwner.Password) {
			return dormOwner
		}

		panic(errortype.ErrUnauthorized)
	}

	panic(errortype.ErrUserNotFound)
}

func comparePassword(password string, hashedPassword string) bool {
	compareError := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return compareError == nil
}
