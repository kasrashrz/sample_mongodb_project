package Controllers

import (
	"github.com/gin-gonic/gin"
	"kasra_medrick.com/mongo/Models"
)

var user_event Models.UserEvent

func CreateUEController(ctx *gin.Context) {
	user_event.AddOneUserEvent(ctx)
}

func GetAllUESController(ctx *gin.Context) {
	user_event.GetAllUserEvents(ctx)
}

func DeleteOneUEController(ctx *gin.Context) {
	user_event.DeleteOneUserEvents(ctx)
}

func FindOneUEController(ctx *gin.Context) {
	user_event.GetOneUserEvent(ctx)
}

func UpdateOneUEController(ctx *gin.Context) {
	user_event.UpdateOneUserEvent(ctx)
}

func GetHistoryController(ctx *gin.Context){
	user_event.GetHistoryAPI(ctx)
}
func ChangeStateController(ctx *gin.Context){
	user_event.ChangeStateAPI(ctx)
}