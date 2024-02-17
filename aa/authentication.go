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

type NGLError struct {
	ErrorCode string
	Message   string
	Status    int
}

func Authenticate(username string, password string) (*Patron, *NGLError) {
	filter := bson.D{{"patron_id", username}, {"user_password", password}}

	var result Patron
	err := config.PatronCollection.FindOne(context.TODO(), filter).Decode(&result)

	fmt.Println(result.FName)

	if err != nil {
		return nil, &NGLError{ErrorCode: "AUTH-001", Message: "Authentication failed", Status: 400}
	}

	return &result, nil
}
