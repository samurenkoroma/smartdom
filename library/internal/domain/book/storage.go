package book

import "context"

type Storage interface {
	Create(ctx context.Context, b Book) (string, error)
	FindOne(ctx context.Context, id string) (Book, error)
	FindAll(ctx context.Context) ([]Book, error)
	Update(ctx context.Context, b Book) error
	Delete(ctx context.Context, id string) error
}
