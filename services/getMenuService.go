package services

import (
	"context"
	"time"

	"com.lopster-pos/configs"
	"com.lopster-pos/db"
	"com.lopster-pos/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ListMenuService() ([]models.MenuResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := db.GetCollection(configs.MenuCollection)

	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []models.MenuResponse

	for cur.Next(ctx) {
		var m models.Menu
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}

		results = append(results, models.MenuResponse{
			ID:             m.ID.Hex(),
			MenuNameEn:     m.MenuNameEn,
			MenuNameLo:     m.MenuNameLo,
			CategoryName:   m.CategoryName,
			MenuImage:      m.MenuImage,
			Price:          m.Price,
			Unit:           m.Unit,
			PromotionName:  m.PromotionName,
			PromotionPrice: m.PromotionPrice,
			CreatedAt:      m.CreatedAt,
			UpdatedAt:      m.UpdatedAt,
		})
	}

	return results, nil
}
