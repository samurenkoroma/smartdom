package book

import (
	"fmt"
	"smartdom/services/library/internal/domain/book"
	"smartdom/services/library/internal/domain/book/path"
	"smartdom/services/library/internal/useCase"
	"smartdom/services/library/internal/useCase/adapters/storage"
)

type UseCase struct {
	store storage.BookStorage
}

func (u UseCase) GetBookByID(id string) (book.Book, error) {
	return u.store.GetById(id)
}

func (u UseCase) List() ([]book.Book, error) {
	listBook, err := u.store.ListBook()
	if err != nil {
		return nil, err
	}
	var result []book.Book
	for _, book := range listBook {
		book.Path = path.Path(fmt.Sprintf("http://localhost:8079/books/%s", book.Id))
		result = append(result, book)
	}
	return result, nil
}

func New(store storage.BookStorage) useCase.BookReader {
	return &UseCase{
		store: store,
	}
}
