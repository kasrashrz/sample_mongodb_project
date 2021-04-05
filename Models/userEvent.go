package Models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserEvent struct {
	ID primitive.ObjectID  		`bson:"_id" json:"id"`
	UUID string 		   		`bson:"UUID" json:"uuid"`
	GlobalUniqueId string  		`bson:"GlobalUniqueId" json:"globalUniqueId"`
	GamePackageName string 		`bson:"GamePackageName" json:"gamePackageName"`
	Env string 			   		`bson:"Env" json:"env"`
	MarketName string 	   		`bson:"MarketName" json:"marketName"`
	UserEventData UserEventData	`bson:"UserEventData" json:"userEventData"`
}

type User_Event_Handler interface {
	AddOneUserEvent(ctx *gin.Context)
	GetOneUserEvent(ctx *gin.Context)
	GetAllUserEvents(ctx *gin.Context)
	DeleteOneUserEvents(ctx *gin.Context)
	UpdateOneUserEvents(ctx *gin.Context)
}

func (UE UserEvent) AddOneUserEvent (ctx *gin.Context){
	user, err := userModel.CheckUserEvent(ctx, "")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		ctx.Abort()
		return
	}

	res, err := user.AddUserEvent("UserEvent")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success", "id": res.InsertedID})
	return
}

func (UE UserEvent) GetAllUserEvents (ctx *gin.Context){
	UserEvent, err := userModel.FindAllUES("UserEvent")
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User Event found!", "UserEvent": UserEvent})
	return
}

func (UE UserEvent) GetOneUserEvent (ctx *gin.Context){
	id := ctx.Param("id")
	if id != "" {
		event, err := userModel.GetUEByID("UserEvent",id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
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

func (UE UserEvent) DeleteOneUserEvents (ctx *gin.Context){
	id := ctx.Param("id")
	if id != "" {
		event, err := userModel.DeleteById("UserEvent",id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User Event has deleted successfully", "ID": event.ID})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	ctx.Abort()
	return
}
