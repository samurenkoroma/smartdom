package useCase

import (
	"smartdom/services/libparser/internal/domain/book"
)

type LibraryParser interface {
	Parse(dir string) ([]book.Book, error)
	Save([]book.Book) error
}
