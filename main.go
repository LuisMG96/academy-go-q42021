package main

import (
	"github.com/LuisMG96/academy-go-q42021/server"
	"log"
	"net/http"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.InitRouter()))
}
