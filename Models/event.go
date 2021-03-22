package Models

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kasra_medrick.com/mongo/Configs/db"
	"log"
	"net/http"
	"time"
)

var client *mongo.Client
type  EventEndTimeType uint
const (
	absolute EventEndTimeType = iota
	relative
)
type Event struct {
	Name string
	Market_name string
	Env string
    EventEndTime EventEndTimeType
	//states map[]
	Repetition Repetition
}

type Event_Handler interface {
	CREATE(ctx *gin.Context)
	RETRIEVE(ctx *gin.Context)

}
var userModel = new(Event)

func (h Event) GetByID(id string) (*Event, error) {
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

func (u Event) RETREIVE(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		user, err := userModel.GetByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User found!", "user": user})
		return
	}
	fmt.Println(c.Param("id"))
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}