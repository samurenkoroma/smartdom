package adapters

import "smartdom/services/libparser/internal/domain/book"

type Storage interface {
	Create(books []book.Book) error
}
