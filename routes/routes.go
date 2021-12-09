package routes

import (
	"net/http"

	"github.com/LuisMG96/academy-go-q42021/controllers"
	"github.com/gorilla/mux"
)

//Router - Struct that contains the mux.Router to access the endpoints
type Router struct {
	router mux.Router
}

//New - Returns an http.Handler with the help of mux
func New() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/getAllCharacters", controllers.GetAllCharacters).Methods(http.MethodGet)
	r.HandleFunc("/getCharacter/{id}", controllers.GetCharacterById).Methods(http.MethodGet)
	r.HandleFunc("/writeCharacters", controllers.WriteCharactersOnCsv).Methods(http.MethodPost)
	r.HandleFunc("/getAllConcurrently", controllers.GetAllCharactersConcurrently).Methods(http.MethodGet)
	r.HandleFunc("/getToken", controllers.GetToken).Methods(http.MethodGet)
	return r
}
