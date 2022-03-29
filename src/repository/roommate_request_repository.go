package repository

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

//TODO: separate into each request type?
type RoommateRequestRepository interface {
	FindRoommateRequestWithRegisteredDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithRegisteredDorm
	FindRoommateRequestWithUnregisteredDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithUnregisteredDorm
	FindRoommateRequestWithNoRooms(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithNoRoom
	FindRoommateRequestWithNoRoom(id string) (model.RoommateRequestWithNoRoom, error)
	FindRoommateRequestWithUnregisteredDorm(id string) (model.RoommateRequestWithUnregisteredDorm, error)
	FindRoommateRequestWithRegisteredDorm(id string) (model.RoommateRequestWithRegisteredDorm, error)
	CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoom model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error)
	CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error)
	CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error)
	UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithRegisteredDorm, error)
	UpdateRoommateRequestWithUnregisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithUnregisteredDorm, error)
	UpdateRoommateRequestRegDorm(id string, req model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error)
	UpdateRoommateRequestUnregDorm(id string, roommateRequest model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error)
	UpdateRoommateRequestNoRoom(id string, roommateRequest model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error)
	DeleteRoommateRequestRegDorm(id string) error
	DeleteRoommateRequestUnregDorm(id string) error
	DeleteRoommateRequestNoRoom(id string) error
}

func RoommateRequestRepositoryHandler(db *gorm.DB) RoommateRequestRepository {
	return &roommateRequestRepository{
		db: db,
	}
}

type roommateRequestRepository struct {
	db *gorm.DB
}

func (repository *roommateRequestRepository) FindRoommateRequestWithRegisteredDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithRegisteredDorm {
	var roommateRequests []model.RoommateRequestWithRegisteredDorm
	condition := repository.getRoomCondition(roommateRequestRoomFilterDTO, constant.RoommateRequestRegDorm)

	repository.db.Preload("RoomPictures").Joins("Dorm").Joins("Room").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateRequestRepository) FindRoommateRequestWithUnregisteredDorms(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []model.RoommateRequestWithUnregisteredDorm {
	var roommateRequests []model.RoommateRequestWithUnregisteredDorm
	condition := repository.getRoomCondition(roommateRequestRoomFilterDTO, constant.RoommateRequestUnregDorm)

	repository.db.Preload("RoomPictures").Preload("RoomFacilities").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateRequestRepository) FindRoommateRequestWithNoRooms(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithNoRoom {
	var roommateRequests []model.RoommateRequestWithNoRoom
	condition := repository.getCondition(roommateRequestFilterDTO, constant.RoommateRequestNoRoom)

	repository.db.Preload("Zones").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateRequestRepository) FindRoommateRequestWithNoRoom(id string) (model.RoommateRequestWithNoRoom, error) {
	var roommateRequest model.RoommateRequestWithNoRoom

	err := repository.db.Preload("Zones").Joins("Student").First(&roommateRequest, id).Error

	return roommateRequest, err
}

func (repository *roommateRequestRepository) FindRoommateRequestWithUnregisteredDorm(id string) (model.RoommateRequestWithUnregisteredDorm, error) {
	var roommateRequest model.RoommateRequestWithUnregisteredDorm

	err := repository.db.Preload("RoomPictures").Preload("RoomFacilities").Joins("Student").First(&roommateRequest, id).Error

	return roommateRequest, err
}

func (repository *roommateRequestRepository) FindRoommateRequestWithRegisteredDorm(id string) (model.RoommateRequestWithRegisteredDorm, error) {
	var roommateRequest model.RoommateRequestWithRegisteredDorm

	err := repository.db.Preload("Room.Pictures").Preload("Room.Facilities").Preload("RoomPictures").Joins("Dorm").Joins("Room").Joins("Student").First(&roommateRequest, id).Error

	return roommateRequest, err
}

func (repository *roommateRequestRepository) CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoom model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error) {
	err := repository.db.Create(&roommateRequestWithNoRoom).Error

	return roommateRequestWithNoRoom, err
}

func (repository *roommateRequestRepository) CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error) {
	err := repository.db.Create(&roommateRequestWithRegisteredDorm).Error

	return roommateRequestWithRegisteredDorm, err
}

