package Models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kasra_medrick.com/mongo/Configs/db"
	"log"
	"time"
)

type Method struct {

}

type MethodHandler interface {
	GetByID(id string)
}

func (event Event) GetByID(id string) (*Event, error) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection("Events")
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Failed to create object id from hex %v", id)
		return nil, err
	}
	var result Event
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		log.Printf("Unable find user by id %f", err)
		return nil, err
	}
	return &result, nil
}