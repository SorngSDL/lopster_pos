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

func MenuTypeService(req models.MenuTypeRequest) (models.MenuTypeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.MenuCollection)

	// check duplicate
	var existing models.MenuType
	err := col.FindOne(ctx, bson.M{"menuTypeName": req.MenuTypeName}).Decode(&existing)
	if err == nil {
		return models.MenuTypeResponse{}, errors.New("menu type already exists")
	}

	newID := primitive.NewObjectID()
	mt := models.MenuType{
		ID:           newID,
		MenuTypeName: req.MenuTypeName,
	}

	_, err = col.InsertOne(ctx, mt)
	if err != nil {
		return models.MenuTypeResponse{}, err
	}

	return models.MenuTypeResponse{
		ID:           newID.Hex(),
		MenuTypeName: mt.MenuTypeName,
	}, nil
}
