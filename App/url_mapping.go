package app

import (
	"kasra_medrick.com/mongo/Controllers"
	_ "kasra_medrick.com/mongo/Controllers"
	"kasra_medrick.com/mongo/Models"
)

var event Models.Event

func mapURLs() {
	//TODO: UPDATE EVENT
	//TODO: SWAGGER!
	router.GET("/ping", Controllers.Ping)
	router.GET("/event/:id", Controllers.GetOneEventController)
	router.GET("/events/all", Controllers.GetAllEventsController)
	router.POST("/event/add", Controllers.CreateEventController)
	router.POST("/event/delete/:id", Controllers.DeleteOneEventController)
	router.POST("/event/update/:id", Controllers.UpdateById)
	router.POST("/event/terminate/:id", Controllers.TerminateEventController)
	/* USER EVENT */
	router.GET("/user_events/all", Controllers.GetAllUESController)
	router.GET("/user_event/:id", Controllers.FindOneUEController)
	router.POST("/user_event/add", Controllers.CreateUEController)
	router.POST("/user_event/delete/:id", Controllers.DeleteOneUEController)
	router.POST("/user_event/update/:id", Controllers.UpdateOneUEController)

}
