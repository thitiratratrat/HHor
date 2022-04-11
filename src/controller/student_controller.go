package controller

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/fieldvalidator"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

type StudentController interface {
	GetHabits(context *gin.Context)
	GetFaculties(context *gin.Context)
	GetStudent(context *gin.Context)
	UpdateStudent(context *gin.Context)
	UpdateHabit(context *gin.Context)
	UpdatePreference(context *gin.Context)
	UploadPicture(context *gin.Context)
}

func StudentControllerHandler(studentService service.StudentService, fieldValidator fieldvalidator.FieldValidator) StudentController {
	return &studentController{
		studentService: studentService,
		fieldValidator: fieldValidator,
	}
}

type studentController struct {
	studentService service.StudentService
	fieldValidator fieldvalidator.FieldValidator
}

// @Summary update student detail
// @Tags student
// @Produce json
// @Param id path string true "Student ID"
// @Param data body dto.UpdateStudentDTO false "student update"
// @Success 200 {object} model.Student "OK"
// @Router /student/{id} [patch]
func (studentController *studentController) UpdateStudent(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	id := context.Param("id")
	var studentUpdateDTO dto.UpdateStudentDTO
	validate := validator.New()
	_ = validate.RegisterValidation("faculty", func(fl validator.FieldLevel) bool {
		return studentController.fieldValidator.ValidFaculty([]string{fl.Field().String()})
	})
	bindErr := context.ShouldBind(&studentUpdateDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(studentUpdateDTO)

	if validateError != nil {
		panic(validateError)
	}

	studentUpdateMap := structs.Map(studentUpdateDTO)
	updatedStudent := studentController.studentService.UpdateStudent(id, studentUpdateMap)

	context.IndentedJSON(http.StatusOK, updatedStudent)
}

// @Summary update student habit
// @Tags student
// @Produce json
// @Param id path string true "Student ID"
// @Param data body dto.UpdateHabitDTO false "habit update"
// @Success 200 {object} model.Student "OK"
// @Router /student/{id}/habit [patch]
func (studentController *studentController) UpdateHabit(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	id := context.Param("id")
	var studentUpdateDTO dto.UpdateHabitDTO
	validate := validator.New()
	bindErr := context.ShouldBind(&studentUpdateDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(studentUpdateDTO)

	if validateError != nil {
		panic(validateError)
	}

	studentUpdateMap := structs.Map(studentUpdateDTO)
	updatedStudent := studentController.studentService.UpdateStudent(id, studentUpdateMap)

	context.IndentedJSON(http.StatusOK, updatedStudent)
}

// @Summary update student preference
// @Tags student
// @Produce json
// @Param id path string true "Student ID"
// @Param data body dto.UpdatePreferenceDTO false "preference update"
// @Success 200 {object} model.Student "OK"
// @Router /student/{id}/preference [patch]
func (studentController *studentController) UpdatePreference(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	id := context.Param("id")
	var studentUpdateDTO dto.UpdatePreferenceDTO
	validate := validator.New()
	bindErr := context.ShouldBind(&studentUpdateDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(studentUpdateDTO)

	if validateError != nil {
		panic(validateError)
	}

	studentUpdateMap := structs.Map(studentUpdateDTO)
	updatedStudent := studentController.studentService.UpdateStudent(id, studentUpdateMap)

	context.IndentedJSON(http.StatusOK, updatedStudent)
}

// @Summary upload profile picture
// @Tags student
// @Accept  multipart/form-data
// @Produce json
// @Success 200 {object} model.Student "OK"
// @Param id path string true "Student ID"
// @Param   profile_picture formData file false  "profile picture"
// @Param   pet_pictures formData file false  "upload multiple pet pictures,test this out in postman"
// @Param data formData dto.StudentPictureDTO true  "profile picture"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /student/{id}/picture [put]
func (studentController *studentController) UploadPicture(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	id := context.Param("id")
	var studentPictureDTO dto.StudentPictureDTO
	validate := validator.New()
	bindErr := context.ShouldBind(&studentPictureDTO)

	if bindErr != nil {
		panic(bindErr)
	}

	validateError := validate.Struct(studentPictureDTO)

	if validateError != nil {
		panic(validateError)
	}

	var updatedStudent model.Student

	if studentPictureDTO.ProfilePicture != nil {
		filename := id + ".png"
		file, _, err := context.Request.FormFile("profile_picture")

		if err != nil {
			panic(err)
		}

		pictureUrl := utils.UploadPicture(file, constant.StudentProfilePictureFolder, filename, context.Request)
		updatedStudent = studentController.studentService.UpdateStudent(id, map[string]interface{}{"picture_url": pictureUrl})
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

			petPictureUrl := utils.UploadPicture(picture, fmt.Sprintf("%s%s/", constant.PetPicturesFolder, id), petPicture.Filename, context.Request)
			petPictureUrls = append(petPictureUrls, petPictureUrl)
		}

		updatedStudent = studentController.studentService.UpdateStudentPetPictures(id, petPictureUrls)
	}

	context.IndentedJSON(http.StatusOK, updatedStudent)
}

// @Summary get student profile
// @Tags student
// @Produce json
// @Success 200 {object} model.Student "OK"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Param id path string true "Student ID"
// @Router /student/{id} [get]
func (studentController *studentController) GetStudent(context *gin.Context) {
	defer utils.RecoverInvalidInput(context)

	id := context.Param("id")
	student := studentController.studentService.GetStudent(id)

	context.IndentedJSON(http.StatusOK, student)
}

// @Summary get faculties
// @Tags student
// @Produce json
// @Success 200 {array} string "OK"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /student/faculty [get]
func (studentController *studentController) GetFaculties(context *gin.Context) {
	faculties := studentController.studentService.GetFaculties()

	context.IndentedJSON(http.StatusOK, faculties)
}

// @Summary get habits
// @Tags student
// @Produce json
// @Success 200 {object} dto.HabitDTO "OK"
// @Failure 400,404,500 {object} dto.ErrorResponse
// @Router /student/habit [get]
func (studentController *studentController) GetHabits(context *gin.Context) {
	habits := studentController.studentService.GetHabits()

	context.IndentedJSON(http.StatusOK, habits)
}
