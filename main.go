package main

import (
	"cv_api/handlers"
	displayData_handler "cv_api/handlers/displayData.handler"
	"cv_api/initializers"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/api/display-data/", displayData_handler.GetDisplayData)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)

}
