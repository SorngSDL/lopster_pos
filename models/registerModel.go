package models

import (
	"time"
)

type RegisterRequest struct {
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

type RegisterResponse struct {
	UserID       string
	AccessToken  string
	RefreshToken string
}
