package constant

const StudentProfilePictureFolder = "student_profile/"
const PetPicturesFolder = "pet_pictures/"
const RoommateRequestRoomPictureFolder = "roommate_request_room_pictures/"

type RoommateRequestType string

const (
	RoommateRequestWithNoRoom           RoommateRequestType = "NO_ROOM"
	RoommateRequestWithRegisteredDorm   RoommateRequestType = "REGISTERED_DORM"
	RoommateRequestWithUnregisteredDorm RoommateRequestType = "UNREGISTERED_DORM"
)
