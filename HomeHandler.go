package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	html := vars["html"]
	res.Header().Set("Content-Type", "text/html")
}
