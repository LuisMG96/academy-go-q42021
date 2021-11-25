package routes

import (
	"github.com/LuisMG96/academy-go-q42021/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	router mux.Router
}

func New() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/getAllCharacters", controllers.GetAllCharacters).Methods(http.MethodGet)
	r.HandleFunc("/getCharacter/{id}", controllers.GetCharacterById).Methods(http.MethodGet)
	return r
}
