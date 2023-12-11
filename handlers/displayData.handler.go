package handlers

import (
	m "cv_api/models"
	"cv_api/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetDisplayData(w http.ResponseWriter, r *http.Request) {

	dds := services.NewDisplayDataDB()

	// TODO: Get the first user from the database until this is handled properly
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

type DisplayDataRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	City        string `json:"city"`
	Country     string `json:"country"`
	ProfilePic  string `json:"profile_pic"`
	CreatedByID int    `json:"created_by_id"`
}

func CreateDisplayData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	decoder := json.NewDecoder(r.Body)
	var requestData DisplayDataRequest
	err := decoder.Decode(&requestData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	// Get the user
	udb := services.NewUserDB()
	user, err := udb.GetById(requestData.CreatedByID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found: %v", err)
		return
	}

	// Create the DisplayData
	newDD := m.DisplayData{
		FirstName:  requestData.FirstName,
		LastName:   requestData.LastName,
		Email:      requestData.Email,
		Phone:      requestData.Phone,
		City:       requestData.City,
		Country:    requestData.Country,
		ProfilePic: requestData.ProfilePic,
		Model: m.Model{
			CreatedByID: int(user.ID),
		},
	}
	ddb := services.NewDisplayDataDB()
	createdDD, _, err := ddb.Create(&newDD)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating DisplayData: %v", err)
		return
	}

	json.NewEncoder(w).Encode(createdDD)
}

func UpdateDisplayData(w http.ResponseWriter, r *http.Request) {

	// Get the id from the query string
	ddid := mux.Vars(r)["id"]
	fmt.Println(ddid)

	// Check if the id is empty or not
	if ddid == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing required parameter: id")
		return
	}

	// Convert the id string to an integer
	id, err := strconv.Atoi(ddid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid id format: %v", err)
		return
	}

	// Read the request body
	decoder := json.NewDecoder(r.Body)
	var requestData DisplayDataRequest
	err = decoder.Decode(&requestData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	// Create the DisplayData
	newDD := m.DisplayData{
		FirstName:  requestData.FirstName,
		LastName:   requestData.LastName,
		Email:      requestData.Email,
		Phone:      requestData.Phone,
		City:       requestData.City,
		Country:    requestData.Country,
		ProfilePic: requestData.ProfilePic,
		Model: m.Model{
			Model: gorm.Model{
				ID: uint(id),
			},
			UpdatedByID: id,
		},
	}
	ddb := services.NewDisplayDataDB()

	// Update the DisplayData
	success, err := ddb.Update(&newDD)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error updating DisplayData: %v", err)
		return
	}

	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func DeleteDisplayData(w http.ResponseWriter, r *http.Request) {
	// Get the id from the query string
	ddid := mux.Vars(r)["id"]
	fmt.Println(ddid)

	// Check if the id is empty or not
	if ddid == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing required parameter: id")
		return
	}

	// Convert the id string to an integer
	id, err := strconv.Atoi(ddid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid id format: %v", err)
		return
	}

	// Delete the DisplayData using the id
	ddb := services.NewDisplayDataDB()
	success, err := ddb.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting DisplayData: %v", err)
		return
	}

	// Respond with the deletion status
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
