package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User datos del usuario
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
}

// Users lista de usuarios
type Users []*User
