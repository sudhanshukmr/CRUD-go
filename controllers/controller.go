package controllers

import (
	"Restapi/entity"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var profiles []entity.Profile = []entity.Profile{}

func AddItem(q http.ResponseWriter, r *http.Request) {
	var newProfile entity.Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(profiles)
}

func GetAllProfiles(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)
}

func GetProfiles(q http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(http.StatusBadRequest)
		q.Write([]byte("Id could not be converted to integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(http.StatusNotFound)
		q.Write([]byte("no profile found with specific ID"))
		return
	}
	profile := profiles[id]
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profile)
}

func UpdateProfiles(q http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("Id could not be converted to integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specific ID"))
		return
	}
	var updateProfiles entity.Profile
	json.NewDecoder(r.Body).Decode(&updateProfiles)
	profiles[id] = updateProfiles
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(updateProfiles)

}

func DeleteProfiles(q http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {

		q.WriteHeader(400)
		q.Write([]byte("Id could not be converted to integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specific ID"))
		return
	}

	profiles = append(profiles[:id], profiles[id+1:]...)
	q.WriteHeader(200)

}
