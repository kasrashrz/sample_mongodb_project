package Models

import (
	"context"
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kasra_medrick.com/mongo/Configs/db"
	"log"
	"time"
)

type MethodHandler interface {
	GetByID(id string)
	FindAll()
	Create()
	AddUser()
}

func (event Event) GetByID(colName string, id string) (*Event, error) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
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

func (event *Event) FindAll (colName string) ([]Event, error) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	var result []Event
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Unable find user by id %f", err)
		return nil, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (event Event) Create (ctx *gin.Context, id string) (Event, error){
	var ev Event
	err := ctx.BindJSON(&ev)
	fmt.Println(ev)
	if err != nil {
		log.Printf("Unable to parse body %f", err)
		return Event{}, err
	}

	if len(id) > 0 {
		docID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Printf("Failed to create object id from hex %v", id)
			return Event{}, err
		}
		ev.ID = docID
	}
	//TODO : CONTROLLER LAYER (VALIDATORS)
	//if len(us.Name) == 0 {
	//	return Event{}, error.New("name is required")
	//}
	//if len(us.Birthday) == 0 {
	//	return Event{}, errors.New("birthday is required")
	//}
	return ev, nil
}

func (event Event) AddEvent(colName string) ( *mongo.InsertOneResult, error) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	//var repetition Repetition
	//var rep = Repetition{
	//	StartPreActiveTime: event.Repetition.StartPreActiveTime,
	//	StartTime:          event.Repetition.StartTime,
	//	EndTime:            event.Repetition.EndTime,
	//	Terminate:          event.Repetition.Terminate,
	//	StartJoinTime:      event.Repetition.StartJoinTime,
	//	EndJoinTime:        repetition.EndJoinTime,
	//}
	//fmt.Println(rep)
	//fmt.Println(event.Repetition)
	res, err := collection.InsertOne(ctx, bson.M{"Name": event.Name, "Market_name": event.Market_name, "Env" : event.Env, "EventEndTime": event.EventEndTime})
	if err != nil {
		log.Printf("Failed to insert user into DB %f", err)
		return nil, err
	}
	log.Printf("Successfully inserted user %f", res.InsertedID)
	return res, nil
}