func (repository *roommateRequestRepository) CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error) {
	err := repository.db.Create(&roommateRequestWithUnregisteredDorm).Error

	return roommateRequestWithUnregisteredDorm, err
}

func (repository *roommateRequestRepository) UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithRegisteredDorm, error) {
	var roomPictures []model.RoommateRequestRegisteredDormPicture
	var roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm

	for _, pictureUrl := range pictureUrls {
		roomPictures = append(roomPictures, model.RoommateRequestRegisteredDormPicture{
			PictureUrl: pictureUrl,
			RoommateRequestWithRegisteredDormStudentID: id,
		})
	}

	repository.db.Table("roommate_request_registered_dorm_pictures").Where("roommate_request_with_registered_dorm_student_id = ?", id).Delete(model.RoommateRequestRegisteredDormPicture{})
	repository.db.Create(&roomPictures)

	err := repository.db.Preload("RoomPictures").Where("student_id = ?", id).First(&roommateRequestWithRegisteredDorm).Error

	return roommateRequestWithRegisteredDorm, err
}

func (repository *roommateRequestRepository) UpdateRoommateRequestWithUnregisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithUnregisteredDorm, error) {
	var roomPictures []model.RoommateRequestUnregisteredDormPicture
	var roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm

	for _, pictureUrl := range pictureUrls {
		roomPictures = append(roomPictures, model.RoommateRequestUnregisteredDormPicture{
			PictureUrl: pictureUrl,
			RoommateRequestWithUnregisteredDormStudentID: id,
		})
	}

	repository.db.Table("roommate_request_unregistered_dorm_pictures").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.RoommateRequestUnregisteredDormPicture{})
	repository.db.Create(&roomPictures)

	err := repository.db.Preload("RoomPictures").Preload("RoomFacilities").Where("student_id = ?", id).First(&roommateRequestWithUnregisteredDorm).Error

	return roommateRequestWithUnregisteredDorm, err
}

func (repository *roommateRequestRepository) UpdateRoommateRequestRegDorm(id string, roommateRequest model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error) {
	err := repository.db.Save(&roommateRequest).Error

	if err != nil {
		return model.RoommateRequestWithRegisteredDorm{}, err
	}

	return repository.FindRoommateRequestWithRegisteredDorm(id)
}

func (repository *roommateRequestRepository) UpdateRoommateRequestUnregDorm(id string, roommateRequest model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error) {
	err := repository.db.Model(&model.RoommateRequestWithUnregisteredDorm{StudentID: roommateRequest.StudentID}).Omit("RoomFacilities").Updates(roommateRequest).Error

	if err != nil {
		return model.RoommateRequestWithUnregisteredDorm{}, err
	}

	repository.db.Table("roommate_request_unregistered_dorm_room_facility").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.AllRoomFacility{})
	repository.db.Model(&roommateRequest).Association("RoomFacilities").Append(roommateRequest.RoomFacilities)

	return repository.FindRoommateRequestWithUnregisteredDorm(id)
}

func (repository *roommateRequestRepository) UpdateRoommateRequestNoRoom(id string, roommateRequest model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error) {
	err := repository.db.Model(&model.RoommateRequestWithNoRoom{StudentID: roommateRequest.StudentID}).Omit("Zones").Updates(roommateRequest).Error

	if err != nil {
		return model.RoommateRequestWithNoRoom{}, err
	}

	repository.db.Table("roommate_request_no_room_zone").Where("roommate_request_with_no_room_student_id = ?", id).Delete(model.DormZone{})
	repository.db.Model(&roommateRequest).Association("Zones").Append(roommateRequest.Zones)

	return repository.FindRoommateRequestWithNoRoom(id)
}

