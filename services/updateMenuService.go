// services/menu_service.go
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

func UpdateMenuService(id string, req models.MenuRequest, imagePath *string) (models.MenuResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.MenuCollection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.MenuResponse{}, errors.New("invalid menu id")
	}

	// find existing
	var existing models.Menu
	if err := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&existing); err != nil {
		return models.MenuResponse{}, errors.New("menu not found")
	}

	// keep old image if no new one
	menuImage := existing.MenuImage
	if imagePath != nil && *imagePath != "" {
		menuImage = *imagePath
	}

	update := bson.M{
		"$set": bson.M{
			"menuNameEn":     req.MenuNameEn,
			"menuNameLo":     req.MenuNameLo,
			"categoryName":   req.CategoryName,
			"menuImage":      menuImage,
			"price":          req.Price,
			"unit":           req.Unit,
			"promotionName":  req.PromotionName,
			"promotionPrice": req.PromotionPrice,
			"updatedAt":      time.Now(),
		},
	}

	if _, err := col.UpdateByID(ctx, objID, update); err != nil {
		return models.MenuResponse{}, err
	}

	if err := col.FindOne(ctx, bson.M{"_id": objID}).Decode(&existing); err != nil {
		return models.MenuResponse{}, err
	}

	return models.MenuResponse{
		ID:             existing.ID.Hex(),
		MenuNameEn:     existing.MenuNameEn,
		MenuNameLo:     existing.MenuNameLo,
		CategoryName:   existing.CategoryName,
		MenuImage:      existing.MenuImage,
		Price:          existing.Price,
		Unit:           existing.Unit,
		PromotionName:  existing.PromotionName,
		PromotionPrice: existing.PromotionPrice,
		CreatedAt:      existing.CreatedAt,
		UpdatedAt:      existing.UpdatedAt,
	}, nil
}
