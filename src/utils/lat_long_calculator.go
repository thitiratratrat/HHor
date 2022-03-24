package utils

import (
	"math"
)

func GetLatLongFromDistance(lat float64, long float64, distance float64, bearingDeg float64) (float64, float64) {
	radius := 6378.1
	pi := math.Pi
	bearing := bearingDeg * (pi / 180)
	currentLatRadian := lat * (pi / 180)
	currentLongRadian := long * (pi / 180)

	destLatRadian := math.Asin(math.Sin(currentLatRadian)*math.Cos(distance/radius) +
		math.Cos(currentLatRadian)*math.Sin(distance/radius)*math.Cos(bearing))

	destLongRadian := currentLongRadian + math.Atan2(math.Sin(bearing)*math.Sin(distance/radius)*math.Cos(currentLatRadian),
		math.Cos(distance/radius)-math.Sin(currentLatRadian)*math.Sin(destLatRadian))

	finalLat := destLatRadian * (180 / pi)
	finalLong := destLongRadian * (180 / pi)

	return finalLat, finalLong
}