func (repository *roommateRequestRepository) DeleteRoommateRequestRegDorm(id string) error {
	repository.db.Table("roommate_request_registered_dorm_pictures").Where("roommate_request_with_registered_dorm_student_id = ?", id).Delete(model.RoommateRequestRegisteredDormPicture{})

	err := repository.db.Delete(&model.RoommateRequestWithRegisteredDorm{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *roommateRequestRepository) DeleteRoommateRequestUnregDorm(id string) error {
	repository.db.Table("roommate_request_unregistered_dorm_pictures").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.RoommateRequestUnregisteredDormPicture{})
	repository.db.Table("roommate_request_unregistered_dorm_room_facility").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.AllRoomFacility{})

	err := repository.db.Delete(&model.RoommateRequestWithUnregisteredDorm{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *roommateRequestRepository) DeleteRoommateRequestNoRoom(id string) error {
	repository.db.Table("roommate_request_no_room_zone").Where("roommate_request_with_no_room_student_id = ?", id).Delete(model.DormZone{})
	err := repository.db.Delete(&model.RoommateRequestWithNoRoom{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *roommateRequestRepository) getRoomCondition(roommateRequestFilterDTO dto.RoommateRequestRoomFilterDTO, requestType constant.RoommateRequestType) string {
	nameCondition := repository.getNameCondition(roommateRequestFilterDTO.DormName, requestType)
	roommateCondition := repository.getNumberOfRoommatesCondition(roommateRequestFilterDTO.NumberOfRoommates)
	roommFacilityCondition := repository.getRoomFacilityCondition(roommateRequestFilterDTO.RoomFacilities, requestType)
	noRoomCondition := repository.getCondition(roommateRequestFilterDTO.RoommateRequestFilterDTO, requestType)

	condition := fmt.Sprintf("%s %s %s AND %s", nameCondition, roommateCondition, roommFacilityCondition, noRoomCondition)

	return condition
}

func (repository *roommateRequestRepository) getCondition(roommateRequestFilterDTO dto.RoommateRequestFilterDTO, requestType constant.RoommateRequestType) string {
	zoneCondition := repository.getZoneCondition(roommateRequestFilterDTO.Zone, requestType)
	genderCondition := repository.getGenderCondition(roommateRequestFilterDTO.Gender)
	facultyCondition := repository.getFacultyCondition(roommateRequestFilterDTO.Faculty)
	yearOfStudyCondition := repository.getYearOfStudyCondition(roommateRequestFilterDTO.YearOfStudy)
	budgetCondition := repository.getBudgetCondition(roommateRequestFilterDTO.LowerPrice, roommateRequestFilterDTO.UpperPrice, requestType)
	preferenceCondition := repository.getPreferenceCondition(roommateRequestFilterDTO.SmokeHabitID, roommateRequestFilterDTO.RoomCareHabitID, roommateRequestFilterDTO.SleepHabitID, roommateRequestFilterDTO.StudyHabitID, roommateRequestFilterDTO.PetHabitID)
	condition := fmt.Sprintf("true %s %s %s %s %s %s", zoneCondition, genderCondition, facultyCondition, yearOfStudyCondition, budgetCondition, preferenceCondition)

	return condition
}

func (repository *roommateRequestRepository) getNameCondition(name *string, requestType constant.RoommateRequestType) string {
	switch {
	case requestType == constant.RoommateRequestRegDorm && name == nil:
		return `"Dorm".name` + " LIKE '%%'"
	case requestType == constant.RoommateRequestRegDorm:
		return `"Dorm".name` + " LIKE '%" + *name + "%'"
	case requestType == constant.RoommateRequestUnregDorm && name == nil:
		return "dorm_name LIKE '%%'"
	default:
		return "dorm_name LIKE '%" + *name + "%'"
	}
}

func (repository *roommateRequestRepository) getZoneCondition(zone *string, requestType constant.RoommateRequestType) string {
	if zone == nil {
		return ""
	}

	switch requestType {
	case constant.RoommateRequestNoRoom:
		return fmt.Sprintf("AND '%s' IN (select dorm_zone_name from roommate_request_no_room_zone where student_id = roommate_request_with_no_room_student_id)", *zone)
	default:
		return fmt.Sprintf("AND dorm_zone_name = '%s'", *zone)
	}
}

func (repository *roommateRequestRepository) getGenderCondition(gender []string) string {
	if len(gender) == 0 {
		return ""
	}

	formattedGender := "'" + strings.Join(gender, "', '") + "'"

	return fmt.Sprintf("AND gender_name IN (%s)", formattedGender)
}

func (repository *roommateRequestRepository) getFacultyCondition(faculty []string) string {
	if len(faculty) == 0 {
		return ""
	}

	formattedFaculty := "'" + strings.Join(faculty, "', '") + "'"

	return fmt.Sprintf("AND faculty_name IN (%s)", formattedFaculty)
}

func (repository *roommateRequestRepository) getYearOfStudyCondition(yearOfStudy []int) string {
	if len(yearOfStudy) == 0 {
		return ""
	}

	sort.Ints(yearOfStudy)
	highestYear := yearOfStudy[len(yearOfStudy)-1]
	highCondition := ""
	currentYear := time.Now().Year()
	formattedYearofStudy := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(yearOfStudy)), ","), "[]")

	if highestYear >= 4 {
		highCondition = fmt.Sprintf("OR %d - enrollment_year >= %d", currentYear, highestYear)
	}

	return fmt.Sprintf("AND (%d - enrollment_year IN (%s) %s)", currentYear, formattedYearofStudy, highCondition)
}

func (repository *roommateRequestRepository) getBudgetCondition(lowerPrice *int, upperPrice *int, requestType constant.RoommateRequestType) string {
	if lowerPrice == nil || upperPrice == nil || *upperPrice < *lowerPrice {
		return ""
	}

	switch requestType {
	case constant.RoommateRequestNoRoom:
		return fmt.Sprintf("AND budget BETWEEN %d AND %d", *lowerPrice, *upperPrice)
	default:
		return fmt.Sprintf("AND shared_room_price BETWEEN %d AND %d", *lowerPrice, *upperPrice)
	}
}

func (repository *roommateRequestRepository) getNumberOfRoommatesCondition(numberOfRoomates []int) string {
	if len(numberOfRoomates) == 0 {
		return ""
	}

	sort.Ints(numberOfRoomates)
	highestRoommate := numberOfRoomates[len(numberOfRoomates)-1]
	highCondition := ""
	formattedRoommates := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numberOfRoomates)), ","), "[]")

	if highestRoommate >= 4 {
		highCondition = fmt.Sprintf("OR number_of_roommates >= %d", highestRoommate)
	}

	return fmt.Sprintf("AND (number_of_roommates IN (%s) %s)", formattedRoommates, highCondition)
}

