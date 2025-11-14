package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MenuType struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	MenuTypeName string             `bson:"menuTypeName" json:"menuTypeName"`
}

type MenuTypeRequest struct {
	MenuTypeName string `json:"menuTypeName"`
}

type MenuTypeResponse struct {
	ID           string `json:"menuTypeId"`
	MenuTypeName string `json:"menuTypeName"`
}
