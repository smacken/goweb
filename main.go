package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Go")
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")

	router.Handle("/", router)
	http.ListenAndServe(":8089", nil)
}
