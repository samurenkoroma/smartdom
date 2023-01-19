package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func (s Settings) connectString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", s.User, s.Password, s.Host, s.Port)
}

func NewClient(ctx context.Context, s Settings) (db *mongo.Database, err error) {

	clientOpts := options.Client().ApplyURI(s.connectString())

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongodb to due error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb to due error: %v", err)
	}
	return client.Database(s.Database), err
}
