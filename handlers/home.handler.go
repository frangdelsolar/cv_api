package handlers

import (

	// userService "cv_api/services/user.service"

	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CV API is ONLINE"))
	w.WriteHeader(http.StatusOK)

	//create User

	// userService.Create(m.User{
	// 	Name:  "Frank",
	// 	Email: "frank@micorreo.com",
	// })

	// create DD
	// uid, _ := primitive.ObjectIDFromHex("6574dfde80cf7f1d96288931")
	// ddService.Create(m.DisplayData{
	// 	CreatedBy:  uid,
	// 	CreatedAt:  time.Now(),
	// 	UpdatedAt:  time.Time{},
	// 	FirstName:  "Frank",
	// 	LastName:   "Gonzalez",
	// 	Email:      "d@f.com",
	// 	Phone:      "5491122223333",
	// 	City:       "Mendoza",
	// 	Country:    "Argentina",
	// 	ProfilePic: uid,
	// })
}
