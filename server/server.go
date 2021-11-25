package server

import (
	"github.com/LuisMG96/academy-go-q42021/routes"
	"net/http"
)

type Api struct {
	Router http.Handler
}
type Server interface {
	InitRouter() http.Handler
}

func New() Server {
	a := &Api{}
	r := routes.New()
	a.Router = r
	return a
}

func (a *Api) InitRouter() http.Handler {
	return a.Router
}
