package services

import (
	"context"
	"errors"
	"log"
	"time"

	configs "com.lopster-pos/config"
	"com.lopster-pos/db"
	"com.lopster-pos/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProfileService(userID string) (models.UsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.UserCollection)

	log.Println("GetProfileService, userID from token:", userID)

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.UsersResponse{}, err
	}

	var u models.Users
	err = col.FindOne(ctx, bson.M{"_id": objID}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		log.Print(err)
		return models.UsersResponse{}, errors.New("user not found")
	}
	if err != nil {
		return models.UsersResponse{}, err
	}

	return models.UsersResponse{
		ID:        u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Role:      u.Role,
		Status:    u.Status,
		Blocked:   u.Blocked,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
	}, nil
}
