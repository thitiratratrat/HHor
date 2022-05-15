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

type CacheNameSpace string

const (
	Dorm              CacheNameSpace = "dorm"
	Room              CacheNameSpace = "room"
	Roommate          CacheNameSpace = "roommate"
	RoommateNoRoom    CacheNameSpace = "roommatenoroom"
	RoommateWithRooms CacheNameSpace = "roommatewithrooms"
	Student           CacheNameSpace = "student"
	DormOwner         CacheNameSpace = "dormowner"
)
