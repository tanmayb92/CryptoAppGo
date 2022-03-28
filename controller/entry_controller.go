package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-crpyto-portfolio/database"
	"user-crpyto-portfolio/entity"

	"github.com/gorilla/mux"
)

func GetEnteries(w http.ResponseWriter, r *http.Request) {
	var enteries []entity.Entry
	database.Connector.Find(&enteries)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enteries)
}

func GetEntryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	entryID := params["entryID"]

	var entry entity.Entry
	database.Connector.First(&entry, entryID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entry)
}

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["pid"]
	requestBody, _ := ioutil.ReadAll(r.Body)
	var entry entity.Entry
	json.Unmarshal(requestBody, &entry)
	val_uint, _ := strconv.ParseUint(key, 10, 64)
	entry.SetPortfolioId(uint(val_uint))
	database.Connector.Create(entry)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entry)
}
