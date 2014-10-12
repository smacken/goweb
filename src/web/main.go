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
	router.Handle("/views", http.FileServer(http.Dir("../web/views")))
	/*router.HandleFunc("/views/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})*/
	//http.HandleFunc("/", homeController.Index)

	http.Handle("/", router)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
