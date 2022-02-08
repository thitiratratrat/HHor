package service

import (
	"errors"
	"strconv"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"gorm.io/gorm"
)

const (
	RoommateRequestWithNoRoom         string = "NO_ROOM"
	RoommateRequestWithRegisteredDorm string = "REGISTERED_DORM"
)

type RoommateRequestService interface {
	CreateRoommateRequestWithNoRoom(dto.RoommateRequestWithNoRoomDTO) model.RoommateRequestWithNoRoom
	CreateRoommateRequestWithRegisteredDorm(dto.RoommateRequestWithRegisteredDormDTO) model.RoommateRequestWithRegisteredDorm
	UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) model.RoommateRequestWithRegisteredDorm
}

func RoommateRequestServiceHandler(roommateRequestRepository repository.RoommateRequestRepository, studentSerivce StudentService) RoommateRequestService {
	return &roommateRequestService{
		roommateRequestRepository: roommateRequestRepository,
		studentService:            studentSerivce,
	}
}

type roommateRequestService struct {
	roommateRequestRepository repository.RoommateRequestRepository
	studentService            StudentService
}

func (roommateRequestService *roommateRequestService) CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoomDTO dto.RoommateRequestWithNoRoomDTO) model.RoommateRequestWithNoRoom {
	if !roommateRequestService.canCreateRoommateRequest(roommateRequestWithNoRoomDTO.StudentID) {
		panic(errors.New("student already has an open roommate request"))
	}

	dormZones := []model.DormZone{}

	for _, inputDormZone := range roommateRequestWithNoRoomDTO.Zone {
		dormZone := model.DormZone{
			Name: inputDormZone,
		}

		dormZones = append(dormZones, dormZone)
	}

	roommateRequestWithNoRoom := &model.RoommateRequestWithNoRoom{
		StudentID: roommateRequestWithNoRoomDTO.StudentID,
		Budget:    roommateRequestWithNoRoomDTO.Budget,
		Zones:     dormZones,
	}

	createdRoommateRequest, err := roommateRequestService.roommateRequestRepository.CreateRoommateRequestWithNoRoom(*roommateRequestWithNoRoom)

	if err != nil {
		panic(err)
	}

	roommateRequestService.studentService.UpdateStudent(roommateRequestWithNoRoomDTO.StudentID, map[string]interface{}{"roommate_request": RoommateRequestWithNoRoom})

	return createdRoommateRequest
}

func (roommateRequestService *roommateRequestService) CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDormDTO dto.RoommateRequestWithRegisteredDormDTO) model.RoommateRequestWithRegisteredDorm {
	if !roommateRequestService.canCreateRoommateRequest(roommateRequestWithRegisteredDormDTO.StudentID) {
		panic(errors.New("student already has an open roommate request"))
	}

	roomID, _ := strconv.Atoi(roommateRequestWithRegisteredDormDTO.RoomID)
	dormID, _ := strconv.Atoi(roommateRequestWithRegisteredDormDTO.DormID)

	roommateRequestWithRegisteredDorm := &model.RoommateRequestWithRegisteredDorm{
		StudentID:         roommateRequestWithRegisteredDormDTO.StudentID,
		SharedRoomPrice:   roommateRequestWithRegisteredDormDTO.SharedRoomPrice,
		NumberOfRoommates: roommateRequestWithRegisteredDormDTO.NumberOfRoommates,
		RoomID:            uint(roomID),
		DormID:            uint(dormID),
	}

	createdRoommateRequest, err := roommateRequestService.roommateRequestRepository.CreateRoommateRequestWithRegisteredDorm(*roommateRequestWithRegisteredDorm)

	if err != nil {
		panic(err)
	}

	roommateRequestService.studentService.UpdateStudent(roommateRequestWithRegisteredDormDTO.StudentID, map[string]interface{}{"roommate_request": RoommateRequestWithRegisteredDorm})

	return createdRoommateRequest
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) model.RoommateRequestWithRegisteredDorm {
	updatedRoommateRequestWithRegisteredDorm, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestWithRegisteredDormPictures(id, pictureUrls)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequestWithRegisteredDorm
}

func (roommateRequestService *roommateRequestService) canCreateRoommateRequest(studentID string) bool {
	student, err := roommateRequestService.studentService.GetStudent(studentID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(errortype.ErrResourceNotFound)
	} else if err != nil {
		panic(err)
	}

	return student.RoommateRequest == nil
}
