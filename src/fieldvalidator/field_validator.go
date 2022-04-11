package fieldvalidator

import (
	"regexp"

	"github.com/thitiratratrat/hhor/src/service"
)

type FieldValidator interface {
	ValidFaculty(inputFaculty []string) bool
	ValidDormZone(inputDormZones []string) bool
	ValidRoomFacility(inputRoomFacilities []string) bool
	ValidDormFacility(inputDormFacilities []string) bool
	ValidPhoneNumber(phohneNumber string) bool
}

func FieldValidatorHandler(dormService service.DormService, roomService service.RoomService, studentService service.StudentService) FieldValidator {
	return &fieldValidator{
		dormService:    dormService,
		roomService:    roomService,
		studentService: studentService,
	}
}

type fieldValidator struct {
	dormService    service.DormService
	roomService    service.RoomService
	studentService service.StudentService
}

func (fieldValidator *fieldValidator) ValidDormZone(inputDormZones []string) bool {
	dormZones := fieldValidator.dormService.GetDormZones()

	return containsAllItems(inputDormZones, dormZones)
}

func (fieldValidator *fieldValidator) ValidDormFacility(inputDormFacilities []string) bool {
	dormFacilities := fieldValidator.dormService.GetAllDormFacilities()

	return containsAllItems(inputDormFacilities, dormFacilities)
}

func (fieldValidator *fieldValidator) ValidFaculty(inputFaculties []string) bool {
	faculties := fieldValidator.studentService.GetFaculties()

	return containsAllItems(inputFaculties, faculties)
}

func (fieldValidator *fieldValidator) ValidRoomFacility(inputRoomFacilities []string) bool {
	roomFacilities := fieldValidator.roomService.GetAllRoomFacilities()

	return containsAllItems(inputRoomFacilities, roomFacilities)
}

func (fieldValidator *fieldValidator) ValidPhoneNumber(phoneNumber string) bool {
	r, _ := regexp.Compile("^(0[689]{1})+([0-9]{8})+$")

	return r.MatchString(phoneNumber)
}

func containsAllItems(inputList []string, mainList []string) bool {
	mainListMap := make(map[string]bool)

	for _, mainItem := range mainList {
		mainListMap[mainItem] = true
	}

	for _, inputItem := range inputList {
		if _, ok := mainListMap[inputItem]; !ok {
			return false
		}
	}

	return true
}
