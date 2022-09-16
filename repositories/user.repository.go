package repositories

import (
	"basesdedatos/dabase"
	"basesdedatos/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var collection = dabase.GetCollection("users")
var ctx = context.Background()

func Create(user models.User) error {

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func Read() (models.Users, error) {

	users := models.Users{}
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func Update(user models.User, userId string) error {
	oid, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"nombre":     user.Nombre,
			"apellido":   user.Apellido,
			"updated_at": time.Now(),
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {

	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
