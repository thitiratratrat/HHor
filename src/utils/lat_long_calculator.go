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

func GetDistanceFromLatLong(lat1, lon1, lat2, lon2 float64) float64 {
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return math.Floor(2*r*math.Asin(math.Sqrt(h))*100) / 100
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
