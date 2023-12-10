package services

import (
	"cv_api/database"
	m "cv_api/models"

	"gorm.io/gorm"
)

type DisplayDataDB struct {
	db *gorm.DB
}

func NewDisplayDataDB() DisplayDataDB {
	db := database.ConnectDB()
	db.AutoMigrate(&m.DisplayData{})

	return DisplayDataDB{db: db}
}

func (ddb *DisplayDataDB) Create(newDD *m.DisplayData) (*m.DisplayData, bool, error) {
	if err := ddb.db.Create(&newDD).Error; err != nil {
		return nil, false, err
	}
	return newDD, true, nil
}

func (ddb *DisplayDataDB) GetById(id int) (*m.DisplayData, error) {
	var dd *m.DisplayData = new(m.DisplayData)
	if err := ddb.db.Preload("CreatedBy").Preload("UpdatedBy").First(&dd, id).Error; err != nil {
		return nil, err
	}

	return dd, nil
}

func (ddb *DisplayDataDB) GetByUserId(userId int) (*m.DisplayData, error) {
	var dd *m.DisplayData = new(m.DisplayData)
	filter := map[string]interface{}{
		"created_by_id": userId,
	}
	if err := ddb.db.Preload("CreatedBy").Preload("UpdatedBy").First(&dd, filter).Error; err != nil {
		return nil, err
	}

	return dd, nil
}

func (ddb *DisplayDataDB) Update(newDD *m.DisplayData) (bool, error) {
	if err := ddb.db.Model(&newDD).Updates(&newDD).Error; err != nil {
		return false, err
	}

	return true, nil
}
