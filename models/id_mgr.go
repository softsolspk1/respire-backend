package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IdCounter is used to track the next available ID for a collection
type IdCounter struct {
	Id    string `bson:"_id"`
	Count int    `bson:"count"`
}

// IdMgr handles ID generation for MongoDB collections
type IdMgr struct {
	DB *mongo.Database
}

// NextId generates the next available ID for a given collection
func (idm *IdMgr) NextId(collectionName string) (int, error) {
	collection := idm.DB.Collection("counters")
	ctx := context.Background()

	// Try to find and update in one atomic operation
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	filter := bson.M{"_id": collectionName}
	update := bson.M{"$inc": bson.M{"count": 1}}

	var counter IdCounter
	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&counter)
	if err != nil {
		return 0, err
	}

	return counter.Count, nil
}