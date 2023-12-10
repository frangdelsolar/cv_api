package services

import (
	"cv_api/database"
	m "cv_api/models"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB() UserDB {
	db := database.ConnectDB()
	db.AutoMigrate(&m.User{})

	return UserDB{db: db}
}

func (u *UserDB) Create(newUser *m.User) (*m.User, bool, error) {
	if err := u.db.Create(&newUser).Error; err != nil {
		return nil, false, err
	}
	return newUser, true, nil
}

func (u *UserDB) GetById(id int) (*m.User, error) {
	var user *m.User = new(m.User)
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserDB) GetAll() (m.Users, error) {
	var users m.Users
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserDB) GetFirst() (*m.User, error) {
	var user *m.User = new(m.User)
	if err := u.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
