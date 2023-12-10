package main

import (
	"cv_api/database"
	"cv_api/handlers"
	"cv_api/initializers"
	"net/http"

	"github.com/rs/cors"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnv()
	DB = database.ConnectDB()
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/api/display-data/", handlers.GetDisplayData)

	http.Handle("/", r)

	handler := cors.AllowAll().Handler(r)

	http.ListenAndServe(":8080", handler)

}
