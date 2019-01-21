package dao

import (
	"context"
	"encoding/json"
	"github.com/mongodb/mongo-go-driver/mongo"
	"one/go-misc/utils"
)

func getMongoCollection(dbName string, collectionName string) (*mongo.Collection) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection

	uri := utils.MongoSetting.Url

	if client, err = mongo.Connect(context.Background(), uri); err != nil {
		panic(err)
	}

	collection = client.Database(dbName).Collection(collectionName)
	return collection
}

func stringify(doc interface{}, opts ...string) string {
	if len(opts) == 2 {
		b, _ := json.MarshalIndent(doc, opts[0], opts[1])
		return string(b)
	}
	b, _ := json.Marshal(doc)
	return string(b)
}
