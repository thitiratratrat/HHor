package constant

const StudentProfilePictureFolder = "student_profile/"
const DormOwnerProfilePictureFolder = "dorm_owner_profile/"
const PetPicturesFolder = "pet_pictures/"
const BankQRPictureFolder = "bank_qr/"
const RoommateRequestRoomPictureFolder = "roommate_request_room_pictures/"
const DormPictureFolder = "dorm_pictures/"
const RoomPictureFolder = "room_pictures/"

type RoommateRequestType string

const (
	RoommateRequestNoRoom    RoommateRequestType = "NO_ROOM"
	RoommateRequestRegDorm   RoommateRequestType = "REGISTERED_DORM"
	RoommateRequestUnregDorm RoommateRequestType = "UNREGISTERED_DORM"
)
