package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web/models"
)

type ApplicationController struct{}

func (controller *ApplicationController) Get(response http.ResponseWriter, request *http.Request) {
	applications := make([]models.Application, 1)
	app := models.Application{ApplicationID: 1, Name: "First", Description: ""}
	applications = append(applications, app)

	appJson, err := json.Marshal(applications)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(appJson)
}

func (controller *ApplicationController) Post(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var application models.Application
	if err := decoder.Decode(&application); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	log.Println(application.Name)
}

func (controller *ApplicationController) Put(response http.ResponseWriter, request *http.Request) {
	app := JsonRequest(request)
	log.Println(app)
}

func (controller *ApplicationController) Delete(response http.ResponseWriter, request *http.Request) {
}

// JSONResponse attempts to set the status code, c, and marshal the given interface, d, into a response that
// is written to the given ResponseWriter.
func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

func JsonRequest(request *http.Request) models.Application {
	decoder := json.NewDecoder(request.Body)
	var application models.Application
	if err := decoder.Decode(&application); err != nil {
		return application
	}
	return application
}
