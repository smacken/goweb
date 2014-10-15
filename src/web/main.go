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
	router.HandleFunc("/", homeController.Index)
	api := router.PathPrefix("/api/").Subrouter()
	router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir("views/"))))
	router.NotFoundHandler = http.HandlerFunc(homeController.Index)
	http.Handle("/", router)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
