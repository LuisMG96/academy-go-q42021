package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuisMG96/academy-go-q42021/common"
	"github.com/LuisMG96/academy-go-q42021/services"
	"github.com/gorilla/mux"
)

//GetAllCharacters - Receive a response and a requeset, it's the entry point for retrieve the full list of characters

func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := services.CsvService{}
	//_, data, err := service.ReadFromCSV()
	data, err := service.GetAllCharacters()
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}

//GetCharacterById - Receive a response and a requeset, it's the entry point for retrieve a character by id
func GetCharacterById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := services.NewCsvService()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	}
	data, err := service.GetCharacterById(id)
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(errorResponse.Status)
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}

func WriteCharactersOnCsv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := services.NewCsvService()
	err := service.WriteCharactersOnCSV()
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(errorResponse.Status)
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		response := common.NewResponse(http.StatusCreated, "Success")
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)
	}
}
