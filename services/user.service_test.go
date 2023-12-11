package services_test

import (
	m "cv_api/models"
	"cv_api/services"
	"testing"
)

func TestCreate(t *testing.T) {

	userDB := services.NewUserDB()

	//create User
	newUser := m.User{
		Name:  "Test Create",
		Email: "test@test.com",
	}
	_, _, err := userDB.Create(&newUser)

	if err != nil {
		t.Error("La prueba de persistencia de datos del usuario ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito!")
	}
}

func TestGetAll(t *testing.T) {
	userDB := services.NewUserDB()

	_, err := userDB.GetAll()

	if err != nil {
		t.Error("La prueba de recuperación de todos los usuarios ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito!")
	}
}

func TestGetById(t *testing.T) {
	userDB := services.NewUserDB()

	_, err := userDB.GetById(1)

	if err != nil {
		t.Error("La prueba de recuperación usuario por ID ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito!")
	}
}

func TestGetFirst(t *testing.T) {
	userDB := services.NewUserDB()

	_, err := userDB.GetFirst()

	if err != nil {
		t.Error("La prueba de recuperación del primer usuario ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito!")
	}
}
