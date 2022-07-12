package main

import (
	"log"
	"net/http"

	"github.com/AshishSingh2001/go-mux-api/configs"
	"github.com/AshishSingh2001/go-mux-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// run database
	configs.ConnectDB()

	routes.UserRoute(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
