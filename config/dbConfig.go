package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
var PatronCollection *mongo.Collection

func ConnectToDatabase() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", viper.GetString("datastore.username"), viper.GetString("datastore.password"), viper.GetString("datastore.host"), viper.GetInt("datastore.port"))

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))

	if err != nil {
		panic(err)
	}

	database = client.Database(viper.GetString("datastore.database"))
	PatronCollection = database.Collection("patron")
}
