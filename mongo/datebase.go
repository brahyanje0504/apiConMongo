package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Database struct {
	url          string
	nameDataBase string
}

func NewDataBase(url, nameDataBase string) *Database {
	return &Database{
		url:          url,
		nameDataBase: nameDataBase,
	}
}

func (d *Database) ConnectionDataBase() *mongo.Database {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(d.url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(d.nameDataBase)

	return database
}
