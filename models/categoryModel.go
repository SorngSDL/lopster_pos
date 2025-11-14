package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CategoryName string             `bson:"categoryName" json:"categoryName"`
}

type CategoryRequest struct {
	CategoryName string `json:"categoryName"`
}

type CategoryResponse struct {
	ID           string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
