package service

import (
	"strconv"

	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type RoommateRequestType string

const (
	RoommateRequestWithNoRoom           RoommateRequestType = "NO_ROOM"
	RoommateRequestWithRegisteredDorm   RoommateRequestType = "REGISTERED_DORM"
	RoommateRequestWithUnregisteredDorm RoommateRequestType = "UNREGISTERED_DORM"
)

type RoommateRequestService interface {
	GetRoommateRequestsWithRoom(dto.RoommateRequestFilterDTO) []dto.RoommateRequestWithRoomDTO
	CreateRoommateRequestWithNoRoom(dto.RoommateRequestWithNoRoomDTO) model.RoommateRequestWithNoRoom
	CreateRoommateRequestWithRegisteredDorm(dto.RoommateRequestWithRegisteredDormDTO) model.RoommateRequestWithRegisteredDorm
	CreateRoommateRequestWithUnregisteredDorm(dto.RoommateRequestWithUnregisteredDormDTO) model.RoommateRequestWithUnregisteredDorm
	UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) model.RoommateRequestWithRegisteredDorm
	UpdateRoommateRequestWithUnregisteredDormPictures(id string, pictureUrls []string) model.RoommateRequestWithUnregisteredDorm
	CanUpdateRoommateRequestPicture(studentID string, requestType RoommateRequestType) bool
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

func (roommateRequestService *roommateRequestService) GetRoommateRequestsWithRoom(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []dto.RoommateRequestWithRoomDTO {
	roommateRequestsWithRoom := []dto.RoommateRequestWithRoomDTO{}
	roommateRequestsWithRegisteredDorm := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithRegisteredDorm(roommateRequestFilterDTO)
	roommateRequestWithUnregisteredDorm := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithUnregisteredDorm(roommateRequestFilterDTO)

	//TODO: if no room pictures, use the dorm owner's room picture
	for _, roommateRequest := range roommateRequestsWithRegisteredDorm {
		student := roommateRequestService.studentService.GetStudent(roommateRequest.Student.ID)
		roommateRequestsWithRoom = append(roommateRequestsWithRoom, dto.RoommateRequestWithRoomDTO{
			ID:              roommateRequest.StudentID,
			RoomPicture:     roommateRequest.RoomPictures[0].PictureUrl,
			DormName:        roommateRequest.Dorm.Name,
			Student:         student,
			SharedRoomPrice: roommateRequest.SharedRoomPrice,
		})
	}

	for _, roommateRequest := range roommateRequestWithUnregisteredDorm {
		student := roommateRequestService.studentService.GetStudent(roommateRequest.Student.ID)
		roommateRequestsWithRoom = append(roommateRequestsWithRoom, dto.RoommateRequestWithRoomDTO{
			ID:              roommateRequest.StudentID,
			RoomPicture:     roommateRequest.RoomPictures[0].PictureUrl,
			DormName:        roommateRequest.DormName,
			Student:         student,
			SharedRoomPrice: roommateRequest.SharedRoomPrice,
		})
	}

	return roommateRequestsWithRoom
}

func (roommateRequestService *roommateRequestService) CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoomDTO dto.RoommateRequestWithNoRoomDTO) model.RoommateRequestWithNoRoom {
	if !roommateRequestService.canCreateRoommateRequest(roommateRequestWithNoRoomDTO.StudentID) {
		panic(errortype.ErrOpenRoommateRequest)
	}

	dormZones := []model.DormZone{}

	for _, inputDormZone := range roommateRequestWithNoRoomDTO.Zone {
		dormZones = append(dormZones, model.DormZone{
			Name: inputDormZone,
		})
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
		panic(errortype.ErrOpenRoommateRequest)
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

func (roommateRequestService *roommateRequestService) CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDormDTO dto.RoommateRequestWithUnregisteredDormDTO) model.RoommateRequestWithUnregisteredDorm {
	if !roommateRequestService.canCreateRoommateRequest(roommateRequestWithUnregisteredDormDTO.StudentID) {
		panic(errortype.ErrOpenRoommateRequest)
	}

	roomFacilities := []model.AllRoomFacility{}

	for _, inputRoomFacility := range roommateRequestWithUnregisteredDormDTO.RoomFacilities {
		roomFacility := model.AllRoomFacility{
			Name: inputRoomFacility,
		}

		roomFacilities = append(roomFacilities, roomFacility)
	}

	roommateRequestWithUnregisteredDorm := &model.RoommateRequestWithUnregisteredDorm{
		StudentID:         roommateRequestWithUnregisteredDormDTO.StudentID,
		DormName:          roommateRequestWithUnregisteredDormDTO.DormName,
		DormZoneName:      roommateRequestWithUnregisteredDormDTO.Zone,
		RoomDescription:   roommateRequestWithUnregisteredDormDTO.RoomDescription,
		RoomPrice:         roommateRequestWithUnregisteredDormDTO.RoomPrice,
		RoomSize:          roommateRequestWithUnregisteredDormDTO.RoomSize,
		RoomFacilities:    roomFacilities,
		NumberOfRoommates: roommateRequestWithUnregisteredDormDTO.NumberOfRoommates,
		SharedRoomPrice:   roommateRequestWithUnregisteredDormDTO.SharedRoomPrice,
	}

	createdRoommateRequest, err := roommateRequestService.roommateRequestRepository.CreateRoommateRequestWithUnregisteredDorm(*roommateRequestWithUnregisteredDorm)

	if err != nil {
		panic(err)
	}

	roommateRequestService.studentService.UpdateStudent(roommateRequestWithUnregisteredDorm.StudentID, map[string]interface{}{"roommate_request": RoommateRequestWithUnregisteredDorm})

	return createdRoommateRequest
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestWithRegisteredDormPictures(id string, pictureUrls []string) model.RoommateRequestWithRegisteredDorm {
	updatedRoommateRequestWithRegisteredDorm, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestWithRegisteredDormPictures(id, pictureUrls)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequestWithRegisteredDorm
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestWithUnregisteredDormPictures(id string, pictureUrls []string) model.RoommateRequestWithUnregisteredDorm {
	updatedRoommateRequestWithUnregisteredDorm, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestWithUnregisteredDormPictures(id, pictureUrls)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequestWithUnregisteredDorm
}

func (roommateRequestService *roommateRequestService) CanUpdateRoommateRequestPicture(studentID string, requestType RoommateRequestType) bool {
	student := roommateRequestService.studentService.GetStudent(studentID)

	return *student.RoommateRequest == string(requestType)
}

func (roommateRequestService *roommateRequestService) canCreateRoommateRequest(studentID string) bool {
	student := roommateRequestService.studentService.GetStudent(studentID)

	return student.RoommateRequest == nil
}
