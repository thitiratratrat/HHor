package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/thitiratratrat/hhor/src/model"
)

type NearbyPlacesService interface {
	GetNearbyPlaces(lat float64, long float64) []model.NearbyLocation
}

func NearbyPlacesHandler() NearbyPlacesService {
	return &nearbyPlacesService{}
}

type nearbyPlacesService struct{}

func (nearbyPlacesService *nearbyPlacesService) GetNearbyPlaces(lat float64, long float64) []model.NearbyLocation {
	apiKey := os.Getenv("PLACES_KEY")
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=-%v,%v&radius=1000&key=%s", lat, long, apiKey)
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

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(body)

	return []model.NearbyLocation{}
}
