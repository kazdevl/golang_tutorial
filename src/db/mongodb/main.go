package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type respData struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Value  float64            `json:"value" bson:"value"`
	Sample string             `json:"sample" bson:"sample"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("access mongodb")

	collection := client.Database("testing").Collection("numbers")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	docs := []interface{}{
		bson.D{{Key: "name", Value: "pi"}, {Key: "value", Value: 3.14159}},
		bson.D{{Key: "name", Value: "pi"}, {Key: "value", Value: 2 * 3.14159}},
		bson.D{{Key: "name", Value: "pi"}, {Key: "sample", Value: "sample"}},
		bson.D{{Key: "name", Value: "sample"}, {Key: "value", Value: 10}},
	}
	res, err := collection.InsertMany(ctx, docs)
	fmt.Println(res)

	filter := bson.D{{Key: "name", Value: "pi"}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var res respData
		cur.Decode(&res)
		fmt.Printf("res(find): %+v\n", res)
	}
}
