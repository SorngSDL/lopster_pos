package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MenuNameEn     string             `bson:"menuNameEn" json:"menuNameEn"`
	MenuNameLo     string             `bson:"menuNameLo" json:"menuNameLo"`
	CategoryName   string             `bson:"categoryName" json:"categoryName"`
	MenuImage      string             `bson:"menuImage" json:"menuImage"`
	Price          float64            `bson:"price" json:"price"`
	Unit           string             `bson:"unit" json:"unit"`
	PromotionName  string             `bson:"promotionName" json:"promotionName"`
	PromotionPrice float64            `bson:"promotionPrice" json:"promotionPrice"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type MenuRequest struct {
	MenuNameEn     string  `form:"menuNameEn" json:"menuNameEn" binding:"required"`
	MenuNameLo     string  `form:"menuNameLo" json:"menuNameLo" binding:"required"`
	CategoryName   string  `form:"categoryName" json:"categoryName" binding:"required"`
	Price          float64 `form:"price" json:"price" binding:"required"`
	Unit           string  `form:"unit" json:"unit" binding:"required"`
	PromotionName  string  `form:"promotionName" json:"promotionName"`
	PromotionPrice float64 `form:"promotionPrice" json:"promotionPrice"`
}

type MenuResponse struct {
	ID             string    `json:"menuId"`
	MenuNameEn     string    `json:"menuNameEn"`
	MenuNameLo     string    `json:"menuNameLo"`
	CategoryName   string    `json:"categoryName"`
	MenuImage      string    `json:"menuImage"`
	Price          float64   `json:"price"`
	Unit           string    `json:"unit"`
	PromotionName  string    `json:"promotionName"`
	PromotionPrice float64   `json:"promotionPrice"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
