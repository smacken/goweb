package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"web/config"
	"web/controllers"
)

func main() {
	router := mux.NewRouter()

	homeController := new(controllers.HomeController)
	applicationController := new(controllers.ApplicationController)

	router.HandleFunc("/", homeController.Index)
	api := router.PathPrefix("/api/v1").Subrouter()
	api = api.StrictSlash(true)
	api.HandleFunc("application", applicationController.Get).Methods("GET")
	api.HandleFunc("application/{id}", applicationController.Get).Methods("GET")
	api.HandleFunc("application", applicationController.Post).Methods("POST")
	api.HandleFunc("application/{id}", applicationController.Put).Methods("PUT")
	api.HandleFunc("application/{id}", applicationController.Delete).Methods("DELETE")

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	router.NotFoundHandler = http.HandlerFunc(homeController.Index)

	http.Handle("/", router)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
