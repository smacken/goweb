package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"web/controllers"
)

func main() {
	fmt.Println("Go")
	router := mux.NewRouter()
	homeController := new(controllers.HomeController)
	router.HandleFunc("/{view}", homeController.Index).Methods("GET")
	//route("/", homeController.Index)

	router.Handle("/", router)
	http.ListenAndServe(":8089", nil)
}

//func route(pattern string, handler func(http.ResponseWriter, *http.Request)) *mux.Route {
//	handler = logRequest(handler)
//	return router.HandleFunc(pattern, handler)
//}

func logRequest(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = time.Now()
		handler(w, r)
		log.Printf("%s %s %6.3fms", r.Method, r.RequestURI, (time.Since(s).Seconds() * 1000))
	}
}

func writeJson(w http.ResponseWriter, v interface{}) {
	// avoid json vulnerabilities, always wrap v in an object literal
	doc := map[string]interface{}{"d": v}

	if data, err := json.Marshal(doc); err != nil {
		log.Printf("Error marshalling json: %v", err)
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func readJson(r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	var (
		body []byte
		err  error
	)

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("ReadJson couldn't read request body %v", err)
		return false
	}

	if err = json.Unmarshal(body, v); err != nil {
		log.Printf("ReadJson couldn't parse request body %v", err)
		return false
	}

	return true
}
