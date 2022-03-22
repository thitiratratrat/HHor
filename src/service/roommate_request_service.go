package service

import (
	"strconv"

	"github.com/thitiratratrat/hhor/src/constant"
	"github.com/thitiratratrat/hhor/src/dto"
	"github.com/thitiratratrat/hhor/src/errortype"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
)

type RoommateRequestService interface {
	GetRoommateRequest(id string) dto.RoommateRequestDTO
	GetRoommateRequestsWithRoom(dto.RoommateRequestRoomFilterDTO) []dto.RoommateRequestWithRoomDTO
	GetRoommateRequestsNoRoom(dto.RoommateRequestFilterDTO) []model.RoommateRequestWithNoRoom
	CreateRoommateRequestNoRoom(dto.RoommateRequestNoRoomDTO) model.RoommateRequestWithNoRoom
	CreateRoommateRequestRegDorm(dto.RoommateRequestRegDormDTO) model.RoommateRequestWithRegisteredDorm
	CreateRoommateRequestUnregDorm(dto.RoommateRequestUnregDormDTO) model.RoommateRequestWithUnregisteredDorm
	UpdateRoommateRequestRegDormPictures(id string, pictureUrls []string) model.RoommateRequestWithRegisteredDorm
	UpdateRoommateRequestUnregDormPictures(id string, pictureUrls []string) model.RoommateRequestWithUnregisteredDorm
	UpdateRoommateRequestRegDorm(roommateRequest dto.RoommateRequestRegDormDTO) model.RoommateRequestWithRegisteredDorm
	UpdateRoommateRequestUnregDorm(roommateRequest dto.RoommateRequestUnregDormDTO) model.RoommateRequestWithUnregisteredDorm
	UpdateRoommateRequestNoRoom(roommateRequest dto.RoommateRequestNoRoomDTO) model.RoommateRequestWithNoRoom
	DeleteRoommateRequest(id string)
	CanUpdateRoommateRequest(studentID string, requestType constant.RoommateRequestType) bool
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

func (roommateRequestService *roommateRequestService) GetRoommateRequest(id string) dto.RoommateRequestDTO {
	student := roommateRequestService.studentService.GetStudent(id)
	var roommateRequest dto.RoommateRequestDTO
	roommateRequest.Student = student

	if student.RoommateRequest == nil {
		panic(errortype.ErrNoRoommateRequest)
	}

	switch *student.RoommateRequest {
	case string(constant.RoommateRequestNoRoom):
		noRoomReq, err := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithNoRoom(id)

		if err != nil {
			panic(err)
		}

		roommateRequest.Type = constant.RoommateRequestNoRoom
		roommateRequest.Budget = &noRoomReq.Budget
		roommateRequest.InterestedDormZones = &noRoomReq.Zones
	case string(constant.RoommateRequestUnregDorm):
		unregDormReq, err := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithUnregisteredDorm(id)

		if err != nil {
			panic(err)
		}

		roomPictures := []model.Picture{}
		for _, roomPicture := range unregDormReq.RoomPictures {
			roomPictures = append(roomPictures, model.Picture{
				PictureUrl: roomPicture.PictureUrl,
			})
		}
		roommateRequest.Type = constant.RoommateRequestUnregDorm
		roommateRequest.Dorm = &dto.Dorm{
			Name: unregDormReq.DormName,
			Zone: unregDormReq.DormZoneName,
		}
		roommateRequest.Room = &dto.Room{
			Description:       unregDormReq.RoomDescription,
			Price:             unregDormReq.RoomPrice,
			Size:              unregDormReq.RoomSize,
			Facilities:        unregDormReq.RoomFacilities,
			NumberOfRoommates: unregDormReq.NumberOfRoommates,
			SharedRoomPrice:   unregDormReq.SharedRoomPrice,
			Pictures:          roomPictures,
		}
	case string(constant.RoommateRequestRegDorm):
		regDormReq, err := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithRegisteredDorm(id)

		if err != nil {
			panic(err)
		}

		roomPictures := []model.Picture{}
		if len(regDormReq.RoomPictures) != 0 {
			for _, roomPicture := range regDormReq.RoomPictures {
				roomPictures = append(roomPictures, model.Picture{
					PictureUrl: roomPicture.PictureUrl,
				})
			}
		} else {
			for _, roomPicture := range regDormReq.Room.Pictures {
				roomPictures = append(roomPictures, model.Picture{
					PictureUrl: roomPicture.PictureUrl,
				})
			}
		}

		roommateRequest.Type = constant.RoommateRequestRegDorm
		roommateRequest.Dorm = &dto.Dorm{
			ID:   &regDormReq.Dorm.ID,
			Name: regDormReq.Dorm.Name,
			Zone: regDormReq.Dorm.DormZoneName,
		}
		roommateRequest.Room = &dto.Room{
			ID:                &regDormReq.RoomID,
			Name:              &regDormReq.Room.Name,
			Description:       regDormReq.Room.Description,
			Price:             regDormReq.Room.Price,
			Size:              regDormReq.Room.Size,
			Facilities:        regDormReq.Room.Facilities,
			NumberOfRoommates: regDormReq.NumberOfRoommates,
			SharedRoomPrice:   regDormReq.SharedRoomPrice,
			Pictures:          roomPictures,
		}
	}

	return roommateRequest
}

func (roommateRequestService *roommateRequestService) GetRoommateRequestsWithRoom(roommateRequestRoomFilterDTO dto.RoommateRequestRoomFilterDTO) []dto.RoommateRequestWithRoomDTO {
	roommateRequestsWithRoom := []dto.RoommateRequestWithRoomDTO{}
	roommateRequestsWithRegisteredDorm := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithRegisteredDorms(roommateRequestRoomFilterDTO)
	roommateRequestWithUnregisteredDorm := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithUnregisteredDorms(roommateRequestRoomFilterDTO)

	for _, roommateRequest := range roommateRequestsWithRegisteredDorm {
		student := roommateRequestService.studentService.GetStudent(roommateRequest.Student.ID)
		var roomPicture *string

		if len(roommateRequest.RoomPictures) != 0 {
			roomPicture = &roommateRequest.RoomPictures[0].PictureUrl
		} else if len(roommateRequest.Room.Pictures) != 0 {
			roomPicture = &roommateRequest.Room.Pictures[0].PictureUrl
		}

		roommateRequestsWithRoom = append(roommateRequestsWithRoom, dto.RoommateRequestWithRoomDTO{
			ID:              roommateRequest.StudentID,
			RoomPicture:     roomPicture,
			DormName:        roommateRequest.Dorm.Name,
			Student:         student,
			SharedRoomPrice: roommateRequest.SharedRoomPrice,
			Latitude:        &roommateRequest.Dorm.Latitude,
			Longitude:       &roommateRequest.Dorm.Longitude,
		})
	}

	for _, roommateRequest := range roommateRequestWithUnregisteredDorm {
		student := roommateRequestService.studentService.GetStudent(roommateRequest.Student.ID)
		var roomPicture *string

		if len(roommateRequest.RoomPictures) != 0 {
			roomPicture = &roommateRequest.RoomPictures[0].PictureUrl
		}

		roommateRequestsWithRoom = append(roommateRequestsWithRoom, dto.RoommateRequestWithRoomDTO{
			ID:              roommateRequest.StudentID,
			RoomPicture:     roomPicture,
			DormName:        roommateRequest.DormName,
			Student:         student,
			SharedRoomPrice: roommateRequest.SharedRoomPrice,
		})
	}

	return roommateRequestsWithRoom
}

func (roommateRequestService *roommateRequestService) GetRoommateRequestsNoRoom(roommateRequestFilterDTO dto.RoommateRequestFilterDTO) []model.RoommateRequestWithNoRoom {
	roommateRequests := roommateRequestService.roommateRequestRepository.FindRoommateRequestWithNoRooms(roommateRequestFilterDTO)
	result := make([]model.RoommateRequestWithNoRoom, len(roommateRequests))

	for index, roommateRequest := range roommateRequests {
		roommateRequest.Student = roommateRequestService.studentService.GetStudent(roommateRequest.Student.ID)
		result[index] = roommateRequest
	}

	return result
}

func (roommateRequestService *roommateRequestService) CreateRoommateRequestNoRoom(roommateRequestWithNoRoomDTO dto.RoommateRequestNoRoomDTO) model.RoommateRequestWithNoRoom {
	if !roommateRequestService.canCreateRoommateRequest(roommateRequestWithNoRoomDTO.StudentID) {
		panic(errortype.ErrOpenRoommateRequest)
	}

	roommateRequestWithNoRoom := roommateRequestService.mapRoommateRequestNoRoomDTO(roommateRequestWithNoRoomDTO)
	createdRoommateRequest, err := roommateRequestService.roommateRequestRepository.CreateRoommateRequestWithNoRoom(roommateRequestWithNoRoom)

	if err != nil {
		panic(err)
	}

	roommateRequestService.studentService.UpdateStudent(roommateRequestWithNoRoomDTO.StudentID, map[string]interface{}{"roommate_request": constant.RoommateRequestNoRoom})

	return createdRoommateRequest
}

func (roommateRequestService *roommateRequestService) CreateRoommateRequestRegDorm(roommateRequestWithRegisteredDormDTO dto.RoommateRequestRegDormDTO) model.RoommateRequestWithRegisteredDorm {
	if !roommateRequestService.canCreateRoommateRequest(roommateRequestWithRegisteredDormDTO.StudentID) {
		panic(errortype.ErrOpenRoommateRequest)
	}

	roommateRequestWithRegisteredDorm := roommateRequestService.mapRoommateRequestRegDormDTO(roommateRequestWithRegisteredDormDTO)
	createdRoommateRequest, err := roommateRequestService.roommateRequestRepository.CreateRoommateRequestWithRegisteredDorm(roommateRequestWithRegisteredDorm)

	if err != nil {
		panic(err)
	}

	roommateRequestService.studentService.UpdateStudent(roommateRequestWithRegisteredDormDTO.StudentID, map[string]interface{}{"roommate_request": constant.RoommateRequestRegDorm})

	return createdRoommateRequest
}

func (roommateRequestService *roommateRequestService) CreateRoommateRequestUnregDorm(roommateRequestWithUnregisteredDormDTO dto.RoommateRequestUnregDormDTO) model.RoommateRequestWithUnregisteredDorm {
	if !roommateRequestService.canCreateRoommateRequest(roommateRequestWithUnregisteredDormDTO.StudentID) {
		panic(errortype.ErrOpenRoommateRequest)
	}

	roommateRequestWithUnregisteredDorm := roommateRequestService.mapRoommateRequestUnregDormDTO(roommateRequestWithUnregisteredDormDTO)
	createdRoommateRequest, err := roommateRequestService.roommateRequestRepository.CreateRoommateRequestWithUnregisteredDorm(roommateRequestWithUnregisteredDorm)

	if err != nil {
		panic(err)
	}

	roommateRequestService.studentService.UpdateStudent(roommateRequestWithUnregisteredDormDTO.StudentID, map[string]interface{}{"roommate_request": constant.RoommateRequestUnregDorm})

	return createdRoommateRequest
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestRegDormPictures(id string, pictureUrls []string) model.RoommateRequestWithRegisteredDorm {
	updatedRoommateRequestWithRegisteredDorm, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestWithRegisteredDormPictures(id, pictureUrls)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequestWithRegisteredDorm
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestUnregDormPictures(id string, pictureUrls []string) model.RoommateRequestWithUnregisteredDorm {
	updatedRoommateRequestWithUnregisteredDorm, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestWithUnregisteredDormPictures(id, pictureUrls)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequestWithUnregisteredDorm
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestRegDorm(roommateRequest dto.RoommateRequestRegDormDTO) model.RoommateRequestWithRegisteredDorm {
	update := roommateRequestService.mapRoommateRequestRegDormDTO(roommateRequest)
	updatedRoommateRequest, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestRegDorm(update.StudentID, update)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequest
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestUnregDorm(roommateRequest dto.RoommateRequestUnregDormDTO) model.RoommateRequestWithUnregisteredDorm {
	update := roommateRequestService.mapRoommateRequestUnregDormDTO(roommateRequest)
	updatedRoommateRequest, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestUnregDorm(update.StudentID, update)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequest
}

func (roommateRequestService *roommateRequestService) UpdateRoommateRequestNoRoom(roommateRequest dto.RoommateRequestNoRoomDTO) model.RoommateRequestWithNoRoom {
	update := roommateRequestService.mapRoommateRequestNoRoomDTO(roommateRequest)
	updatedRoommateRequest, err := roommateRequestService.roommateRequestRepository.UpdateRoommateRequestNoRoom(update.StudentID, update)

	if err != nil {
		panic(err)
	}

	return updatedRoommateRequest
}

func (roommateRequestService *roommateRequestService) DeleteRoommateRequest(id string) {
	student := roommateRequestService.studentService.GetStudent(id)
	var roommateRequest dto.RoommateRequestDTO
	var err error
	roommateRequest.Student = student

	if student.RoommateRequest == nil {
		panic(errortype.ErrNoRoommateRequest)
	}

	switch *student.RoommateRequest {
	case string(constant.RoommateRequestRegDorm):
		err = roommateRequestService.roommateRequestRepository.DeleteRoommateRequestRegDorm(id)
	case string(constant.RoommateRequestUnregDorm):
		err = roommateRequestService.roommateRequestRepository.DeleteRoommateRequestUnregDorm(id)
	case string(constant.RoommateRequestNoRoom):
		err = roommateRequestService.roommateRequestRepository.DeleteRoommateRequestNoRoom(id)
	}

	roommateRequestService.studentService.UpdateStudent(id, map[string]interface{}{"roommate_request": nil})

	if err != nil {
		panic(err)
	}
}

func (roommateRequestService *roommateRequestService) CanUpdateRoommateRequest(studentID string, requestType constant.RoommateRequestType) bool {
	student := roommateRequestService.studentService.GetStudent(studentID)

	return *student.RoommateRequest == string(requestType)
}

func (roommateRequestService *roommateRequestService) canCreateRoommateRequest(studentID string) bool {
	student := roommateRequestService.studentService.GetStudent(studentID)

	return student.RoommateRequest == nil
}

func (roommateRequestService *roommateRequestService) mapRoommateRequestNoRoomDTO(roommateRequestNoRoomDTO dto.RoommateRequestNoRoomDTO) model.RoommateRequestWithNoRoom {
	dormZones := []model.DormZone{}

	for _, inputDormZone := range roommateRequestNoRoomDTO.Zone {
		dormZones = append(dormZones, model.DormZone{
			Name: inputDormZone,
		})
	}

	return model.RoommateRequestWithNoRoom{
		StudentID: roommateRequestNoRoomDTO.StudentID,
		Budget:    roommateRequestNoRoomDTO.Budget,
		Zones:     dormZones,
	}
}

func (roommateRequestService *roommateRequestService) mapRoommateRequestRegDormDTO(roommateRequestWithRegisteredDormDTO dto.RoommateRequestRegDormDTO) model.RoommateRequestWithRegisteredDorm {
	roomID, _ := strconv.Atoi(roommateRequestWithRegisteredDormDTO.RoomID)
	dormID, _ := strconv.Atoi(roommateRequestWithRegisteredDormDTO.DormID)

	return model.RoommateRequestWithRegisteredDorm{
		StudentID:         roommateRequestWithRegisteredDormDTO.StudentID,
		SharedRoomPrice:   roommateRequestWithRegisteredDormDTO.SharedRoomPrice,
		NumberOfRoommates: roommateRequestWithRegisteredDormDTO.NumberOfRoommates,
		RoomID:            uint(roomID),
		DormID:            uint(dormID),
	}
}

func (roommateRequestService *roommateRequestService) mapRoommateRequestUnregDormDTO(roommateRequestWithUnregisteredDormDTO dto.RoommateRequestUnregDormDTO) model.RoommateRequestWithUnregisteredDorm {
	roomFacilities := []model.AllRoomFacility{}

	for _, inputRoomFacility := range roommateRequestWithUnregisteredDormDTO.RoomFacilities {
		roomFacility := model.AllRoomFacility{
			Name: inputRoomFacility,
		}

		roomFacilities = append(roomFacilities, roomFacility)
	}

	return model.RoommateRequestWithUnregisteredDorm{
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
}
