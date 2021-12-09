package main

import (
	"log"
	"net/http"

	"github.com/LuisMG96/academy-go-q42021/common"
	"github.com/LuisMG96/academy-go-q42021/server"
)

func main() {
	common.InitLogFile()

	common.InfoLogger.Println("Starting the application")
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.InitRouter()))
}
