package connetion

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewConnectionMongo() (*mongo.Client, error) {
	uri := "mongodb://root:12344321@localhost:27017/?maxPoolSize=20&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	return client, nil
}
