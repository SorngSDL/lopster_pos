package services

import (
	"context"
	"errors"
	"time"

	"com.lopster-pos/configs"
	"com.lopster-pos/db"
	"com.lopster-pos/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CategoryService(req models.CategoryRequest) (models.CategoryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.CategoryCollection)

	var existing models.Category
	err := col.FindOne(ctx, bson.M{"categoryName": req.CategoryName}).Decode(&existing)
	if err == nil {
		return models.CategoryResponse{}, errors.New("category already exists")
	}

	newID := primitive.NewObjectID()
	mt := models.Category{
		ID:           newID,
		CategoryName: req.CategoryName,
	}

	_, err = col.InsertOne(ctx, mt)
	if err != nil {
		return models.CategoryResponse{}, err
	}

	return models.CategoryResponse{
		ID:           newID.Hex(),
		CategoryName: mt.CategoryName,
	}, nil
}
