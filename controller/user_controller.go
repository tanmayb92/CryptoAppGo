package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user-crpyto-portfolio/database"
	"user-crpyto-portfolio/entity"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []entity.User
	database.Connector.Preload("Portfolios").Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

//method to fetch user by id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputOrderID := params["userID"]

	var user entity.User
	database.Connector.First(&user, inputOrderID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//method to create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestBody, &user)

	database.Connector.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
