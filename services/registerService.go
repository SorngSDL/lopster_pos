package services

import (
	"context"
	"errors"
	"time"

	configs "com.lopster-pos/config"
	"com.lopster-pos/db"
	"com.lopster-pos/models"
	"com.lopster-pos/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterService(req models.RegisterRequest) (models.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.UserCollection)

	var existing models.Users
	err := col.FindOne(ctx, bson.M{"phone": req.Phone}).Decode(&existing)
	if err == nil {
		return models.RegisterResponse{}, errors.New("phone number already exists")
	}
	if err != mongo.ErrNoDocuments {
		return models.RegisterResponse{}, err
	}

	hash, err := utils.HashPassword(req.PasswordHash)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	now := time.Now()
	newID := primitive.NewObjectID()

	user := models.Users{
		ID:           newID,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Gender:       req.Gender,
		Phone:        req.Phone,
		Role:         req.Role,
		PasswordHash: hash,
		Status:       "ACTIVE",
		Blocked:      false,
		CreatedAt:    now,
		CreatedBy:    req.CreatedBy,
	}

	if _, err := col.InsertOne(ctx, user); err != nil {
		return models.RegisterResponse{}, err
	}

	access, refresh, err := utils.GenerateTokens(newID.Hex(), user.Phone)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return models.RegisterResponse{
		UserID:       newID.Hex(),
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
