package Controllers

import (
	"github.com/gin-gonic/gin"
	"kasra_medrick.com/mongo/Models"
)
var event Models.Event

func CreateEvent(ctx *gin.Context){
	event.AddOne(ctx)
}

func GetAllEvents(ctx *gin.Context){
	event.GetAll(ctx)
}

func GetOneEvent(ctx *gin.Context){
	event.GetOne(ctx)
}

func DeleteOneEvent(ctx *gin.Context){
	event.DeleteOne(ctx)
}