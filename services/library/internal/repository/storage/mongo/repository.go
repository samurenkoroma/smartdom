package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"smartdom/services/library/internal/domain/book"
	"smartdom/services/library/internal/domain/book/path"
	"smartdom/services/library/internal/domain/book/title"
	"smartdom/services/library/internal/repository/storage/mongo/dao"
	"smartdom/services/library/internal/useCase/adapters/storage"
)

type Repository struct {
	collection *mongo.Collection
}

func (r Repository) GetById(id string) (book.Book, error) {
	var book book.Book
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}

	result := r.collection.FindOne(context.Background(), bson.M{"_id": hex})
	if err := result.Decode(&book); err != nil {
		return book, err

	}
	return book, nil
}

func (r Repository) ListBook() (books []book.Book, err error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if cursor.Err() != nil {
		return nil, fmt.Errorf("failed to find all user due to err: %v", err)
	}
	var dbBooks []dao.Book

	if err := cursor.All(context.Background(), &dbBooks); err != nil {
		fmt.Println(err)
		return books, fmt.Errorf("failed to read all documents from cursor due to err %v", err)
	}
	for _, item := range dbBooks {
		books = append(books, book.Book{
			Id:    item.Id,
			Title: title.Title(item.Title),
			Path:  path.Path(item.Path),
		})
	}

	return books, nil
}

func New(db *mongo.Database) storage.BookStorage {
	return &Repository{
		collection: db.Collection("Books"),
	}
}
