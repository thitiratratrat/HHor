package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/thitiratratrat/hhor/docs"
	"github.com/thitiratratrat/hhor/src/controller"
	"github.com/thitiratratrat/hhor/src/fieldvalidator"
	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"github.com/thitiratratrat/hhor/src/router"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

var server *gin.Engine
var dbConnector utils.DBConnector
var controllers router.Controllers

func setUpSwagger() {
	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	docs.SwaggerInfo.Title = "HHor API"
	docs.SwaggerInfo.Description = "HHor is a mobile application that helps students to find nearby dormitories around the university and roommates by allowing the students to search for interested dorms, view dorm’s details, chat with dorm owners, and reserve the room."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
}

func setRoutes() {
	router.InitRoutes(server, controllers)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func migrateDatabase() {
	dbConnector.GetDB().AutoMigrate(&model.AllDormFacility{}, &model.DormOwner{}, &model.AllRoomFacility{}, &model.DormZone{}, &model.Dorm{}, &model.NearbyLocation{}, &model.Room{}, &model.RoomPicture{}, &model.DormPicture{}, &model.PetHabit{}, &model.SmokeHabit{}, &model.StudyHabit{}, &model.RoomCareHabit{}, &model.Gender{}, &model.Faculty{}, &model.Student{}, &model.PetPicture{}, &model.RoommateRequestWithNoRoom{}, &model.RoommateRequestWithRegisteredDorm{}, &model.RoommateRequestRegisteredDormPicture{}, &model.RoommateRequestWithUnregisteredDorm{}, &model.RoommateRequestUnregisteredDormPicture{})

}

func init() {
	server = gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	dbConnector = utils.DBConnectorHandler()

	dbConnector.Open()

	dormRepository := repository.DormRepositoryHandler(dbConnector.GetDB())
	roomRepository := repository.RoomRepositoryHandler(dbConnector.GetDB())
	dormOwnerRepository := repository.DormOwnerRepositoryHandler(dbConnector.GetDB())
	studentRepository := repository.StudentRepositoryHandler(dbConnector.GetDB())
	roommateRequestRpository := repository.RoommateRequestRepositoryHandler(dbConnector.GetDB())

	encryptor := utils.EncryptorHandler()

	jwtService := service.JWTServiceHandler()
	nearbyPlacesService := service.NearbyPlacesHandler()
	dormService := service.DormServiceHandler(nearbyPlacesService, dormRepository, roomRepository, dormOwnerRepository)
	roomService := service.RoomServiceHandler(dormRepository, roomRepository)
	authService := service.AuthServiceHandler(studentRepository, dormOwnerRepository)
	studentService := service.StudentServiceHandler(studentRepository)
	roommateRequestService := service.RoommateRequestServiceHandler(roommateRequestRpository, studentService)
	dormOwnerService := service.DormOwnerServiceHandler(dormOwnerRepository, encryptor)

	fieldValidator := fieldvalidator.FieldValidatorHandler(dormService, roomService, studentService)

	dormController := controller.DormControllerHandler(dormService, jwtService, fieldValidator)
	roomController := controller.RoomControllerHandler(roomService, jwtService, fieldValidator)
	authController := controller.AuthControllerHandler(authService, jwtService, fieldValidator)
	studentController := controller.StudentControllerHandler(studentService, fieldValidator)
	roommateRequestController := controller.RoommateRequestControllerHandler(roommateRequestService, dormService, roomService, fieldValidator)
	dormOwnerController := controller.DormOwnerControllerHandler(dormOwnerService, fieldValidator)

	controllers = router.Controllers{
		DormController:            dormController,
		RoomController:            roomController,
		AuthController:            authController,
		StudentController:         studentController,
		RoommateRequestController: roommateRequestController,
		DormOwnerController:       dormOwnerController,
	}
}

func main() {
	setUpSwagger()
	setRoutes()
	migrateDatabase()

	port := os.Getenv("PORT")

	defer dbConnector.Close()

	server.Run(":" + port)
}
