package database

import (
	"fmt"

	"github.com/kamva/mgm/v3"
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
	return err
}
