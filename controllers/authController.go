package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisMG96/academy-go-q42021/common"
	"github.com/LuisMG96/academy-go-q42021/services"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	common.InfoLogger.Println("AuthController | GetToken request")
	w.Header().Set("Content-Type", "application/json")
	service := services.AuthService{}
	//_, data, err := service.ReadFromCSV()
	data, err := service.GetToken("rootuserofapi")
	if err != nil {
		common.ErrorLogger.Println("AuthController | GetToken Error")
		errorResponse := common.NewError(err)
		w.WriteHeader(int(errorResponse.Status))
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		common.InfoLogger.Println("AuthController | GetToken Success")
		response := common.NewResponse(http.StatusCreated, "Success")
		response.Token = data
		w.WriteHeader(response.Status)
		json.NewEncoder(w).Encode(response)

	}
}
