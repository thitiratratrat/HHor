package utils

import (
	"testing"

	"github.com/thitiratratrat/hhor/src/utils"
)

func TestLatLong(t *testing.T) {
	expectedLat := 13.729884598494714
	expectedLong := 100.8244695328533
	originLat := 13.7298889
	originLong := 100.7782323
	distance := 5.0
	bearingDeg := 90.0
	actualLat, actualLong := utils.GetLatLongFromDistance(originLat, originLong, distance, bearingDeg)

	if expectedLat != actualLat || expectedLong != actualLong {
		t.Errorf("got %v, wanted %v", actualLat, expectedLat)
		t.Errorf("got %v, wanted %v", actualLong, expectedLong)
	}
}

func TestDistance(t *testing.T) {
	expectedLat := 13.729871693981964
	expectedLong := 100.8707067623141
	originLat := 13.7298889
	originLong := 100.7782323
	distance := 10.0
	bearingDeg := 90.0
	actualLat, actualLong := utils.GetLatLongFromDistance(originLat, originLong, distance, bearingDeg)

	if expectedLat != actualLat || expectedLong != actualLong {
		t.Errorf("got %v, wanted %v", actualLat, expectedLat)
		t.Errorf("got %v, wanted %v", actualLong, expectedLong)
	}
}

func TestDegree(t *testing.T) {
	expectedLat := 13.684972875233155
	expectedLong := 100.7782323
	originLat := 13.7298889
	originLong := 100.7782323
	distance := 5.0
	bearingDeg := 180.0
	actualLat, actualLong := utils.GetLatLongFromDistance(originLat, originLong, distance, bearingDeg)

	if expectedLat != actualLat || expectedLong != actualLong {
		t.Errorf("got %v, wanted %v", actualLat, expectedLat)
		t.Errorf("got %v, wanted %v", actualLong, expectedLong)
	}
}
