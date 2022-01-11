package database

import (
	"context"
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	env "auth-service/pkg/utils"
)

func Init() error {
	mongoDBUri := env.GoDotEnvVariable("MONGODB_URI")
	err := mgm.SetDefaultConfig(nil, "auth", options.Client().ApplyURI(mongoDBUri))
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to mongo on URI", mongoDBUri)
	fmt.Println("Setting Indexes Rules")
	fmt.Println("TEST TEXT")
	setUserIndexes()
	fmt.Println("Finished Setting Indexes Rules")

	return err
}

func setUserIndexes() {
	mod := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)}
	_, err := mgm.CollectionByName("users").Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}
}
