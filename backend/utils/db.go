package utils

import (
	"context"
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDb() (*mongo.Database, func()) {
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}

	var disconnect = func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}

	dbName := "service-courses"
	db := client.Database(dbName)

	return db, disconnect
}

func GetUsersDb() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))

	if err != nil {
		panic(err)
	}

	return db
}
