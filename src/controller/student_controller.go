package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
	"gorm.io/gorm"
)

type StudentController interface {
	GetStudent(context *gin.Context)
	UpdateStudent(context *gin.Context)
	UploadPicture(context *gin.Context)
}

func StudentControllerHandler(studentService service.StudentService) StudentController {
	return &studentController{
		studentService: studentService,
	}
}

type studentController struct {
	studentService service.StudentService
}

//TODO: make it able to update null values with golang validator in field.
//form data null is not null. cannot differentiate if passing null value or
//field not present in json.

// @Summary update student detail
// @Description update student detail
// @Tags student
// @Produce json
// @Param id path string true "Student ID"
// @Param data body dto.StudentUpdateDTO false "student update"
// @Success 200 {object} model.Student "OK"
// @Router /student/{id} [patch]
func (studentController *studentController) UpdateStudent(context *gin.Context) {
	id := context.Param("id")

	defer utils.RecoverInvalidInput(context)

	validate := validator.New()
	var studentUpdateDTO dto.StudentUpdateDTO

	bindErr := context.ShouldBind(&studentUpdateDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(studentUpdateDTO)

	if validateError != nil {
		panic(validateError)
	}

	studentUpdateMap := structs.Map(studentUpdateDTO)
	updatedStudent, updateError := studentController.studentService.UpdateStudent(id, studentUpdateMap)

	if updateError != nil {
		panic(updateError)
	}

	context.IndentedJSON(http.StatusOK, updatedStudent)
}

// @Summary upload profile picture
// @Description upload profile picture
// @Tags student
// @Accept  multipart/form-data
// @Produce json
// @Success 200 {object} model.Student "OK"
// @Param id path string true "Student ID"
// @Param   profile_picture formData file false  "profile picture"
// @Param   pet_pictures formData file false  "upload multiple pet pictures,test this out in postman"
// @Param data formData dto.StudentPictureDTO true  "profile picture"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /student/picture/{id} [post]
func (studentController *studentController) UploadPicture(context *gin.Context) {
	id := context.Param("id")

	defer utils.RecoverInvalidInput(context)

	validate := validator.New()

	var studentPictureDTO dto.StudentPictureDTO

	bindErr := context.ShouldBind(&studentPictureDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(studentPictureDTO)

	if validateError != nil {
		panic(validateError)
	}

	student, err := studentController.studentService.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.IndentedJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: errortype.ErrResourceNotFound.Error(),
		})

		return
	} else if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	var updatedStudent model.Student

	if studentPictureDTO.ProfilePicture != nil {
		filename := student.ID + ".png"
		file, _, err := context.Request.FormFile("profile_picture")

		if err != nil {
			panic(err)
		}

		pictureUrl := utils.UploadPicture(file, constant.StudentProfilePictureFolder, filename, context.Request)
		updatedStudent, err = studentController.studentService.UpdateStudent(id, map[string]interface{}{"picture_url": pictureUrl})

		if err != nil {
			panic(err)
		}
	}

	//TODO: delete old files from bucket storage too
	if studentPictureDTO.PetPictures != nil {
		files := context.Request.MultipartForm.File["pet_pictures"]
		var petPictureUrls []string

		for _, petPicture := range files {
			picture, err := petPicture.Open()

			if err != nil {
				panic(err)
			}

			petPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.PetPicturesFolder, student.ID), petPicture.Filename, context.Request)
			petPictureUrls = append(petPictureUrls, petPictureUrl)
		}

		updatedStudent, err = studentController.studentService.UpdateStudentPetPictures(id, petPictureUrls)

		if err != nil {
			panic(err)
		}
	}

	context.IndentedJSON(http.StatusOK, updatedStudent)
}

//TODO: get student's roommate request

// @Summary get student profile
// @Description returns student profile
// @Tags student
// @Produce json
// @Success 200 {object} model.Student "OK"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Param id path string true "Student ID"
// @Router /student/{id} [get]
func (studentController *studentController) GetStudent(context *gin.Context) {
	id := context.Param("id")

	student, err := studentController.studentService.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.IndentedJSON(http.StatusNotFound, dto.ErrorResponse{
			Message: errortype.ErrResourceNotFound.Error(),
		})

		return
	} else if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: err.Error(),
		})

		return
	}

	context.IndentedJSON(http.StatusOK, student)
}
