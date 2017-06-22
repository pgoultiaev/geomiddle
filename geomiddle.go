package geomiddle

import (
	"math"
)

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// CalculateMidPoint of multiple geo coordinates
// http://www.geomidpoint.com/example.html
func CalculateMidPoint(locations []Location) Location {

	numberOfLocations := float64(len(locations))

	sumCartesianX := 0.0
	sumCartesianY := 0.0
	sumCartesianZ := 0.0

	for _, loc := range locations {
		latInRadians := loc.Lat * math.Pi / 180
		longInRadians := loc.Long * math.Pi / 180

		cartesianX := math.Cos(latInRadians) * math.Cos(longInRadians)
		cartesianY := math.Cos(latInRadians) * math.Sin(longInRadians)
		cartesianZ := math.Sin(latInRadians)

		sumCartesianX += cartesianX
		sumCartesianY += cartesianY
		sumCartesianZ += cartesianZ
	}

	averageCartesianX := sumCartesianX / numberOfLocations
	averageCartesianY := sumCartesianY / numberOfLocations
	averageCartesianZ := sumCartesianZ / numberOfLocations

	middleLongitude := math.Atan2(averageCartesianY, averageCartesianX)
	hyp := math.Sqrt(averageCartesianX*averageCartesianX + averageCartesianY*averageCartesianY)
	middleLatitude := math.Atan2(averageCartesianZ, hyp)

	return Location{middleLatitude * 180 / math.Pi, middleLongitude * 180 / math.Pi}
}
