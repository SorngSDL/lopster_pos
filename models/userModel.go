package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string
	LastName     string
	Gender       string
	Phone        string
	Role         string
	PasswordHash string
	Status       string
	Blocked      bool
	CreatedAt    time.Time
	CreatedBy    string
}

type UsersResponse struct {
	ID        string `json:"userId"`
	FirstName string
	LastName  string
	Gender    string
	Phone     string
	Role      string
	Status    string
	Blocked   bool
	CreatedAt time.Time
	CreatedBy string
}
