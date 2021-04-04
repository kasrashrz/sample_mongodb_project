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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	res, err := user.AddUserEvent("UserEvent")
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success", "id": res.InsertedID})
}

func (UE UserEvent) GetAllUserEvents (ctx *gin.Context){
	UserEvent, err := userModel.FindAllUES("UserEvent")
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event found!", "UserEvent": UserEvent})
	return
}