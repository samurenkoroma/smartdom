package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"smarthome/library/internal/domain/book"
)

type db struct {
	collection *mongo.Collection
}

func (d *db) Create(ctx context.Context, b book.Book) (string, error) {
	result, err := d.collection.InsertOne(ctx, b)
	fmt.Println(result)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error %v", b)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", fmt.Errorf("failed to convert objectid to hex: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (book.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) FindAll(ctx context.Context) (books []book.Book, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return books, fmt.Errorf("failed to find all user due to err: %v", err)
	}

	if err := cursor.All(ctx, &books); err != nil {
		return books, fmt.Errorf("failed to read all documents from cursor due to err %v", err)
	}

	return books, nil
}

func (d *db) Update(ctx context.Context, b book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string) book.Storage {
	return &db{
		collection: database.Collection(collection),
	}
}
