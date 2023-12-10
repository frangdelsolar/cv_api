package handlers

import (
	"cv_api/services"
	"encoding/json"
	"net/http"
)

func GetDisplayData(w http.ResponseWriter, r *http.Request) {

	dds := services.NewDisplayDataDB()

	us := services.NewUserDB()
	user, err := us.GetFirst()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	data, err := dds.GetByUserId(int(user.ID))

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// Encode the dd object to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}
