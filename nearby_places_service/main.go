package main

import (
	"fmt"
	"os"

	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/repository"
	"github.com/thitiratratrat/hhor/src/service"
	"github.com/thitiratratrat/hhor/src/utils"
)

var dbConnector utils.DBConnector

func migrateDatabase() {
	dbConnector.GetDB().AutoMigrate(&model.Dorm{}, &model.NearbyLocation{})
}

func init() {
	dbConnector = utils.DBConnectorHandler()

	dbConnector.Open()

	dormRepository := repository.DormRepositoryHandler(dbConnector.GetDB())
	nearbyPlacesService := service.NearbyPlacesHandler()

	args := os.Args

	if len(args) > 1 {
		id := args[1]
		dorm, err := dormRepository.FindDorm(id)

		if err != nil {
			fmt.Println("invalid dorm ID")

			return
		}

		nearbyLocations := nearbyPlacesService.GetNearbyPlaces(dorm.ID, dorm.Latitude, dorm.Longitude)
		dormRepository.UpdateNearbyLocations(id, nearbyLocations)

		return
	}

	var dorms []model.Dorm
	dbConnector.GetDB().Find(&dorms)

	for _, dorm := range dorms {
		nearbyLocations := nearbyPlacesService.GetNearbyPlaces(dorm.ID, dorm.Latitude, dorm.Longitude)
		dormRepository.UpdateNearbyLocations(fmt.Sprint(dorm.ID), nearbyLocations)
	}
}

func main() {
	migrateDatabase()

	defer dbConnector.Close()
}
