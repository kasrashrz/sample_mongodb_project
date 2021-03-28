package app
import
(
    "kasra_medrick.com/mongo/Controllers"
    _ "kasra_medrick.com/mongo/Controllers"
    "kasra_medrick.com/mongo/Models"
)
var event Models.Event
func mapURLs (){
    router.GET("/ping", Controllers.Ping)
    router.GET("/event/:id", Controllers.GetOneEvent)
    router.GET("/events/all", Controllers.GetAllEvents)
	router.POST("/events/add", Controllers.CreateEvent)

}
