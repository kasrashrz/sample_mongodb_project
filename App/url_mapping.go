package app
import
(
    "kasra_medrick.com/mongo/Controllers"
    _ "kasra_medrick.com/mongo/Controllers"
    "kasra_medrick.com/mongo/Models"
)
var event Models.Event
func mapURLs (){
    //TODO: UPDATE EVENT
    //TODO: USER_EVENT MODEL AND ROUTES
    //TODO: SWAGGER!
    router.GET("/ping", Controllers.Ping)
    router.GET("/event/:id", Controllers.GetOneEventController)
    router.GET("/events/all", Controllers.GetAllEventsController)
    router.POST("/event/add", Controllers.CreateEventController)
    router.POST("/event/delete/:id", Controllers.DeleteOneEventController)
    router.POST("/event/update/:id", Controllers.UpdateById)
    /* USER EVENT */
    router.POST("/user_event/add", Controllers.CreateUEController)
    router.GET("/user_events/all", Controllers.GetAllUESController)
    router.POST("/user_event/delete/:id", Controllers.DeleteOneUEController)
	router.GET("/user_event/:id", Controllers.FindOneUEController)


}
