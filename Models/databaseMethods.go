package Models

import (
	"context"
	_ "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
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
	var result Event
	EventCollection := db.GetCollection("Events")
	ins := UserEvent{
		ID:                        primitive.NewObjectID(),
		UUID:                      UE.UUID,
		GlobalUniqueId:            UE.GlobalUniqueId,
		GamePackageName:           UE.GamePackageName,
		Env:                       UE.Env,
		JoinedEventRepetitionUuId: UE.JoinedEventRepetitionUuId,
	}
	for _, UserEventData := range UE.UserEventData {
		id, _ := primitive.ObjectIDFromHex(UserEventData.EventId)
		filter := bson.M{
			"_id": id,
			"Repetition": bson.M{
				"$elemMatch": bson.M{
					"Terminate": true,
				},
			},
		}
		_ = EventCollection.FindOne(ctx, filter).Decode(&result)
		for _, obj := range result.Repetition {
			if obj.Terminate != true && obj.EndTime == UserEventData.EndTime {
				fmt.Println("+-+-", obj)
				UserEventData.EndTime = obj.EndTime
				UserEventData.StartTime = obj.StartTime
				UserEventData.PreActiveTime = obj.StartPreActiveTime
				ins.UserEventData = append(ins.UserEventData, UserEventData)
			}
		}
	}
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
			"Name":                           event.Name,
			"Env":                            event.Env,
			"EventEndType":                   event.EventEndType,
			"ClientType":                     event.ClientType,
			"PeriodTimeForMiddleJoinTillEnd": event.PeriodTimeForMiddleJoinTillEnd,
			"ConfigVersion":                  event.ConfigVersion,
			"States":                         event.States,
			"VersionMetaData":                event.VersionMetaData,
			"Repetition":                     ins.Repetition,
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
	ins := UserEvent{
		UUID:                      UE.UUID,
		GlobalUniqueId:            UE.GlobalUniqueId,
		GamePackageName:           UE.GamePackageName,
		Env:                       UE.Env,
		JoinedEventRepetitionUuId: UE.JoinedEventRepetitionUuId,
	}
	for _, UserEventData := range UE.UserEventData {
		ins.UserEventData = append(ins.UserEventData, UserEventData)
	}
	update := bson.D{
		{"$set", bson.D{
			{"UuId", UE.UUID},
			{"GlobalUniqueId", UE.GlobalUniqueId},
			{"GamePackageName", UE.GamePackageName},
			{"Env", UE.Env},
			{"JoinedEventRepetitionUuId", UE.JoinedEventRepetitionUuId},
			{"UserEventData", ins.UserEventData},
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
	UserEventCollection := db.GetCollection("UserEvent")
	Time_Now := dates.EpchoConvertor()
	NewId, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(Time_Now)
	EventFilter := bson.M{
		"_id": NewId,
		"Repetition": bson.M{"$elemMatch": bson.M{
			"EndTime":   bson.M{"$gt": Time_Now},
			"StartTime": bson.M{"$lt": Time_Now},
			"Terminate": false,
		}},
	}
	EventUpdate := bson.M{
		"$set": bson.M{
			"Repetition.$.Terminate": true,
		},
	}
	UserEventFilter := bson.M{
		"UserEventData": bson.M{"$elemMatch": bson.M{
			"EventId":        id,
			"EndTime":        bson.M{"$gt": Time_Now},
			"UserEventStage": bson.M{"$ne": "terminated"},
		}},
	}
	UserEventUpdate := bson.M{
		"$set": bson.M{
			"UserEventData.$.UserEventStage": "terminated",
		},
	}
	res, err := UserEventCollection.UpdateOne(ctx, UserEventFilter, UserEventUpdate)
	if err != nil {
		ServerError := Errors.ServerError("Failed To update")
		log.Printf("Faild to update id %v %v", id, err)
		return nil, ServerError
	}
	_, UserEventErr := collection.UpdateOne(ctx, EventFilter, EventUpdate)
	if UserEventErr != nil {
		ServerError := Errors.ServerError("Failed To update")
		log.Printf("Faild to update id %v %v", id, err)
		return nil, ServerError
	}
	return res, nil
}

func (UserEvents UserEvent) GetHistory(colName string, id string) (*UserEvent, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	NewId, err := primitive.ObjectIDFromHex(id)
	fmt.Println(NewId)
	if err != nil {
		ServerError := Errors.ServerError("Failed to create")
		log.Printf("Failed to create object id from hex %v", id)
		return nil, ServerError
	}
	filter := bson.M{
		"UserEventData.EventId": id,
	}
	findOps := options.Find()
	findOps.SetProjection(bson.M{
		"UserEventData.$": 1,
	})
	var results UserEvent
	err = collection.FindOne(ctx, filter).Decode(&results)
	if err != nil {
		BadReqError := Errors.BadRequest("Invalid ID")
		log.Printf("Unable find event by id %f", err)
		return nil, BadReqError
	}
	return &results, nil
}

func (event *Event) GetActiveAPI(colName string) ([]*Event, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	Time_Now := dates.EpchoConvertor()
	filter := bson.M{
		"Repetition": bson.M{"$elemMatch": bson.M{
			"EndTime":   bson.M{"$gt": Time_Now},
			"StartTime": bson.M{"$lt": Time_Now},
			"Terminate": false,
		}},
	}
	var results []*Event
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		BadReqError := Errors.BadRequest("Invalid ID")
		log.Printf("Unable find event by id", err)
		return results, BadReqError
	}
	for cur.Next(ctx) {
		var t Event
		err := cur.Decode(&t)
		if err != nil {
			BadReqError := Errors.BadRequest("Invalid ID")
			return results, BadReqError
		}
		results = append(results, &t)
	}
	return results, nil
}

func (UES UserEvent) ChangeStateUserEvent(UE UserEvent, colName string, UserEventId string) (*mongo.UpdateResult, *Errors.RestError) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection(colName)
	//Time_Now := dates.EpchoConvertor()
	filter := bson.M{
		"UserEventData": bson.M{
			"$elemMatch": bson.M{
				"EventId": UserEventId,
				//"EndTime": bson.M{"$gt": Time_Now},
				//"StartTime": bson.M{"$lt": Time_Now},
				//"Terminate": false,
			}},
	}
	ins := UserEvent{
		UUID:            UE.UUID,
		GlobalUniqueId:  UE.GlobalUniqueId,
		GamePackageName: UE.GamePackageName,
		Env:             UE.Env,
	}
	for _, UserEventData := range UE.UserEventData {
		ins.UserEventData = append(ins.UserEventData, UserEventData)
	}
	update := bson.D{
		{"$set", bson.D{
			{"UserEventData", ins.UserEventData},
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
