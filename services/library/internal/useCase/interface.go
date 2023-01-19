package useCase

import "smartdom/services/library/internal/domain/book"

type BookReader interface {
	List() ([]book.Book, error)
	GetBookByID(id string) (book.Book, error)
}
