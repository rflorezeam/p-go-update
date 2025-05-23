package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConectarDB() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@libro-mongodb:27017")
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado a MongoDB!")
}

func GetCollection() *mongo.Collection {
	return Client.Database("libreria").Collection("libros")
} 