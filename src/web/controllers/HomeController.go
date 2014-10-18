package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type HomeController struct{}

func (controller *HomeController) Index(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	view := vars["view"]
	fmt.Printf(view)
	response.Header().Set("Content-Type", "text/html")
	webpage, err := ioutil.ReadFile("../web/public/index.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("index.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}
