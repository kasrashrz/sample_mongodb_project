package Controllers

import (
	"github.com/gin-gonic/gin"
	"kasra_medrick.com/mongo/Models"
)

var event Models.Event

func CreateEventController(ctx *gin.Context) {
	event.AddOneEvent(ctx)
}

func GetAllEventsController(ctx *gin.Context) {
	event.GetAllEvents(ctx)
}

func GetOneEventController(ctx *gin.Context) {
	event.GetOneEvent(ctx)
}

func DeleteOneEventController(ctx *gin.Context) {
	event.DeleteOneEvent(ctx)
}

func UpdateById(ctx *gin.Context) {
	event.UpdateOneEvent(ctx)
}

func TerminateEventController(ctx *gin.Context) {
	event.Terminate(ctx)
}
func GetActiveEventController(ctx *gin.Context) {
	event.GetActive(ctx)
}
