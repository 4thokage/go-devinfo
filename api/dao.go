package api

import (
	"context"
	"fmt"
	"github.com/4thokage/go-devinfo/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func init() {
	var conf config.Config

	conf.Read()

	clientOptions := options.Client().ApplyURI(conf.Server)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(conf.Database)

	fmt.Println("Connected to MongoDB!")
}

// GetAllJobs returns all jobs from DB
func GetAllJobs() []Job {
	cur, err := db.Collection("jobs").Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var elements []Job

	var elem Job

	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		elements = append(elements, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.Background())

	return elements
}

// GetAllProjects returns all projects from DB
func GetAllProjects() []Project {
	cur, err := db.Collection("projects").Find(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	var elements []Project

	var elem Project
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		elements = append(elements, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(context.Background())

	return elements
}
