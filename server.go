package main

import (
	"encoding/json"
	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//"github.com/rs/cors"
	"github.com/thoas/stats"
	//"github.com/unrolled/secure"
	//"goweb/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	// env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := NewRouter()
	app := NewApp()
	app.UseHandler(router)
	app.Run(":" + os.Getenv("port"))
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stats := stats.New().Data()

	b, _ := json.Marshal(stats)

	w.Write(b)
}
