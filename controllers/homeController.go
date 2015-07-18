package controllers

import (
	"fmt"
	"net/http"
)

// GET "/"
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello %v\n")
}

// GET "/"
func DocsHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Api reference %v\n")
}
