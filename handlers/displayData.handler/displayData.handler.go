package displayData_handler

import (
	ddService "cv_api/services/displayData.service"
	"encoding/json"
	"net/http"
)

const userId = "6574dfde80cf7f1d96288931"

func GetDisplayData(w http.ResponseWriter, r *http.Request) {

	dd, err := ddService.Read(userId)

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	// Encode the dd object to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(dd)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}
