package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/thitiratratrat/hhor/src/model"
	"github.com/thitiratratrat/hhor/src/utils"
)

type NearbyPlacesService interface {
	GetNearbyPlaces(dormID uint, lat float64, long float64) []model.NearbyLocation
}

func NearbyPlacesHandler() NearbyPlacesService {
	return &nearbyPlacesService{}
}

type nearbyPlacesService struct{}

func (nearbyPlacesService *nearbyPlacesService) GetNearbyPlaces(dormID uint, lat float64, long float64) []model.NearbyLocation {
	var nearbyLocations []model.NearbyLocation
	apiKey := os.Getenv("PLACES_KEY")
	urlUniversity := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%v,%v&radius=500&type=university&rankby=prominence&key=%s", lat, long, apiKey)
	urlStore := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%v,%v&radius=300&type=store&rankby=prominence&key=%s", lat, long, apiKey)

	nearbyLocations = append(nearbyLocations, nearbyPlacesService.findNearbyPlaces(dormID, lat, long, urlUniversity)...)
	nearbyLocations = append(nearbyLocations, nearbyPlacesService.findNearbyPlaces(dormID, lat, long, urlStore)...)

	return nearbyLocations
}

func (nearbyPlacesService) findNearbyPlaces(dormID uint, lat float64, long float64, url string) []model.NearbyLocation {
	limit := 5
	distanceLimit := 600.0
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		panic(err)
	}

	var results []model.NearbyLocation
	nearbyLocations := data["results"].([]interface{})

	for index, nearbyLocation := range nearbyLocations {
		if index == limit {
			break
		}

		nearbyLocationMap := nearbyLocation.(map[string]interface{})
		location := nearbyLocationMap["geometry"].(map[string]interface{})["location"].(map[string]interface{})
		locationLat := location["lat"].(float64)
		locationLong := location["lng"].(float64)
		distance := utils.GetDistanceFromLatLong(lat, long, locationLat, locationLong)

		if distance > distanceLimit {
			continue
		}

		nearby := model.NearbyLocation{
			DormID:    int(dormID),
			Name:      nearbyLocationMap["name"].(string),
			Longitude: locationLong,
			Latitude:  locationLat,
			Distance:  distance,
		}

		results = append(results, nearby)
	}

	return results
}
