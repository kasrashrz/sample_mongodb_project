package Models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserEvent struct {
	ID                        primitive.ObjectID          `bson:"_id" json:"id"`
	UUID                      string                      `bson:"UuId" json:"uuid"`
	GlobalUniqueId            string                      `bson:"GlobalUniqueId" json:"globalUniqueId"`
	GamePackageName           string                      `bson:"GamePackageName" json:"gamePackageName"`
	Env                       string                      `bson:"Env" json:"env"`
	JoinedEventRepetitionUuId []JoinedEventRepetitionUuId `bson:"JoinedEventRepetitionUuId" json:"JoinedEventRepetitionUuId"`
	UserEventData             []UserEventData             `bson:"UserEventData" json:"userEventData"`
}

type User_Event_Handler interface {
	AddOneUserEvent(ctx *gin.Context)
	GetOneUserEvent(ctx *gin.Context)
	GetAllUserEvents(ctx *gin.Context)
	DeleteOneUserEvents(ctx *gin.Context)
	UpdateOneUserEvents(ctx *gin.Context)
}

func (UserEvent UserEvent) AddOneUserEvent(ctx *gin.Context) {
	user, err := UserEvent.CheckUserEvent(ctx, "")

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

func (UserEvents UserEvent) GetAllUserEvents(ctx *gin.Context) {
	UserEvent, err := userModel.FindAllUES("UserEvent")
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User Event found!", "UserEvent": UserEvent})
	return
}

func (UserEvents UserEvent) GetOneUserEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		UserEvent, err := userModel.GetUEByID("UserEvent", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Event found!", "event": UserEvent})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	ctx.Abort()
	return
}

func (UserEvents UserEvent) DeleteOneUserEvents(ctx *gin.Context) {
	id := ctx.Param("id")
	if id != "" {
		UserEvent, err := userModel.DeleteById("UserEvent", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User Event has deleted successfully", "ID": UserEvent.ID})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	ctx.Abort()
	return
}

func (UserEvent UserEvent) UpdateOneUserEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		ctx.Abort()
		return
	}

	SpecUserEvent, err := UserEvent.CheckUserEvent(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	_, err = UserEvent.UpdateUserEvent(SpecUserEvent, "UserEvent", id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User successfully updated ", "user": SpecUserEvent.ID.Hex()})
	return
}

func (UserEvent UserEvent) GetHistoryAPI(ctx *gin.Context) {
	id := ctx.Param("UserEventDataId")
	if id != "" {
		result, err := UserEvent.GetHistory("UserEvent", id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Get History API", "History": result.UserEventData})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Bad Request"})
	ctx.Abort()
	return
}

func (UserEvent UserEvent) ChangeStateAPI(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		ctx.Abort()
		return
	}

	SpecUserEvent, err := UserEvent.CheckUserEvent(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	_, err = UserEvent.ChangeStateUserEvent(SpecUserEvent, "UserEvent", id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Bad Request"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User successfully updated ", "user": SpecUserEvent.ID.Hex()})
	return
}
