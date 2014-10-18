package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"web/controllers"
)

func main() {
	router := mux.NewRouter()
	homeController := new(controllers.HomeController)
	applicationController := new(controllers.ApplicationController)
	router.HandleFunc("/", homeController.Index)
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("application", applicationController.Get)
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	router.NotFoundHandler = http.HandlerFunc(homeController.Index)
	http.Handle("/", router)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
