package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/LuisMG96/academy-go-q42021/common"
	"github.com/LuisMG96/academy-go-q42021/services"
	"github.com/gorilla/mux"
)

//GetAllCharacters - Receive a response and a requeset, it's the entry point for retrieve the full list of characters
func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	common.InfoLogger.Println("CsvController | GetAllCharacters request")
	w.Header().Set("Content-Type", "application/json")
	authService := services.NewAuthService()
	token, err := authService.ExtractTokenFromRequest(r)
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	} else if isValid, err := authService.ValidateToken(token); isValid {
		service := services.CsvService{}
		//_, data, err := service.ReadFromCSV()
		data, err := service.GetAllCharacters()
		if err != nil {
			common.ErrorLogger.Println("CsvController | GetAllCharacters Error")
			errorResponse := common.NewError(err)
			w.WriteHeader(int(errorResponse.Status))
			json.NewEncoder(w).Encode(errorResponse)
		} else {
			common.InfoLogger.Println("CsvController | GetAllCharacters Success")
			json.NewEncoder(w).Encode(data)
		}
	} else {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	}

}

//GetCharacterById - Receive a response and a requeset, it's the entry point for retrieve a character by id
func GetCharacterById(w http.ResponseWriter, r *http.Request) {
	common.InfoLogger.Println("CsvController | GetCharacterById request")
	w.Header().Set("Content-Type", "application/json")
	authService := services.NewAuthService()
	token, err := authService.ExtractTokenFromRequest(r)
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	} else if isValid, err := authService.ValidateToken(token); isValid {
		service := services.NewCsvService()
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			common.ErrorLogger.Println("CsvController | GetCharacterById Error")
			errorResponse := common.NewError(err)
			w.WriteHeader(int(errorResponse.Status))
			json.NewEncoder(w).Encode(errorResponse)
		}
		data, err := service.GetCharacterById(id)
		if err != nil {
			common.ErrorLogger.Println("CsvController | GetCharacterById Error")
			errorResponse := common.NewError(err)
			w.WriteHeader(errorResponse.Status)
			json.NewEncoder(w).Encode(errorResponse)
		} else {
			common.InfoLogger.Println("CsvController | GetCharacterById Success")
			json.NewEncoder(w).Encode(data)
		}
	} else {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	}
}

func WriteCharactersOnCsv(w http.ResponseWriter, r *http.Request) {
	common.InfoLogger.Println("CsvController | WriteCharactersOnCsv request")
	w.Header().Set("Content-Type", "application/json")
	authService := services.NewAuthService()
	token, err := authService.ExtractTokenFromRequest(r)
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	} else if isValid, err := authService.ValidateToken(token); isValid {
		service := services.NewCsvService()
		err := service.WriteCharactersOnCSV()
		if err != nil {
			common.ErrorLogger.Println("CsvController | WriteCharactersOnCsv Error")
			errorResponse := common.NewError(err)
			w.WriteHeader(errorResponse.Status)
			json.NewEncoder(w).Encode(errorResponse)
		} else {
			common.InfoLogger.Println("CsvController | WriteCharactersOnCsv Success")
			response := common.NewResponse(http.StatusCreated, "Success")
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response)
		}
	} else {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	}

}

//GetAllCharactersConcurrently - Receive a response and a requeset, it's the entry point for retrieve the full list of characters
func GetAllCharactersConcurrently(w http.ResponseWriter, r *http.Request) {
	common.InfoLogger.Println("CsvController | GetAllCharactersConcurrently request")
	w.Header().Set("Content-Type", "application/json")
	authService := services.NewAuthService()
	token, err := authService.ExtractTokenFromRequest(r)
	if err != nil {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	} else if isValid, err := authService.ValidateToken(token); isValid {
		service := services.CsvService{}
		filters, err := getQueryParams(r)
		if err != nil {
			common.ErrorLogger.Println("CsvController | GetAllCharactersConcurrently Error")
			errorResponse := common.NewError(err)
			w.WriteHeader(int(errorResponse.Status))
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		data, err := service.GetAllConcurrently(filters)
		if err != nil {
			common.ErrorLogger.Println("CsvController | GetAllCharactersConcurrently Error")
			errorResponse := common.NewError(err)
			w.WriteHeader(int(errorResponse.Status))
			json.NewEncoder(w).Encode(errorResponse)
			return
		} else {
			common.InfoLogger.Println("CsvController | GetAllCharactersConcurrently Success")
			json.NewEncoder(w).Encode(data)
			return
		}
	} else {
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	}
}

func getQueryParams(r *http.Request) (*common.Filter, error) {
	v := r.URL.Query()
	var itemsPerWorkerInt, itemsInt int64
	var err error
	typeFilter := v.Get("type")
	if typeFilter != "" && typeFilter != "odd" && typeFilter != "even" {
		return nil, errors.New("400")
	}
	items := v.Get("items")
	if items != "" {
		itemsInt, err = strconv.ParseInt(items, 10, 32)
		if err != nil {
			return nil, errors.New("400")
		}
	} else {
		itemsInt = -1
	}
	itemsPerWorker := v.Get("items_per_worker")
	if itemsPerWorker != "" {
		itemsPerWorkerInt, err = strconv.ParseInt(itemsPerWorker, 10, 32)
		if err != nil {
			return nil, errors.New("400")
		}
	} else {
		itemsPerWorkerInt = 1
	}
	return common.NewFilter(typeFilter, itemsInt, itemsPerWorkerInt), nil

}
