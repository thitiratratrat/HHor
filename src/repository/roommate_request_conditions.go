package repository

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
)

func getRoomCondition(roommateRequestFilterDTO dto.RoommateRequestRoomFilterDTO, requestType constant.RoommateRequestType) string {
	nameCondition := getNameCondition(roommateRequestFilterDTO.DormName, requestType)
	roommateCondition := getNumberOfRoommatesCondition(roommateRequestFilterDTO.NumberOfRoommates)
	roommFacilityCondition := getRoomFacilityCondition(roommateRequestFilterDTO.RoomFacilities, requestType)
	noRoomCondition := getCondition(roommateRequestFilterDTO.RoommateRequestFilterDTO, requestType)

	condition := fmt.Sprintf("%s %s %s AND %s", nameCondition, roommateCondition, roommFacilityCondition, noRoomCondition)

	return condition
}

func getCondition(roommateRequestFilterDTO dto.RoommateRequestFilterDTO, requestType constant.RoommateRequestType) string {
	zoneCondition := getZoneCondition(roommateRequestFilterDTO.Zone, requestType)
	genderCondition := getGenderCondition(roommateRequestFilterDTO.Gender)
	facultyCondition := getFacultyCondition(roommateRequestFilterDTO.Faculty)
	yearOfStudyCondition := getYearOfStudyCondition(roommateRequestFilterDTO.YearOfStudy)
	budgetCondition := getBudgetCondition(roommateRequestFilterDTO.LowerPrice, roommateRequestFilterDTO.UpperPrice, requestType)
	preferenceCondition := getPreferenceCondition(roommateRequestFilterDTO.SmokeHabitID, roommateRequestFilterDTO.RoomCareHabitID, roommateRequestFilterDTO.SleepHabitID, roommateRequestFilterDTO.StudyHabitID, roommateRequestFilterDTO.PetHabitID)
	condition := fmt.Sprintf("true %s %s %s %s %s %s", zoneCondition, genderCondition, facultyCondition, yearOfStudyCondition, budgetCondition, preferenceCondition)

	return condition
}

func getNameCondition(name *string, requestType constant.RoommateRequestType) string {
	switch {
	case requestType == constant.RoommateRequestRegDorm && name == nil:
		return `"Dorm".name` + " LIKE '%%'"
	case requestType == constant.RoommateRequestRegDorm:
		if len(*name) == 1 {
			return `"Dorm".name` + ` LIKE '` + *name + `%'`
		}

		return fmt.Sprintf("'%s'", *name) + ` % ANY(STRING_TO_ARRAY("Dorm".name,' '))`
	case requestType == constant.RoommateRequestUnregDorm && name == nil:
		return "dorm_name LIKE '%%'"
	default:
		if len(*name) == 1 {
			return "dorm_name" + ` LIKE '` + *name + `%'`
		}

		return fmt.Sprintf("'%s'", *name) + "% ANY(STRING_TO_ARRAY(dorm_name,' '))"
	}
}

func getZoneCondition(zone *string, requestType constant.RoommateRequestType) string {
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

func getGenderCondition(gender []string) string {
	if len(gender) == 0 {
		return ""
	}

	formattedGender := "'" + strings.Join(gender, "', '") + "'"

	return fmt.Sprintf("AND gender_name IN (%s)", formattedGender)
}

func getFacultyCondition(faculty []string) string {
	if len(faculty) == 0 {
		return ""
	}

	formattedFaculty := "'" + strings.Join(faculty, "', '") + "'"

	return fmt.Sprintf("AND faculty_name IN (%s)", formattedFaculty)
}

func getYearOfStudyCondition(yearOfStudy []int) string {
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

func getBudgetCondition(lowerPrice *int, upperPrice *int, requestType constant.RoommateRequestType) string {
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

func getNumberOfRoommatesCondition(numberOfRoomates []int) string {
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

func getRoomFacilityCondition(roomFacilities []string, requestType constant.RoommateRequestType) string {
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

func getPreferenceCondition(smokeHabit *string, roomCareHabit *string, sleepHabit *string, studyHabit *string, petHabit *string) string {
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
