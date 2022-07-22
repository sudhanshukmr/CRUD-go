package main

import (
	"Restapi/controllers"
	"Restapi/database"
	"Restapi/entity"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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

}
