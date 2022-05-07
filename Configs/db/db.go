package db

import (
	"context"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
	"time"
)

var once sync.Once
var instance *mongo.Client

//*** DOT ENV ***//
func ViperConfigVariable(key string) string {
	viper.SetConfigName("setup")
	viper.AddConfigPath("./Configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

//*** DATA BASE ***//
func Init() *mongo.Client {

	once.Do(func() {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + ViperConfigVariable("DB_HOST") + ":27017"))
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_HOST")+":27017/"))

		//fmt.Println("mongodb://"+ViperConfigVariable("DB_HOST")+":27017")

		if err != nil {
			log.Fatalf("Failed to connect to mongo db %f", err)
		}

		instance = client
	})

	return instance
}
func GetCollection(collection string) *mongo.Collection {
	coll := instance.Database(os.Getenv("DB_NAME")).Collection(collection)
	return coll
}
