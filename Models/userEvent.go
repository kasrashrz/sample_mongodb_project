package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEvent struct {
	ID primitive.ObjectID  `bson:"_id" json:"id"`
	UUID string 		   `bson:"uuid" json:"uuid"`
	GlobalUniqueId string  `bson:"globalUniqueId" json:"globalUniqueId"`
	GamePackageName string `bson:"gamePackageName" json:"gamePackageName"`
	Env string 			   `bson:"env" json:"env"`
	MarketName string 	   `bson:"marketName" json:"marketName"`
	UserEventData UserEventData
}