package models

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model  `gorm:"embedded"`
	CreatedByID int  `json:"CreatedByID"`
	CreatedBy   User `json:"CreatedBy"`
	UpdatedByID int  `json:"UpdatedByID,omitempty,omitzero"`
	UpdatedBy   User `json:"UpdatedBy,omitempty,omitzero"`
}
