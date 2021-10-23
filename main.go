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
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func setRoutes() {
	router.InitRoutes(server, controllers)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func migrateDatabase() {
	dbConnector.GetDB().AutoMigrate(&model.AllDormFacility{}, &model.Account{}, &model.AllRoomFacility{}, &model.DormZone{}, &model.Location{}, &model.Dorm{}, &model.NearbyLocation{}, &model.Room{}, &model.RoomPicture{}, &model.DormPicture{})
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

	dormService := service.DormServiceHandler(dormRepository, dormFacilityRepository, dormZoneRepository)
	roomService := service.RoomServiceHandler(roomFacilityRepository)

	dormController := controller.DormControllerHandler(dormService)
	roomController := controller.RoomControllerHandler(roomService)

	controllers = router.Controllers{
		DormController: dormController,
		RoomController: roomController,
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
