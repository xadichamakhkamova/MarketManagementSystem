package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MongoProduct struct {
	ImageUrl     string             `bson:"image_url" json:"image_url"`
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `json:"name" bson:"name"`
	UniqueNumber string             `json:"unique_number" bson:"unique_number"`
	BagID        string             `json:"bag_id" bson:"bag_id"`
	Price        int64              `json:"price" bson:"price"`
	Size         string             `json:"size" bson:"size"`
	Colors       map[string]int32   `json:"colors" bson:"colors"`
	Count        int32              `json:"count" bson:"count"`
	CreatedAt    string             `json:"created_at" bson:"created_at"`
	UpdatedAt    string             `json:"updated_at" bson:"updated_at"`
	DeletedAt    int32              `json:"deleted_at" bson:"deleted_at"`
}

type DeletedProduct struct {
	Name    string `json:"name" bson:"name"`
	Message string `json:"message"`
}
