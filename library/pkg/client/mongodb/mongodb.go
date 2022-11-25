package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, database string) (db *mongo.Database, err error) {
	var mongoDbUrl string
	mongoDbUrl = fmt.Sprintf("mongodb://library:library@%s:%s", host, port)

	clientOpts := options.Client().ApplyURI(mongoDbUrl)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongodb to due error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb to due error: %v", err)
	}
	return client.Database(database), err
}
