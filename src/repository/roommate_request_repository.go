package repository

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/model"
	"gorm.io/gorm"
)

type RoommateRequestRepository interface {
	FindRoommateRequestWithRegisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithRegisteredDorm
	FindRoommateRequestWithUnregisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithUnregisteredDorm
	CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoom model.RoommateRequestWithNoRoom) (model.RoommateRequestWithNoRoom, error)
	CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDorm model.RoommateRequestWithRegisteredDorm) (model.RoommateRequestWithRegisteredDorm, error)
	CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDorm model.RoommateRequestWithUnregisteredDorm) (model.RoommateRequestWithUnregisteredDorm, error)
	UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithRegisteredDorm, error)
	UpdateRoommateRequestWithUnregisteredDormPictures(id string, pictureUrls []string) (model.RoommateRequestWithUnregisteredDorm, error)
}

func RoommateRequestRepositoryHandler(db *gorm.DB) RoommateRequestRepository {
	return &roommateRequestRepository{
		db: db,
	}
}

type roommateRequestRepository struct {
	db *gorm.DB
}

func (repository *roommateRequestRepository) FindRoommateRequestWithRegisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithRegisteredDorm {
	var roommateRequests []model.RoommateRequestWithRegisteredDorm
	var dormNameCondition string

	if roommateRequestFilterDTO.DormName == nil {
		dormNameCondition = `"Dorm".name` + " LIKE '%%'"
	} else {
		dormNameCondition = `"Dorm".name` + " LIKE '%" + *roommateRequestFilterDTO.DormName + "%'"
	}

	condition := dormNameCondition + repository.getCondition(roommateRequestFilterDTO)

	repository.db.Preload("RoomPictures").Joins("Dorm").Joins("Room").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
}

func (repository *roommateRequestRepository) FindRoommateRequestWithUnregisteredDorm(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithUnregisteredDorm {
	var roommateRequests []model.RoommateRequestWithUnregisteredDorm
	var dormNameCondition string

	if roommateRequestFilterDTO.DormName == nil {
		dormNameCondition = "dorm_name LIKE '%%'"
	} else {
		dormNameCondition = "dorm_name LIKE '%" + *roommateRequestFilterDTO.DormName + "%'"
	}

	condition := dormNameCondition + repository.getCondition(roommateRequestFilterDTO)

	repository.db.Preload("RoomPictures").Preload("RoomFacilities").Joins("Student").Where(condition).Find(&roommateRequests)

	return roommateRequests
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
		roomPictures = append(roomPictures, model.RoommateRequestRegisteredDormPicture{PictureUrl: pictureUrl,
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
		roomPictures = append(roomPictures, model.RoommateRequestUnregisteredDormPicture{PictureUrl: pictureUrl,
			RoommateRequestWithUnregisteredDormStudentID: id,
		})
	}

	repository.db.Table("roommate_request_unregistered_dorm_pictures").Where("roommate_request_with_unregistered_dorm_student_id = ?", id).Delete(model.RoommateRequestUnregisteredDormPicture{})

	repository.db.Create(&roomPictures)

	err := repository.db.Preload("RoomPictures").Preload("RoomFacilities").Where("student_id = ?", id).First(&roommateRequestWithUnregisteredDorm).Error

	return roommateRequestWithUnregisteredDorm, err
}

func (repository *roommateRequestRepository) getCondition(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) string {
	zoneCondition := repository.getZoneCondition(roommateRequestFilterDTO.Zone)
	genderCondition := repository.getGenderCondition(roommateRequestFilterDTO.Gender)
	facultyCondition := repository.getFacultyCondition(roommateRequestFilterDTO.Faculty)
	yearOfStudyCondition := repository.getYearOfStudyCondition(roommateRequestFilterDTO.YearOfStudy)
	budgetCondition := repository.getBudgetCondition(roommateRequestFilterDTO.LowerPrice, roommateRequestFilterDTO.UpperPrice)
	roommateCondition := repository.getNumberOfRoommatesCondition(roommateRequestFilterDTO.NumberOfRoommates)
	condition := fmt.Sprintf("%s %s %s %s %s %s", zoneCondition, genderCondition, facultyCondition, yearOfStudyCondition, budgetCondition, roommateCondition)

	fmt.Println(condition)

	return condition
}

func (repository *roommateRequestRepository) getZoneCondition(zone *string) string {
	if zone == nil {
		return ""
	}

	return fmt.Sprintf("AND dorm_zone_name = '%s'", *zone)
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

func (repository *roommateRequestRepository) getBudgetCondition(lowerPrice *int, upperPrice *int) string {
	if lowerPrice == nil || upperPrice == nil || *upperPrice < *lowerPrice {
		return ""
	}

	return fmt.Sprintf("AND shared_room_price BETWEEN %d AND %d", *lowerPrice, *upperPrice)
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
