package user_repository

import (
	"context"
	"cv_api/database"
	m "cv_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "displayData"

var collection = database.GetCollection(collectionName)
var ctx = context.Background()

func Create(displayData m.DisplayData) error {

	var err error

	_, err = collection.InsertOne(ctx, displayData)

	if err != nil {
		return err
	}

	return nil
}

func Read(userId string) (m.DisplayData, error) {

	var displayData m.DisplayData

	var err error
	var oid primitive.ObjectID

	var cur *mongo.Cursor

	oid, err = primitive.ObjectIDFromHex(userId)

	if err != nil {
		return displayData, err
	}

	filter := bson.M{"created_by": oid}

	cur, err = collection.Find(ctx, filter)

	cur.Next(ctx)
	cur.Decode(&displayData)

	return displayData, err
}
