package db
import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var once sync.Once
var instance *mongo.Client
//*** DATA BASE ***//
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
//*** DOT ENV ***//
func DotEnv(input string) string {

	err := godotenv.Load(filepath.Join("/home/kasra/mong_golang/Configs/", "setup.env"))
	if err != nil {
		fmt.Println(err)
	}
	return os.Getenv(input)
}