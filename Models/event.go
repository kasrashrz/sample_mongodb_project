package Models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var client *mongo.Client
type  EventEndTimeType uint
var userModel = new(Event)

const (
	absolute EventEndTimeType = iota
	relative
)

type Event struct {
	ID 		  primitive.ObjectID  `bson:"_id" json:"id"`
	Name string	     			  `json:"name"`
	Market_name string 			  `json:"market_name"`
	Env string 					  `json:"env"`
    EventEndTime string 		  `json:"eventEndTime"`
	Repetition Repetition 		  `json:"repetition"`
}

type Event_Handler interface {
	AddOneEvent(ctx *gin.Context)
	GetOneEvent(ctx *gin.Context)
	GetAllEvents(ctx *gin.Context)
	DeleteOneEvent(ctx *gin.Context)
	UpdateOneEvent(ctx *gin.Context)
}

func (event Event) GetOneEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		event, err := userModel.GetByID("Events",id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve event", "error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Event found!", "event": event})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	ctx.Abort()
	return
}

func (events Event) GetAllEvents(ctx *gin.Context){
		event, err := userModel.FindAll("Events")
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			ctx.Abort()
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Event found!", "event": event})
		return
}

func (event Event) AddOneEvent(ctx *gin.Context){
	user, err := userModel.Check(ctx, "")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	res, err := user.AddEvent("Events")

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success", "id": res.InsertedID})
}

func (event Event) DeleteOneEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		event, err := userModel.DeleteById("Events",id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error to delete event", "error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Event has deleted successfully", "ID": event.ID})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	ctx.Abort()
	return
}

func (event Event) UpdateOneEvent(ctx *gin.Context) {
	if ctx.Param("id") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		ctx.Abort()
		return
	}

	SpecEvent, err := event.Check(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	_, err = event.Update(SpecEvent,"Events")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User successfully updated ", "user": SpecEvent.ID.Hex()})
	ctx.Abort()
	return
}