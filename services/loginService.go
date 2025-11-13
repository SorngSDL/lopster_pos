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
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginService(req models.LoginRequest) (models.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.UserCollection)

	var user models.Users
	err := col.FindOne(ctx, bson.M{"phone": req.Phone}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return models.LoginResponse{}, errors.New("invalid phone or password")
	}
	if err != nil {
		return models.LoginResponse{}, err
	}

	if !utils.CheckPassword(user.PasswordHash, req.Password) {
		return models.LoginResponse{}, errors.New("invalid phone or password")
	}

	access, refresh, err := utils.GenerateTokens(user.ID.Hex(), user.Phone)
	if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		ID:           user.ID.Hex(),
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
