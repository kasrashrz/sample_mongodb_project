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
	Insert()
	AddEvent()
	AddUserEvent()
	DeleteById(id string)
	Update()
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

func (events Event) CheckEvent(ctx *gin.Context, id string) (Event, error){
	var event Event
	err := ctx.BindJSON(&event)
	fmt.Println(event)
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
		event.ID = docID
	}
	//TODO : CONTROLLER LAYER (VALIDATORS)
	return event, nil
}

func (events Event) CheckUserEvent (ctx *gin.Context, id string) (UserEvent, error) {
	var UE UserEvent
	err := ctx.BindJSON(&UE)
	fmt.Println(UE)
	if err != nil {
		log.Printf("Unable to parse body %f", err)
		return UserEvent{}, err
	}

	if len(id) > 0 {
		docID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Printf("Failed to create object id from hex %v", id)
			return UserEvent{}, err
		}
		UE.ID = docID
	}
	//TODO : CONTROLLER LAYER (VALIDATORS)
	return UE, nil
}

func (event Event) AddEvent(colName string) ( *mongo.InsertOneResult, error) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	ins := Event{
		ID: 		  primitive.NewObjectID(),
		Name:         event.Name,
		Market_name:  event.Market_name,
		Env:          event.Env,
		EventEndTime: event.EventEndTime,
		Repetition:   Repetition{
			StartPreActiveTime: event.Repetition.StartPreActiveTime,
			StartTime:          event.Repetition.StartTime,
			EndTime:            event.Repetition.EndTime,
			Terminate:          event.Repetition.Terminate,
			StartJoinTime:      event.Repetition.StartJoinTime,
			EndJoinTime:        event.Repetition.EndJoinTime,
		},
	}
	//res, err := collection.InsertOne(ctx, bson.M{"Name": event.Name, "Market_name": event.Market_name, "Env" : event.Env})
	res, err := collection.InsertOne(ctx, ins)
	if err != nil {
		log.Printf("Failed to insert user into DB %f", err)
		return nil, err
	}
	log.Printf("Successfully inserted user %f", res.InsertedID)
	return res, nil
}

func (UE UserEvent) AddUserEvent (colName string) ( *mongo.InsertOneResult, error) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	ins := UserEvent{
		ID:              primitive.NewObjectID(),
		UUID:            UE.UUID,
		GlobalUniqueId:  UE.GlobalUniqueId,
		GamePackageName: UE.GamePackageName,
		Env:             UE.Env,
		MarketName:      UE.MarketName,
		UserEventData:   UserEventData{
			EventId:        UE.UserEventData.EventId,
			UserEventStage: UE.UserEventData.UserEventStage,
			Score:          UE.UserEventData.Score,
			JoinTime:       UE.UserEventData.JoinTime,
			EndTime:        UE.UserEventData.EndTime,
			StartTime:      UE.UserEventData.StartTime,
			PreActiveTime:  UE.UserEventData.PreActiveTime,
		},
	}
	//res, err := collection.InsertOne(ctx, bson.M{"Name": event.Name, "Market_name": event.Market_name, "Env" : event.Env})
	res, err := collection.InsertOne(ctx, ins)
	if err != nil {
		log.Printf("Failed to insert user into DB %f", err)
		return nil, err
	}
	log.Printf("Successfully inserted user %f", res.InsertedID)
	return res, nil
}

func (event Event) DeleteById(colName string, id string) (*Event, error) {
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
	err = collection.FindOneAndDelete(ctx, bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		log.Printf("Unable find user by id %f", err)
		return nil, err
	}
	return &result, nil
}

func (events Event) Update(event Event,colName string)(*mongo.UpdateResult, error){
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection :=  db.GetCollection(colName)
	id, err := primitive.ObjectIDFromHex("6065c2203a84db42c53c8d04")
	filter := bson.M{"_id": id}
	if err != nil {
		log.Printf("Faild to update id %v %v", event.ID, err)
		return nil, err
	}
	update :=  bson.D{
		{"$set", bson.D{
			{"name", string(event.Name)},
			{"market_name", event.Market_name},
			{"env", event.Env},
			{"eventEndTime", event.EventEndTime},
			{"repetition.startPreActiveTime", event.Repetition.StartPreActiveTime},
			{"repetition.startTime", event.Repetition.StartTime},
			{"repetition.endTime", event.Repetition.EndTime},
			{"repetition.Terminate", event.Repetition.Terminate},
			{"repetition.startJoinTime", event.Repetition.StartJoinTime},
			{"repetition.endJoinTime", event.Repetition.EndJoinTime},
		}},
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Faild to update id %v %v", event.ID, err)
		return nil, err
	}
	return res, err
}