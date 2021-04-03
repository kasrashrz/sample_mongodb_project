package Models

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserEvent struct {
	ID primitive.ObjectID  		`bson:"_id" json:"id"`
	UUID string 		   		`json:"uuid"`
	GlobalUniqueId string  		`json:"globalUniqueId"`
	GamePackageName string 		`json:"gamePackageName"`
	Env string 			   		`json:"env"`
	MarketName string 	   		`json:"marketName"`
	UserEventData UserEventData	`json:"userEventData"`
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
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success", "id": res.InsertedID})
}