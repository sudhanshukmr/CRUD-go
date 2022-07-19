package main

import (
	"Restapi/controllers"
	"Restapi/database"
	"Restapi/entity"
	"log"
	"net/http"

	// "net/http"
	"fmt"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var profiles []entity.Profile = []entity.Profile{}

/*
// type User struct {
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// 	Email     string `json:"email"`
// }
// type Profile struct {
// 	Department  string `json:"department"`
// 	Designation string `json:"designation"`
// 	Employee    User   `json:"employee"`
// }

func addItem(q http.ResponseWriter, r *http.Request) {
	var newProfile entity.Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(profiles)
}

func getAllProfiles(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)
}

func getProfiles(q http.ResponseWriter, r *http.Request) {

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

func updateProfiles(q http.ResponseWriter, r *http.Request) {

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

func deleteProfiles(q http.ResponseWriter, r *http.Request) {

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
*/
func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))

}
func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "root",
			DB:         "employee",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Profile{})
}
func initaliseHandlers(router *mux.Router) {

	fmt.Println("In main.go , router")
	router.HandleFunc("/profiles", controllers.AddItem).Methods("POST")

	router.HandleFunc("/profiles", controllers.GetAllProfiles).Methods("GET")

	router.HandleFunc("/profiles/{id}", controllers.GetProfiles).Methods("GET")

	router.HandleFunc("/profiles/{id}", controllers.UpdateProfiles).Methods("PUT")

	router.HandleFunc("/profiles/{id}", controllers.DeleteProfiles).Methods("DELETE")

	// http.ListenAndServe(":5001", router)
}
