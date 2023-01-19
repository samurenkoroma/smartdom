package mongodbrepo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"smartdom/services/libparser/internal/domain/book"
	"smartdom/services/libparser/internal/useCase/parser/adapters"
)

type Repository struct {
	db *mongo.Collection
}

func (r Repository) Create(books []book.Book) error {
	inserts := make([]interface{}, len(books))
	for i := range books {
		inserts[i] = books[i]
	}
	_, err := r.db.InsertMany(context.Background(), inserts)
	if err != nil {
		return fmt.Errorf("failed to create user due to error %v", err)
	}
	return nil
}

func New(db *mongo.Database) adapters.Storage {
	return &Repository{
		db: db.Collection("Books"),
	}
}
