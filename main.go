package main

import (
	"coffe-service/database"
	"coffe-service/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDB()
	RouterStart()
}

//
func Services(router *mux.Router) {
	staticDir := "/static/"
	router.PathPrefix(staticDir).Handler(
		http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	router.HandleFunc("/coffe", routers.GetAllCoffe).Methods("GET")
	router.HandleFunc("/coffe", routers.CreateNewCoffe).Methods("POST")
	router.HandleFunc("/coffe", routers.UpdateCoffeByID).Methods("PUT")

}

func RouterStart() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println(`Running on port 3000`)
	Services(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}

//
func initDB() {
	config :=
		database.Config{
			ServerName: "34.132.22.44:3306",
			User:       "root",
			Password:   "",
			DB:         "coffe_db",
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}
	database.Migrate()

}
