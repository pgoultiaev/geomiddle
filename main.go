package geomiddle

import (
	"encoding/json"
	"math"
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

type location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func init() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleMidpoint).Methods("POST")
	http.Handle("/", router)
}

func handleMidpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var locations []location

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&locations)
	c := appengine.NewContext(r)
	if err != nil {
		http.Error(w, "could not parse json body", http.StatusInternalServerError)
		log.Errorf(c, "could not parse json body: %v", err)
	}

	midPoint, err := getMidPoint(locations)
	if err != nil {
		http.Error(w, "could not get midpoint", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(midPoint)
}

// Calculate geographic midpoint
// http://www.geomidpoint.com/example.html
func getMidPoint(locations []location) (location, error) {

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

	return location{middleLatitude * 180 / math.Pi, middleLongitude * 180 / math.Pi}, nil
}
