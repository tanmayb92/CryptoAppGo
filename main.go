package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controllers "user-crpyto-portfolio/controller"
	"user-crpyto-portfolio/database"
	"user-crpyto-portfolio/entity"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	intializeDbConnection()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	//user api
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserByID).Methods("GET")

	//portfolio api
	router.HandleFunc("/user/{id}/portfolio", controllers.CreatePortfolio).Methods("POST")
	router.HandleFunc("/user/{id}/portfolio", controllers.GetAllPortfolios).Methods("GET")
	router.HandleFunc("/user/{id}/portfolio/{pid}", controllers.GetPortfolioByID).Methods("GET")
}

func intializeDbConnection() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "admin",
			DB:         "mydb",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.MigrateUser(&entity.User{})
	database.MigratePortfolio(&entity.Portfolio{})
}
