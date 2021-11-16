package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/thitiratratrat/hhor/docs"
	"github.com/thitiratratrat/hhor/src/controller"
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
	dbConnector.GetDB().AutoMigrate(&model.AllDormFacility{}, &model.DormOwner{}, &model.AllRoomFacility{}, &model.DormZone{}, &model.Location{}, &model.Dorm{}, &model.NearbyLocation{}, &model.Room{}, &model.RoomPicture{}, &model.DormPicture{}, &model.PetHabit{}, &model.SmokeHabit{}, &model.StudyHabit{}, &model.RoomCareHabit{}, &model.Gender{}, &model.Faculty{}, &model.Student{})
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
	dormFacilityRepository := repository.DormFacilityRepositoryHandler(dbConnector.GetDB())
	dormZoneRepository := repository.DormZoneRepositoryHandler(dbConnector.GetDB())
	roomFacilityRepository := repository.RoomFacilityRepositoryHandler(dbConnector.GetDB())
	facultyRepository := repository.FacultyRepositoryHandler(dbConnector.GetDB())
	dormOwnerRepository := repository.DormOwnerRepositoryHandler(dbConnector.GetDB())
	studentRepository := repository.StudentRepositoryHandler(dbConnector.GetDB())

	dormService := service.DormServiceHandler(dormRepository, dormFacilityRepository, dormZoneRepository)
	roomService := service.RoomServiceHandler(roomFacilityRepository)
	authService := service.AuthServiceHandler(facultyRepository, studentRepository, dormOwnerRepository)
	studentService := service.StudentServiceHandler(studentRepository)

	dormController := controller.DormControllerHandler(dormService)
	roomController := controller.RoomControllerHandler(roomService)
	authController := controller.AuthControllerHandler(authService, studentService)
	studentController := controller.StudentControllerHandler(studentService)

	controllers = router.Controllers{
		DormController:    dormController,
		RoomController:    roomController,
		AuthController:    authController,
		StudentController: studentController,
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
