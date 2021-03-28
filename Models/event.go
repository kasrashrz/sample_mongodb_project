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
    EventEndTime string `json:"eventEndTime"`
	Repetition Repetition 		  `json:"repetition"`
}

type Event_Handler interface {
	Add(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

func (event Event) GetOne(ctx *gin.Context) {
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

func (events Event) GetAll(ctx *gin.Context){
		event, err := userModel.FindAll("Events")
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			ctx.Abort()
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Event found!", "event": event})
		return
}

func (event Event) Add (ctx *gin.Context){
	user, err := userModel.Create(ctx, "")

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