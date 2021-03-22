package app
import 
(
    _ "kasra_medrick.com/mongo/Controllers"
    "kasra_medrick.com/mongo/Models"
)
var event Models.Event
func mapURLs (){
    router.GET("/ping/:id", event.RETREIVE)
}
