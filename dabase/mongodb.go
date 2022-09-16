package dabase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var NameDatabase = "personas"

func GetCollection(collection string) *mongo.Collection {
	url := "mongodb+srv://admin:tueliges123@tueligescedula.i6zqf.mongodb.net/tueligescedula?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	return client.Database(NameDatabase).Collection(collection)
}
