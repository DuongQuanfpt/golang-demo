package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Address struct {
	Building string 
	City     string 
	State    string 
}

func main() {
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, cancel, err := connectDB("mongodb+srv://Cluster07816:zGwLyXH2Jm0erB45@cluster07816.2p23g.mongodb.net/?retryWrites=true&w=majority&appName=Cluster07816")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer closeDB(client, ctx, cancel)
	coll := client.Database("LearningMongo").Collection("Address")

	docs := []interface{}{
		Address{Building: "Abc", City: "abc",State: "abc"},
		Address{Building: "idk", City: "idk",State: "idkv"},
	}

	_, err = coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []Address
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}