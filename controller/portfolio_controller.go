package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user-crpyto-portfolio/database"
	"user-crpyto-portfolio/entity"

	"github.com/gorilla/mux"
)

func GetAllPortfolios(w http.ResponseWriter, r *http.Request) {
	var portfolios []entity.Portfolio
	database.Connector.Find(&portfolios)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portfolios)
}

//method to fetch user by id
func GetPortfolioByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var portfolio entity.Portfolio
	database.Connector.First(&portfolio, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(portfolio)
}

//method to create user
func CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var portfolio entity.Portfolio
	json.Unmarshal(requestBody, &portfolio)

	database.Connector.Create(portfolio)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(portfolio)
}
