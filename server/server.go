package server

import (
	"net/http"

	"github.com/LuisMG96/academy-go-q42021/routes"
)

//Api - Struct of the Api who contains the Router
type Api struct {
	Router http.Handler
}

//Server - Interface that have a InitRouter() function who initialize the server
type Server interface {
	InitRouter() http.Handler
}

//New - Constructor for Server Struct
func New() Server {
	a := &Api{}
	r := routes.New()
	a.Router = r
	return a
}

//InitRouter - implementation of InitRouter function of Server interface in Api Struct
func (a *Api) InitRouter() http.Handler {
	return a.Router
}
