package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pgoultiaev/geomiddle"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleMidpoint).Methods("POST")
	http.Handle("/", router)
}

func handleMidpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var locations []Location

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err := decoder.Decode(&locations)
	c := appengine.NewContext(r)
	if err != nil {
		http.Error(w, "could not parse json body", http.StatusInternalServerError)
		log.Errorf(c, "could not parse json body: %v", err)
	}

	midPoint := geomiddle.CaculateMidPoint(locations)
	if err != nil {
		http.Error(w, "could not get midpoint", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(midPoint)
}
