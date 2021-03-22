package db
import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)


var once sync.Once
var instance *mongo.Client

func Init() *mongo.Client {

once.Do(func() {

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

if err != nil {
log.Fatalf("Failed to connect to mongo db %f", err)
}

instance = client
})

return instance
}

func GetCollection(collection string) *mongo.Collection {
coll := instance.Database("golang").Collection(collection)
return coll
}