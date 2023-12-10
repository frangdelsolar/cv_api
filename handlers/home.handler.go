package handlers

import (
	m "cv_api/models"
	"cv_api/services"
	"fmt"

	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CV API is ONLINE"))
	w.WriteHeader(http.StatusOK)

	//create User
	us := services.NewUserDB()
	user, err := us.GetById(1)
	fmt.Println(user, err)

	newUser := m.User{
		Name:  "Frank",
		Email: "franhhk@micorreo.com",
	}
	createdUser, _, _ := us.Create(&newUser)

	// create DD
	newDD := m.DisplayData{
		Model: m.Model{
			CreatedByID: int(createdUser.ID),
		},
		FirstName:  "Frank",
		LastName:   "Gonzalez",
		Email:      "d@f.com",
		Phone:      "5491122223333",
		City:       "Mendoza",
		Country:    "Argentina",
		ProfilePic: "2",
	}

	dds := services.NewDisplayDataDB()

	one, two, three := dds.Create(&newDD)

	fmt.Println(one, two, three)

}
