package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DisplayData struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedBy  primitive.ObjectID `bson:"created_by" json:"created_by"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
	FirstName  string             `json:"first_name"`
	LastName   string             `json:"last_name"`
	Email      string             `json:"email"`
	Phone      string             `json:"phone"`
	City       string             `json:"city"`
	Country    string             `json:"country"`
	ProfilePic primitive.ObjectID `bson:"profile_pic" json:"profile_pic"`
}

type DisplayDatas []DisplayData
