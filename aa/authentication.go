package aa

import (
	"context"
	"fmt"
	"nglxtreme/auth/config"

	"go.mongodb.org/mongo-driver/bson"
)

type Patron struct {
	PatronId string `bson:"patron_id"`
	FName    string `bson:"fname"`
}

func Authenticate() {
	filter := bson.D{{"patron_id", "1SI07TE401"}}

	var result Patron
	err := config.PatronCollection.FindOne(context.TODO(), filter).Decode(&result)

	fmt.Println(result.FName)

	if err != nil {
		panic(err)
	}
}
