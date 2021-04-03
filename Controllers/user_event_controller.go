package Controllers

import (
	"github.com/gin-gonic/gin"
	"kasra_medrick.com/mongo/Models"
)

var user_event Models.UserEvent

func CreateUEController(ctx *gin.Context){
	user_event.AddOneUserEvent(ctx)
}