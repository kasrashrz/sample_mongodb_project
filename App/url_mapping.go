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
    router.GET("/event/:id", Controllers.GetOneEvent)
    router.GET("/events/all", Controllers.GetAllEvents)
	router.POST("/event/add", Controllers.CreateEvent)
    router.POST("/event/delete/:id", Controllers.DeleteOneEvent)
    router.POST("/event/update/:id", Controllers.UpdateById)


}
