package controllers

import (
	"encoding/json"
	"github.com/LuisMG96/academy-go-q42021/common"
	"github.com/LuisMG96/academy-go-q42021/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := services.CsvService{}
	//_, data, err := service.ReadFromCSV()
	data, err := service.GetAllCharacters()
	if err != nil {
		errorResponse := common.New(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}

func GetCharacterById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := services.NewCsvService()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorResponse := common.New(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	}
	data, err := service.GetCharacterById(id)
	if err != nil {
		errorResponse := common.New(err)
		w.WriteHeader(errorResponse.Status)
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}