func (repository *roommateRequestRepository) getRoomFacilityCondition(roomFacilities []string, requestType constant.RoommateRequestType) string {
	if len(roomFacilities) == 0 {
		return ""
	}

	formattedRoomFacilities := "'" + strings.Join(roomFacilities, "', '") + "'"

	switch requestType {
	case constant.RoommateRequestRegDorm:
		return fmt.Sprintf("AND %d = (select count(*) from room_facility where"+`"Room".id `+"= room_facility.room_id and all_room_facility_name IN (%s))", len(roomFacilities), formattedRoomFacilities)
	default:
		return fmt.Sprintf("AND %d = (select count(*) from roommate_request_unregistered_dorm_room_facility where student_id = roommate_request_with_unregistered_dorm_student_id and all_room_facility_name IN (%s))", len(roomFacilities), formattedRoomFacilities)
	}
}

func (repository *roommateRequestRepository) getPreferenceCondition(smokeHabit *string, roomCareHabit *string, sleepHabit *string, studyHabit *string, petHabit *string) string {
	var smokeCondition, roomCareCondition, sleepingCondition, studyCondition, petCondition string

	if smokeHabit != nil && len(*smokeHabit) != 0 {
		smokeCondition = fmt.Sprintf("AND smoke_habit_id = %v", *smokeHabit)
	}

	if roomCareHabit != nil && len(*roomCareHabit) != 0 {
		roomCareCondition = fmt.Sprintf("AND room_care_habit_id = %v", *roomCareHabit)
	}

	if sleepHabit != nil && len(*sleepHabit) != 0 {
		sleepingCondition = fmt.Sprintf("AND sleep_habit_id = %v", *sleepHabit)
	}

	if studyHabit != nil && len(*studyHabit) != 0 {
		studyCondition = fmt.Sprintf("AND study_habit_id = %v", *studyHabit)
	}

	if petHabit != nil && len(*petHabit) != 0 {
		petCondition = fmt.Sprintf("AND pet_habit_id = %v", *petHabit)
	}

	return fmt.Sprintf("%s %s %s %s %s", smokeCondition, roomCareCondition, sleepingCondition, studyCondition, petCondition)
}
