package services_test

import (
	m "cv_api/models"
	"cv_api/services"
	"testing"
)

func TestCreateDisplayData(t *testing.T) {

	ddb := services.NewDisplayDataDB()
	udb := services.NewUserDB()

	newUser := m.User{
		Name:  "Test Create",
		Email: "test1@test.com",
	}
	nUser, _, _ := udb.Create(&newUser)

	// Create valid DisplayData
	newDD := &m.DisplayData{
		Model: m.Model{
			CreatedByID: int(nUser.ID),
		},
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@email.com",
	}
	_, _, err := ddb.Create(newDD)

	if err != nil {
		t.Error("There was an error creating valid DisplayData")
		t.Fail()
	} else {
		t.Log("DisplayData created successfully!")
	}

}

func TestUpdateDisplayData(t *testing.T) {

	ddb := services.NewDisplayDataDB()
	udb := services.NewUserDB()

	newUser := m.User{
		Name:  "Test Create",
		Email: "test1@test.com",
	}
	nUser, _, _ := udb.Create(&newUser)

	// Create valid DisplayData
	newDD := &m.DisplayData{
		Model: m.Model{
			CreatedByID: int(nUser.ID),
		},
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@email.com",
	}
	newData, _, _ := ddb.Create(newDD)

	// Update DisplayData
	newData.Country = "Argentina"
	newData.City = "Mendoza"
	ddb.Update(newData)

	// Check if updated
	retrievedDD, _ := ddb.GetById(int(newData.ID))

	if retrievedDD.Country != "Argentina" || retrievedDD.City != "Mendoza" {
		t.Error("There was an error updating DisplayData")
		t.Fail()
	} else {
		t.Log("DisplayData updated successfully!")
	}

}
