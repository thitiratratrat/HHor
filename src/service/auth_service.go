package service

import (
	"fmt"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"github.com/thitiratratrat/hhor/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterStudent(dto.RegisterStudentDTO)
	RegisterDormOwner(dto.RegisterDormOwnerDTO)
	LoginStudent(dto.LoginCredentialsDTO) model.Student
	LoginDormOwner(dto.LoginCredentialsDTO) model.DormOwner
	VerifyCodeStudent(dto.VerifyCodeDTO)
	VerifyCodeDormOwner(dto.VerifyCodeDTO)
	ResendCodeStudent(dto.ResendCodeDTO)
	ResendCodeDormOwner(dto.ResendCodeDTO)
}

func AuthServiceHandler(emailService EmailService, studentRepository repository.StudentRepository, dormOwnerRepository repository.DormOwnerRepository) AuthService {
	return &authService{
		emailService:        emailService,
		studentRepository:   studentRepository,
		dormOwnerRepository: dormOwnerRepository,
	}
}

type authService struct {
	emailService        EmailService
	studentRepository   repository.StudentRepository
	dormOwnerRepository repository.DormOwnerRepository
}

func (authService *authService) RegisterStudent(registerStudentDTO dto.RegisterStudentDTO) {
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(registerStudentDTO.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		panic(hashErr)
	}

	code, codeErr := utils.GenerateCode()

	if codeErr != nil {
		panic(codeErr)
	}

	student := model.Student{
		Firstname:        registerStudentDTO.Firstname,
		Lastname:         registerStudentDTO.Lastname,
		ID:               registerStudentDTO.StudentID,
		Email:            registerStudentDTO.Email,
		Password:         string(hashedPassword),
		EnrollmentYear:   registerStudentDTO.EnrollmentYear,
		GenderName:       registerStudentDTO.Gender,
		FacultyName:      registerStudentDTO.Faculty,
		VerificationCode: &code,
		HasVerified:      false,
	}

	_, err := authService.studentRepository.CreateStudent(student)

	if err != nil {
		panic(err)
	}

	authService.emailService.SendEmail(student.Email, code)
}

func (authService *authService) RegisterDormOwner(registerDormOwnerDTO dto.RegisterDormOwnerDTO) {
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(registerDormOwnerDTO.Password), bcrypt.DefaultCost)

	if hashErr != nil {
		panic(hashErr)
	}

	code, codeErr := utils.GenerateCode()

	if codeErr != nil {
		panic(codeErr)
	}

	dormOwner := model.DormOwner{
		Firstname:        registerDormOwnerDTO.Firstname,
		Lastname:         registerDormOwnerDTO.Lastname,
		Email:            registerDormOwnerDTO.Email,
		Password:         string(hashedPassword),
		LineID:           registerDormOwnerDTO.LineID,
		PhoneNumber:      registerDormOwnerDTO.PhoneNumber,
		VerificationCode: &code,
		HasVerified:      false,
	}

	_, err := authService.dormOwnerRepository.CreateDormOwner(dormOwner)

	if err != nil {
		panic(err)
	}

	authService.emailService.SendEmail(dormOwner.Email, code)
}

func (authService *authService) LoginStudent(loginCredentialsDTO dto.LoginCredentialsDTO) model.Student {
	student, err := authService.studentRepository.FindStudentByEmail(loginCredentialsDTO.Email)

	if err == nil {
		if comparePassword(loginCredentialsDTO.Password, student.Password) {
			if !student.HasVerified {
				panic(errortype.ErrHasNotVerifiedCode)
			}

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
			if !dormOwner.HasVerified {
				panic(errortype.ErrHasNotVerifiedCode)
			}

			return dormOwner
		}

		panic(errortype.ErrUnauthorized)
	}

	panic(errortype.ErrUserNotFound)
}

func (authService *authService) VerifyCodeStudent(verifyCodeDTO dto.VerifyCodeDTO) {
	student, err := authService.studentRepository.FindStudentByEmail(verifyCodeDTO.Email)

	if err != nil {
		panic(errortype.ErrUserNotFound)
	}

	if *student.VerificationCode != verifyCodeDTO.Code {
		panic(errortype.ErrIncorrectCode)
	}

	authService.studentRepository.UpdateStudent(student.ID, map[string]interface{}{"has_verified": true})
}

func (authService *authService) VerifyCodeDormOwner(verifyCodeDTO dto.VerifyCodeDTO) {
	dormOwner, err := authService.dormOwnerRepository.FindDormOwnerByEmail(verifyCodeDTO.Email)

	if err != nil {
		panic(errortype.ErrUserNotFound)
	}

	if *dormOwner.VerificationCode != verifyCodeDTO.Code {
		panic(errortype.ErrIncorrectCode)
	}

	authService.dormOwnerRepository.UpdateDormOwner(fmt.Sprintf("%v", dormOwner.ID), model.DormOwner{HasVerified: true})
}

func (authService *authService) ResendCodeDormOwner(resendCodeDTO dto.ResendCodeDTO) {
	dormOwner, err := authService.dormOwnerRepository.FindDormOwnerByEmail(resendCodeDTO.Email)

	if err != nil {
		panic(errortype.ErrUserNotFound)
	}

	if dormOwner.HasVerified {
		panic(errortype.ErrAlreadyVerified)
	}

	code, codeErr := utils.GenerateCode()

	if codeErr != nil {
		panic(codeErr)
	}

	authService.dormOwnerRepository.UpdateDormOwner(fmt.Sprintf("%v", dormOwner.ID), model.DormOwner{VerificationCode: &code})
	authService.emailService.SendEmail(dormOwner.Email, code)
}

func (authService *authService) ResendCodeStudent(resendCodeDTO dto.ResendCodeDTO) {
	student, err := authService.studentRepository.FindStudentByEmail(resendCodeDTO.Email)

	if err != nil {
		panic(errortype.ErrUserNotFound)
	}

	if student.HasVerified {
		panic(errortype.ErrAlreadyVerified)
	}

	code, codeErr := utils.GenerateCode()

	if codeErr != nil {
		panic(codeErr)
	}

	authService.studentRepository.UpdateStudent(string(student.ID), map[string]interface{}{"verification_code": code})
	authService.emailService.SendEmail(student.Email, code)
}

func comparePassword(password string, hashedPassword string) bool {
	compareError := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return compareError == nil
}
