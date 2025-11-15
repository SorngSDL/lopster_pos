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

func CreateMenuService(req models.MenuRequest, imagePath string) (models.MenuResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.MenuCollection)

	// ตัวอย่างเช็คซ้ำตามชื่อ (แล้วแต่ business rule)
	var existing models.Menu
	err := col.FindOne(ctx, bson.M{
		"menuNameEn":   req.MenuNameEn,
		"categoryName": req.CategoryName,
	}).Decode(&existing)
	if err == nil {
		return models.MenuResponse{}, errors.New("menu already exists")
	}

	now := time.Now()
	newID := primitive.NewObjectID()

	mt := models.Menu{
		ID:             newID,
		MenuNameEn:     req.MenuNameEn,
		MenuNameLo:     req.MenuNameLo,
		CategoryName:   req.CategoryName,
		MenuImage:      imagePath, // เก็บ path หรือ URL
		Price:          req.Price,
		Unit:           req.Unit,
		PromotionName:  req.PromotionName,
		PromotionPrice: req.PromotionPrice,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	_, err = col.InsertOne(ctx, mt)
	if err != nil {
		return models.MenuResponse{}, err
	}

	return models.MenuResponse{
		ID:             newID.Hex(),
		MenuNameEn:     mt.MenuNameEn,
		MenuNameLo:     mt.MenuNameLo,
		CategoryName:   mt.CategoryName,
		MenuImage:      mt.MenuImage,
		Price:          mt.Price,
		Unit:           mt.Unit,
		PromotionName:  mt.PromotionName,
		PromotionPrice: mt.PromotionPrice,
		CreatedAt:      mt.CreatedAt,
		UpdatedAt:      mt.UpdatedAt,
	}, nil
}
