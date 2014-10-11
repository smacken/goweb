package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type HomeController struct{}

func (controller *HomeController) Index(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(req)
	html := vars["html"]
	fmt.Printf(html)
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(response, "Hello world")
}
