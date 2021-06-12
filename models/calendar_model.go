package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Calendar struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	IsDefault bool               `json:"isDefault,omitempty" bson:"isDefault,omitempty"`
	UserId    int                `json:"userId,omitempty" bson:"userId,omitempty"`
	Color     string             `json:"color" bson:"color"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}
