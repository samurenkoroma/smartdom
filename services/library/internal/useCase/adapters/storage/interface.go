package storage

import "smartdom/services/library/internal/domain/book"

type BookStorage interface {
	ListBook() ([]book.Book, error)
	GetById(id string) (book.Book, error)
}
