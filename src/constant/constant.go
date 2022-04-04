package constant

const StudentProfilePictureFolder = "student_profile/"
const PetPicturesFolder = "pet_pictures/"
const RoommateRequestRoomPictureFolder = "roommate_request_room_pictures/"
const DormPictureFolder = "dorm_pictures/"
const RoomPictureFolder = "room_pictures/"

type RoommateRequestType string

const (
	RoommateRequestNoRoom    RoommateRequestType = "NO_ROOM"
	RoommateRequestRegDorm   RoommateRequestType = "REGISTERED_DORM"
	RoommateRequestUnregDorm RoommateRequestType = "UNREGISTERED_DORM"
)
