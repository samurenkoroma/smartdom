package book

import (
	"smartdom/services/library/internal/domain/book/path"
	"smartdom/services/library/internal/domain/book/title"
)

type Book struct {
	Id    string
	Title title.Title
	Path  path.Path
}
