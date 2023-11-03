package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetDb(env string) (*mongo.Database, func()) {
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@service-course-db:27017"))
	if err != nil {
		panic(err)
	}

	var disconnect = func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	
	dbName := "service-courses"
	if env == "test" {
		dbName = "service-courses-test"
	}
	db := client.Database(dbName)

	return db, disconnect
}