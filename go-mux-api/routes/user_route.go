package routes

import (
	"github.com/AshishSingh2001/go-mux-api/controllers"
	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router)  {
	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetAUser()).Methods("Get")
}
