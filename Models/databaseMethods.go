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
	"kasra_medrick.com/mongo/Utils"
	"kasra_medrick.com/mongo/Utils/Errors"
	"kasra_medrick.com/mongo/Utils/dates"
	"log"
	"time"
)

type MethodHandler interface {
	GetByID(id string)
	DeleteById(id string)
	FindAllUES()
	FindAll()
	Insert()
	Update()
	GetUEByID()
	AddEvent()
	CheckEvent()
	CheckUserEvent()
	AddUserEvent()
}

func (event Event) GetUEByID(colName string, id string) (*UserEvent, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ServerError := Errors.ServerError("Something Went Wrong")
		log.Printf("Failed to create object id from hex %v", id)
		return nil, ServerError
	}
	var result UserEvent
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		BadReqError := Errors.BadRequest("Invalid ID")
		log.Printf("Unable find event by id %f", err)
		return nil, BadReqError
	}
	return &result, nil
}

func (event Event) GetByID(colName string, id string) (*Event, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ServerError := Errors.ServerError("Something Went Wrong")
		log.Printf("Failed to create object id from hex %v", id)
		return nil, ServerError
	}
	var result Event
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		BadReqError := Errors.BadRequest("Invalid ID")
		log.Printf("Unable find event by id %f", err)
		return nil, BadReqError
	}
	return &result, nil
}

func (event *Event) FindAll(colName string) ([]Event, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	var result []Event
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		BadReqError := Errors.BadRequest("Invalid ID")
		log.Printf("Unable find event by id %f", err)
		return nil, BadReqError
	}
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (event *Event) FindAllUES(colName string) ([]UserEvent, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	var result []UserEvent
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		BadReqError := Errors.BadRequest("invalid ID")
		log.Printf("Unable find event by id %f", err)
		return nil, BadReqError
	}
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func (events Event) CheckEvent(ctx *gin.Context, id string) (Event, *Errors.RestError) {
	var event Event
	err := ctx.BindJSON(&event)
	if err != nil {
		BadReqError := Errors.BadRequest("Unable To Parse Body")
		log.Printf("Unable to parse body %f", err)
		return Event{}, BadReqError
	}

	if len(id) > 0 {
		docID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ServerError := Errors.ServerError("Can't Create object id")
			log.Printf("Failed to create object id from hex %v", id)
			return Event{}, ServerError
		}
		event.ID = docID
	}
	for _, i := range event.Repetition {
		i.RandomRepetitionUuId = Utils.RandomId()
		i.StartTime = dates.EpchoConvertor()

	}
	fmt.Println(event)
	return event, nil
}

func (UES UserEvent) CheckUserEvent(ctx *gin.Context, id string) (UserEvent, *Errors.RestError) {
	var UE UserEvent
	err := ctx.BindJSON(&UE)
	fmt.Println(UE)
	if err != nil {
		BadReqError := Errors.BadRequest("Unable To Parse Body")
		log.Printf("Unable to parse body %f", err)
		return UserEvent{}, BadReqError
	}

	if len(id) > 0 {
		docID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ServerError := Errors.ServerError("Can't Create object id")
			log.Printf("Failed to create object id from hex %v", id)
			return UserEvent{}, ServerError
		}
		UE.ID = docID
	}
	//TODO : CONTROLLER LAYER (VALIDATORS)
	return UE, nil
}

func (event Event) AddEvent(colName string) (*mongo.InsertOneResult, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	ins := Event{
		ID:                             primitive.NewObjectID(),
		Name:                           event.Name,
		Env:                            event.Env,
		EventEndType:                   event.EventEndType,
		ClientType:                     event.ClientType,
		PeriodTimeForMiddleJoinTillEnd: event.PeriodTimeForMiddleJoinTillEnd,
		ConfigVersion:                  event.ConfigVersion,
		States:                         event.States,
		VersionMetaData:                event.VersionMetaData,
	}
	for _, repetitions := range event.Repetition {
		repetitions.RandomRepetitionUuId = Utils.RandomId()
		repetitions.StartTime = dates.EpchoConvertor()
		ins.Repetition = append(ins.Repetition, repetitions)
	}
	res, err := collection.InsertOne(ctx, ins)
	if err != nil {
		ServerError := Errors.ServerError("Failed To Insert")
		log.Printf("Failed to insert event into DB %f", err)
		return nil, ServerError
	}
	log.Printf("Successfully inserted event %f", res.InsertedID)
	return res, nil
}

