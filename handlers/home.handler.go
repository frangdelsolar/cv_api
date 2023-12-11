package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CV API is ONLINE"))
	w.WriteHeader(http.StatusOK)

}