func (UE UserEvent) AddUserEvent(colName string) (*mongo.InsertOneResult, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	EventCollection := db.GetCollection("Events")
	var result Event
	EventId, _ := primitive.ObjectIDFromHex(UE.UserEventData.EventId)
	filter := bson.M{
		"_id": EventId,
		"Repetition.StartTime": UE.UserEventData.StartTime,
	}
	err := EventCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("*+*+*++*+*+*+*+*+*+*+", result.Repetition)
	ins := UserEvent{
	ID:              primitive.NewObjectID(),
		UUID:            UE.UUID,
		GlobalUniqueId:  UE.GlobalUniqueId,
		GamePackageName: UE.GamePackageName,
		Env:             UE.Env,
		JoinedEventRepetitionUuId: UE.JoinedEventRepetitionUuId,
		UserEventData: UserEventData{
			EventId:        UE.UserEventData.EventId,
			UserEventStage: UE.UserEventData.UserEventStage,
			Score:          UE.UserEventData.Score,
			JoinTime:       UE.UserEventData.JoinTime,
			EndTime:        UE.UserEventData.EndTime,
			StartTime:      UE.UserEventData.StartTime,
			PreActiveTime:  UE.UserEventData.PreActiveTime,
		},
	}
	//for _, repetitionData := range result.Repetition {
	//	repetitionData.RandomRepetitionUuId = Utils.RandomId()
	//	repetitionData.StartTime = dates.EpchoConvertor()
	//	ins.UserEventData = append(ins.UserEventData, repetitionData)
	//}
	res, err := collection.InsertOne(ctx, ins)
	if err != nil {
		ServerError := Errors.ServerError("Failed to insert")
		log.Printf("Failed to insert event into DB %f", err)
		return nil, ServerError
	}
	log.Printf("Successfully inserted event %f", res.InsertedID)
	return res, nil
}

func (event Event) DeleteById(colName string, id string) (*Event, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ServerError := Errors.ServerError("Failed to create")
		log.Printf("Failed to create object id from hex %v", id)
		return nil, ServerError
	}
	var result Event
	err = collection.FindOneAndDelete(ctx, bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		BadReqError := Errors.BadRequest("Invalid ID")
		log.Printf("Unable find Event by id %f", err)
		return nil, BadReqError
	}
	return &result, nil
}

func (events Event) Update(event Event, colName string, EventId string) (*mongo.UpdateResult, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	id, err := primitive.ObjectIDFromHex(EventId)
	filter := bson.M{"_id": id}
	if err != nil {
		ServerError := Errors.ServerError("failed to update")
		log.Printf("Faild to update id %v %v", event.ID, err)
		return nil, ServerError
	}
	ins := Event{
		Name:         event.Name,
		Env:          event.Env,
		EventEndType: event.EventEndType,
		Repetition:   nil,
	}
	for _, repetitions := range event.Repetition {
		repetitions.RandomRepetitionUuId = Utils.RandomId()
		repetitions.StartTime = dates.EpchoConvertor()
		ins.Repetition = append(ins.Repetition, repetitions)
	}
	// TODO: FULL UPDATE
	update := bson.M{
		"$set": bson.M{
			"Name": event.Name,
			"Env":  event.Env,
			"EventEndType": event.EventEndType,
			"ClientType": event.ClientType,
			"PeriodTimeForMiddleJoinTillEnd": event.PeriodTimeForMiddleJoinTillEnd,
			"ConfigVersion": event.ConfigVersion,
			"States": event.States,
			"VersionMetaData": event.VersionMetaData,
			"Repetition":ins.Repetition,
		},
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ServerError := Errors.ServerError("Failed To update")
		log.Printf("Faild to update id %v %v", event.ID, err)
		return nil, ServerError
	}
	return res, nil
}

func (UES UserEvent) UpdateUserEvent(UE UserEvent, colName string, UserEventId string) (*mongo.UpdateResult, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	id, err := primitive.ObjectIDFromHex(UserEventId)
	filter := bson.M{"_id": id}
	if err != nil {
		ServerError := Errors.ServerError("failed to update")
		log.Printf("Faild to update id %v %v", UE.ID, err)
		return nil, ServerError
	}
	update := bson.D{
		{"$set", bson.D{
			{"UUID", UE.UUID},
			{"GlobalUniqueId", UE.GlobalUniqueId},
			{"GamePackageName", UE.GamePackageName},
			{"Env", UE.Env},
			{"JoinedEventRepetitionUuId", UE.JoinedEventRepetitionUuId},
			{"UserEventData.EventId", UE.UserEventData.EventId},
			{"UserEventData.UserEventStage", UE.UserEventData.UserEventStage},
			{"UserEventData.Score", UE.UserEventData.Score},
			{"UserEventData.JoinTime", UE.UserEventData.JoinTime},
			{"UserEventData.EndTime", UE.UserEventData.EndTime},
			{"UserEventData.StartTime", UE.UserEventData.StartTime},
			{"UserEventData.PreActiveTime", UE.UserEventData.PreActiveTime},
		}},
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ServerError := Errors.ServerError("Failed To update")
		log.Printf("Faild to update id %v %v", UE.ID, err)
		return nil, ServerError
	}
	return res, nil
}

func (events Event) TerminateAPI(colName string, id string) (*mongo.UpdateResult, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	Time_Now := dates.EpchoConvertor()
	NewId, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(Time_Now)
	filter := bson.M{
		"_id" : NewId,
		"Repetition.EndTime": bson.M{"$lt":Time_Now},
		"Repetition.Terminate": false,

	}
	update := bson.M{
		"$set": bson.M{
			"Repetition.$.Terminate" : true,
		},
	}
	res, err := collection.UpdateMany(ctx, filter, update)

	if err != nil {
		ServerError := Errors.ServerError("Failed To update")
		log.Printf("Faild to update id %v %v", id, err)
		return nil, ServerError
	}
	return res, nil
}
